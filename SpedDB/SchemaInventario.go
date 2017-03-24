package SpedDB

import (
	"fmt"
	"github.com/chapzin/GoInventario/Models"
	"github.com/jinzhu/gorm"
)

func CreateSchemaInventario(db gorm.DB) {
	fmt.Println("Criando schema do inventario...")
	db.AutoMigrate(&Models.Inventario{})
}

func DropSchemaInventario(db gorm.DB) {
	fmt.Println("Deletando schema do inventario...")
	db.DropTable(&Models.Inventario{})
}
