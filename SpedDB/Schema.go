package SpedDB

import (
	"github.com/chapzin/parse-efd-fiscal/createMigration"
	"fmt"
	"github.com/jinzhu/gorm"
)

func Schema(db gorm.DB) {

	fmt.Println("Deletando Tabelas....")
	createMigration.Drop(db)
	fmt.Println("Criando Tabelas....")
	createMigration.Create(db)

}
