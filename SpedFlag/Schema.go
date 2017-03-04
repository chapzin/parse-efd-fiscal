package SpedFlag

import (
	"github.com/chapzin/parse-efd-fiscal/createMigration"
	"flag"
	"fmt"
)

func SpedFlag() bool{

	create := flag.Bool("create",false, "cria tabela")
	drop := flag.Bool("drop", false, "Dropa Tabela")

	flag.Parse()

	if *drop  {
		fmt.Println("Deletando Tabelas....")
		createMigration.Drop()
		return *drop
	} else {
		return *drop
	}


	if *create {
		fmt.Println("Criando Tabelas....")
		createMigration.Create()
		return *create
	} else {
		return *create
	}




}
