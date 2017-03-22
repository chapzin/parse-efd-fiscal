package main

import (
	"github.com/chapzin/parse-efd-fiscal/SpedError"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"flag"
	"github.com/chapzin/parse-efd-fiscal/SpedDB"
	"github.com/chapzin/parse-efd-fiscal/SpedRead"

	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	db, err := gorm.Open("mysql", "root@/auditoria2?charset=utf8")
	schema := flag.Bool("schema", false, "Recria as tabelas")
	flag.Parse()
	if *schema {
		// Recria o Schema do banco de dados
		SpedDB.Schema(*db)
	}
	SpedError.CheckErr(err)
	wg.Add(1)
	time.Sleep(15 * time.Second)
	// Lendo todos arquivos da pasta speds
	fmt.Println("Iniciando processamento ",time.Now())
	go SpedRead.RecursiveSpeds("./speds", &wg)
	// Pega cada arquivo e ler linha a linha e envia para o banco de dados
	//SpedRead.AddAllSpeds(filesSpeds,*db)
	wg.Wait()
	fmt.Println("Final processamento ",time.Now())
	var msg string
	fmt.Scanln(&msg)

}
