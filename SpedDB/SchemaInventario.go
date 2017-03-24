package SpedDB

import (
	"fmt"
	"github.com/chapzin/GoInventario/Models"
	"github.com/jinzhu/gorm"
)
// funcao cria estrutura do DB
func CreateSchemaInventario(db gorm.DB) {
	fmt.Println("Criando schema do inventario...")
	db.AutoMigrate(&Models.Inventario{})
}

// Funcao dropa estrutura do DB
func DropSchemaInventario(db gorm.DB) {
	fmt.Println("Deletando schema do inventario...")
	db.DropTable(&Models.Inventario{})
}
