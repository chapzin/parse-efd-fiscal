package main

import (
	"flag"
	"fmt"
	"github.com/chapzin/parse-efd-fiscal/Controllers"
	"github.com/chapzin/parse-efd-fiscal/SpedDB"
	"github.com/chapzin/parse-efd-fiscal/SpedRead"
	"github.com/chapzin/parse-efd-fiscal/config"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tealeg/xlsx"
	"os"
	"sync"
	"time"
)

var schema = flag.Bool("schema", false, "Recria as tabelas")
var importa = flag.Bool("importa", false, "Importa os xmls e speds ")
var inventario = flag.Bool("inventario", false, "Fazer processamento do inventario")
var anoInicial = flag.Int("anoInicial", 0, "Ano inicial do processamento do inventário")
var anoFinal = flag.Int("anoFinal", 0, "Ano inicial do processamento do inventário")
var excel = flag.Bool("excel", false, "Gera arquivo excel do inventario")
var h010 = flag.Bool("h010", false, "Gera arquivo h010 e 0200 no layout sped para ser importado")

func init() {
	flag.Parse()
	cfg := new(config.Configurador)
	config.InicializaConfiguracoes(cfg)
}

func main() {
	dialect, err := config.Propriedades.ObterTexto("bd.dialect")
	conexao, err := config.Propriedades.ObterTexto("bd.conexao.mysql")
	digitos, err := config.Propriedades.ObterTexto("bd.digit.cod")
	db, err := gorm.Open(dialect, conexao)
	db.LogMode(true)
	//defer db.Close()
	if err != nil {
		fmt.Println("Falha ao abrir conexão. dialect=?, Linha de Conexao=?", dialect, conexao)
		return
	}

	if *schema {
		// Recria o Schema do banco de dados
		SpedDB.Schema(*db)
	}

	if *importa {
		// Lendo todos arquivos da pasta speds
		fmt.Println("Iniciando processamento ", time.Now())
		SpedRead.RecursiveSpeds("./speds", dialect, conexao, digitos)
		// Pega cada arquivo e ler linha a linha e envia para o banco de dados
		fmt.Println("Final processamento ", time.Now())
		var s string
		fmt.Scanln(&s)
	}

	if *inventario {
		//Recria tabela de inventário
		SpedDB.DropSchemaInventario(*db)
		SpedDB.CreateSchemaInventario(*db)

		// Processa o inventário
		fmt.Println("Inventario começou a processar as ?", time.Now())
		var wg sync.WaitGroup

		if *anoInicial == 0 || *anoFinal == 0 {
			fmt.Println("Favor informar o ano inicial que deseja processar. Exemplo -anoInicial=2011")
			return
		} else if *anoInicial <= 2011 || *anoFinal <= 2011 {
			fmt.Println("Favor informar um ano maior que 2011")
			return
		} else if *anoInicial <= 999 || *anoFinal <= 999 {
			fmt.Println("Favor informar o ano com 4 digitos. Exemplo 2017")
			return
		} else if *anoInicial > *anoFinal {
			fmt.Println("O ano inicial deve ser menor que o ano final")
		}

		wg.Add(2)
		go Controllers.ProcessarFatorConversao(*db, &wg)
		go Controllers.DeletarItensNotasCanceladas(*db, "2012-01-01", "2016-12-31", &wg)
		wg.Wait()

		wg.Add(2)
		go Controllers.PopularReg0200(*db, &wg)
		go Controllers.PopularItensXmls(*db, &wg)
		wg.Wait()

		wg.Add(3)
		go Controllers.PopularInventarios(*anoInicial, *anoFinal, &wg, *db)
		go Controllers.PopularEntradas(*anoInicial, *anoFinal, &wg, *db)
		go Controllers.PopularSaidas(*anoInicial, *anoFinal, &wg, *db)
		wg.Wait()

		// Quando finalizar todas essas deve rodar o processar diferencas
		Controllers.ProcessarDiferencas(*db)
		time.Sleep(90 * time.Second)
		fmt.Println(time.Now())
		color.Green("TERMINOU")
	}

	if *excel {
		var file *xlsx.File
		var sheet *xlsx.Sheet
		var err error

		file = xlsx.NewFile()

		sheet, err = file.AddSheet(tools.PLANILHA)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		Controllers.ExcelMenu(sheet)
		Controllers.ExcelAdd(*db, sheet)

		err = file.Save("AnaliseInventario.xlsx")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("Arquivo de Analise Inventario Gerado com Sucesso!!!")
		}
	}

	if *h010 {

		if *anoInicial != 0 {
			//Controllers.CriarH010InvInicial(*ano, *db)
			//Controllers.CriarH010InvFinal(*ano, *db)
		} else {
			fmt.Println("Favor informar a tag ano. Exemplo: -ano=2016")
		}

	}

}
