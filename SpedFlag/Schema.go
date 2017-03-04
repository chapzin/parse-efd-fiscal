package SpedFlag

import (
	"github.com/chapzin/parse-efd-fiscal/createMigration"
	"flag"
	"fmt"
)
// TODO verificar como adicionar essa funcao no main
func Schema(){

	create := flag.Bool("create",false, "cria tabela")
	drop := flag.Bool("drop", false, "Dropa Tabela")

	flag.Parse()

	if *drop {
		fmt.Println("Deletando Tabelas....")
		createMigration.Drop()

	}


	if *create {
		fmt.Println("Criando Tabelas....")
		createMigration.Create()

	}



}
