package main

import (
	"./createMigration"
	"flag"
	"fmt"
)

func main() {

	create := flag.Bool("create",false, "cria tabela")
	drop := flag.Bool("drop", false, "Dropa Tabela")

	flag.Parse()

	if *drop  {
		fmt.Println("Deletando Tabelas....")
		createMigration.Drop()
	}


	if *create {
		fmt.Println("Criando Tabelas....")
		createMigration.Create()
	}


}
