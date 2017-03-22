package SpedDB

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func Schema(db gorm.DB) {

	fmt.Println("Deletando Tabelas....")
	Drop(db)
	fmt.Println("Criando Tabelas....")
	Create(db)

}
