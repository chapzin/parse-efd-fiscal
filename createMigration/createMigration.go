package createMigration

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"../model/Bloco0"
	"../model/BlocoC"
	"../model/BlocoH"
)




func Create(){

	db, err := gorm.Open("mysql","root@/auditoria2?charset=utf8")
	defer db.Close()
	if err != nil {
		panic(err.Error()+"Erro ao acessar banco de dados")

	}

	// Migrate the schema
	db.AutoMigrate(&Bloco0.Reg0000{})
	db.AutoMigrate(&Bloco0.Reg0150{})
	db.AutoMigrate(&Bloco0.Reg0190{})
	db.AutoMigrate(&Bloco0.Reg0200{})
	db.AutoMigrate(&Bloco0.Reg0220{})
	db.AutoMigrate(&BlocoC.RegC100{})
	db.AutoMigrate(&BlocoC.RegC170{})
	db.AutoMigrate(&BlocoH.RegH010{})

}

func Drop(){
	db, err := gorm.Open("mysql","root@/auditoria2?charset=utf8")
	defer db.Close()
	if err != nil {
		panic(err.Error()+"Erro ao acessar banco de dados")

	}
	// Drop the tables
	db.DropTable(&Bloco0.Reg0000{})
	db.DropTable(&Bloco0.Reg0150{})
	db.DropTable(&Bloco0.Reg0190{})
	db.DropTable(&Bloco0.Reg0200{})
	db.DropTable(&Bloco0.Reg0220{})
	db.DropTable(&BlocoC.RegC100{})
	db.DropTable(&BlocoC.RegC170{})
	db.DropTable(&BlocoH.RegH010{})

}


