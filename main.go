package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"flag"
	"github.com/chapzin/parse-efd-fiscal/SpedDB"
	"github.com/chapzin/parse-efd-fiscal/SpedRead"
	"fmt"
	"time"
	"github.com/chapzin/parse-efd-fiscal/config"
)


func main() {
	cfg := new(config.Configurador)
	config.InicializaConfiguracoes(cfg)
	dialect, err := config.Propriedades.ObterTexto("bd.dialect")
	conexao, err := config.Propriedades.ObterTexto("bd.conexao")
	db, err := gorm.Open(dialect, conexao)
	if err != nil {
		fmt.Println("Falha ao abrir conexão. dialect=%s, Linha de Conexao=%s", dialect, conexao)
	}
	schema := flag.Bool("schema", false, "Recria as tabelas")
	flag.Parse()
	if *schema {
		// Recria o Schema do banco de dados
		SpedDB.Schema(*db)
	}
	// Lendo todos arquivos da pasta speds
	fmt.Println("Iniciando processamento ",time.Now())
	SpedRead.RecursiveSpeds("./speds", dialect,conexao)
	// Pega cada arquivo e ler linha a linha e envia para o banco de dados
	//SpedRead.AddAllSpeds(filesSpeds,*db)
	fmt.Println("Final processamento ",time.Now())
	var msg string
	fmt.Scanln(&msg)

}
