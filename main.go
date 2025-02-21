package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Controllers"
	"github.com/chapzin/parse-efd-fiscal/SpedDB"
	"github.com/chapzin/parse-efd-fiscal/config"
	"github.com/chapzin/parse-efd-fiscal/pkg/worker"
	"github.com/chapzin/parse-efd-fiscal/read"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tealeg/xlsx"
)

var (
	schema       = flag.Bool("schema", false, "Recria as tabelas")
	importarSped = flag.Bool("importar-sped", false, "Importa os speds")
	importarXml  = flag.Bool("importar-xml", false, "Importa os xmls")
	inventario   = flag.Bool("inventario", false, "Fazer processamento do inventario")
	anoInicial   = flag.Int("anoInicial", 0, "Ano inicial do processamento do inventário")
	anoFinal     = flag.Int("anoFinal", 0, "Ano inicial do processamento do inventário")
	excel        = flag.Bool("excel", false, "Gera arquivo excel do inventario")
	h010         = flag.Bool("h010", false, "Gera arquivo h010 e 0200 no layout sped para ser importado")
)

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

	// Conecta ao banco de dados - conexão compartilhada
	db, err := gorm.Open(cfg.DB.Dialect, cfg.GetMySQLConnectionString())
	if err != nil {
		log.Fatalf("Falha ao abrir conexão com banco de dados: %v", err)
	}
	db.LogMode(true)
	defer db.Close()

	// Cria schema se necessário
	if *schema {
		SpedDB.Schema(db)
	}

	// Importa arquivos XML
	if *importarXml {
		log.Printf("Iniciando processamento de XML em %v", time.Now())

		if err := read.RecursiveXmls(db, cfg.SpedsPath, cfg.DigitCode); err != nil {
			log.Fatalf("erro ao processar XMLs: %v", err)
		}
		log.Printf("Final processamento em %v", time.Now())
	}

	// Importa arquivos SPED
	if *importarSped {
		log.Printf("Iniciando processamento de Sped em %v", time.Now())

		if err := read.RecursiveSpeds(db, cfg.SpedsPath, cfg.DigitCode); err != nil {
			log.Fatalf("erro ao processar SPEDs: %v", err)
		}
		log.Printf("Final processamento em %v", time.Now())
	}

	// Processa inventário
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

	// Gera arquivo Excel
	if *excel {
		log.Printf("Iniciando geração do arquivo Excel")
		if err := generateExcel(db); err != nil {
			log.Fatalf("Erro ao gerar arquivo Excel: %v", err)
		}
		log.Printf("Arquivo de Análise de Inventário gerado com sucesso")
	}

	// Processa H010 se necessário
	if *h010 && *anoInicial != 0 {
		log.Printf("Iniciando processamento H010 para o ano %d", *anoInicial)
		// Controllers.CriarH010InvInicial(*anoInicial, db)
		// Controllers.CriarH010InvFinal(*anoInicial, db)
	} else if *h010 {
		log.Fatal("Favor informar a tag ano. Exemplo: -ano=2016")
	}
}
