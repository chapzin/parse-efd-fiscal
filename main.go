package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Controllers"
	"github.com/chapzin/parse-efd-fiscal/SpedDB"
	"github.com/chapzin/parse-efd-fiscal/SpedRead"
	"github.com/chapzin/parse-efd-fiscal/config"
	"github.com/chapzin/parse-efd-fiscal/pkg/worker"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tealeg/xlsx"
)

var schema = flag.Bool("schema", false, "Recria as tabelas")
var importa = flag.Bool("importa", false, "Importa os xmls e speds ")
var inventario = flag.Bool("inventario", false, "Fazer processamento do inventario")
var anoInicial = flag.Int("anoInicial", 0, "Ano inicial do processamento do inventário")
var anoFinal = flag.Int("anoFinal", 0, "Ano inicial do processamento do inventário")
var excel = flag.Bool("excel", false, "Gera arquivo excel do inventario")
var h010 = flag.Bool("h010", false, "Gera arquivo h010 e 0200 no layout sped para ser importado")

func validateInventoryFlags(anoInicial, anoFinal int) error {
	if anoInicial == 0 || anoFinal == 0 {
		return fmt.Errorf("favor informar o ano inicial e final que deseja processar. Exemplo -anoInicial=2011 -anoFinal=2016")
	}
	if anoInicial <= 2011 || anoFinal <= 2011 {
		return fmt.Errorf("favor informar um ano maior que 2011")
	}
	if anoInicial <= 999 || anoFinal <= 999 {
		return fmt.Errorf("favor informar o ano com 4 digitos. Exemplo 2017")
	}
	if anoInicial > anoFinal {
		return fmt.Errorf("o ano inicial deve ser menor que o ano final")
	}
	return nil
}

func generateExcel(db *gorm.DB) error {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet(tools.PLANILHA)
	if err != nil {
		return fmt.Errorf("erro ao criar planilha: %v", err)
	}

	Controllers.ExcelMenu(sheet)
	Controllers.ExcelAdd(db, sheet)

	if err := file.Save("AnaliseInventario.xlsx"); err != nil {
		return fmt.Errorf("erro ao salvar arquivo Excel: %v", err)
	}

	return nil
}

func processInventory(db *gorm.DB, anoInicial, anoFinal int) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("erro ao carregar configurações: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Worker.TaskTimeout)
	defer cancel()

	// Recria tabela de inventário
	SpedDB.DropSchemaInventario(db)
	SpedDB.CreateSchemaInventario(db)

	// Cria pool com configurações carregadas
	pool := worker.NewPool(ctx, cfg.Worker.MaxWorkers)
	pool.Start()

	// Primeira fase
	log.Printf("Iniciando primeira fase do processamento")
	pool.Submit(func() error {
		Controllers.ProcessarFatorConversao(db, nil)
		return nil
	})
	pool.Submit(func() error {
		Controllers.DeletarItensNotasCanceladas(db, "2012-01-01", "2016-12-31", nil)
		return nil
	})

	if errs := pool.Wait(); len(errs) > 0 {
		return fmt.Errorf("erros na primeira fase: %v", errs)
	}

	// Segunda fase
	log.Printf("Iniciando segunda fase do processamento")
	pool.Submit(func() error {
		Controllers.PopularReg0200(db, nil)
		return nil
	})
	pool.Submit(func() error {
		Controllers.PopularItensXmls(db, nil)
		return nil
	})

	if errs := pool.Wait(); len(errs) > 0 {
		return fmt.Errorf("erros na segunda fase: %v", errs)
	}

	// Terceira fase
	log.Printf("Iniciando terceira fase do processamento")
	pool.Submit(func() error {
		Controllers.PopularInventarios(anoInicial, anoFinal, nil, db)
		return nil
	})
	pool.Submit(func() error {
		Controllers.PopularEntradas(anoInicial, anoFinal, nil, db)
		return nil
	})
	pool.Submit(func() error {
		Controllers.PopularSaidas(anoInicial, anoFinal, nil, db)
		return nil
	})

	if errs := pool.Wait(); len(errs) > 0 {
		return fmt.Errorf("erros na terceira fase: %v", errs)
	}

	// Processamento final
	log.Printf("Iniciando processamento final")
	Controllers.ProcessarDiferencas(db)

	return nil
}

func main() {
	flag.Parse()

	// Carrega configurações do ambiente
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Conecta ao banco de dados
	db, err := gorm.Open(cfg.DB.Dialect, cfg.GetMySQLConnectionString())
	if err != nil {
		log.Fatalf("Falha ao abrir conexão com banco de dados: %v", err)
	}
	db.LogMode(true)
	defer db.Close()

	if *schema {
		SpedDB.Schema(db)
	}

	if *importa {
		log.Printf("Iniciando processamento em %v", time.Now())
		SpedRead.RecursiveSpeds(cfg.SpedsPath, cfg.DB.Dialect, cfg.GetMySQLConnectionString(), cfg.DigitCode)
		log.Printf("Final processamento em %v", time.Now())
		var s string
		fmt.Scanln(&s)
	}

	if *inventario {
		if err := validateInventoryFlags(*anoInicial, *anoFinal); err != nil {
			log.Fatal(err)
		}

		log.Printf("Iniciando processamento do inventário em %v", time.Now())
		if err := processInventory(db, *anoInicial, *anoFinal); err != nil {
			log.Fatalf("Erro no processamento do inventário: %v", err)
		}
		log.Printf("Processamento do inventário finalizado em %v", time.Now())
	}

	if *excel {
		log.Printf("Iniciando geração do arquivo Excel")
		if err := generateExcel(db); err != nil {
			log.Fatalf("Erro ao gerar arquivo Excel: %v", err)
		}
		log.Printf("Arquivo de Análise de Inventário gerado com sucesso")
	}

	if *h010 {
		if *anoInicial != 0 {
			log.Printf("Funcionalidade h010 ainda não implementada")
			//Controllers.CriarH010InvInicial(*ano, *db)
			//Controllers.CriarH010InvFinal(*ano, *db)
		} else {
			log.Fatal("Favor informar a tag ano. Exemplo: -ano=2016")
		}
	}
}
