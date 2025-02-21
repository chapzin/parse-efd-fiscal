package SpedDB

import (
	"github.com/chapzin/parse-efd-fiscal/Models"
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoC"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoH"
	"github.com/chapzin/parse-efd-fiscal/Models/NotaFiscal"
	"github.com/jinzhu/gorm"
)

func Create(db *gorm.DB) {

	// Migrate the schema
	db.AutoMigrate(&Bloco0.Reg0000{})
	db.AutoMigrate(&Bloco0.Reg0150{})
	db.AutoMigrate(&Bloco0.Reg0190{})
	db.AutoMigrate(&Bloco0.Reg0200{})
	db.AutoMigrate(&Bloco0.Reg0220{})
	db.AutoMigrate(&BlocoC.RegC100{})
	db.AutoMigrate(&BlocoC.RegC170{})
	db.AutoMigrate(&BlocoC.RegC400{})
	db.AutoMigrate(&BlocoC.RegC405{})
	db.AutoMigrate(&BlocoC.RegC420{})
	db.AutoMigrate(&BlocoC.RegC425{})
	db.AutoMigrate(&BlocoC.RegC460{})
	db.AutoMigrate(&BlocoC.RegC465{})
	db.AutoMigrate(&BlocoC.RegC470{})
	db.AutoMigrate(&BlocoC.RegC490{})
	db.AutoMigrate(&BlocoH.RegH005{})
	db.AutoMigrate(&BlocoH.RegH010{})
	db.AutoMigrate(&Models.Inventario{})
	db.AutoMigrate(&NotaFiscal.Emitente{})
	db.AutoMigrate(&NotaFiscal.Destinatario{})
	db.AutoMigrate(&NotaFiscal.Item{})
	db.AutoMigrate(&NotaFiscal.NotaFiscal{})

}

func Drop(db *gorm.DB) {

	// Drop the tables
	db.DropTable(&Bloco0.Reg0000{})
	db.DropTable(&Bloco0.Reg0150{})
	db.DropTable(&Bloco0.Reg0190{})
	db.DropTable(&Bloco0.Reg0200{})
	db.DropTable(&Bloco0.Reg0220{})
	db.DropTable(&BlocoC.RegC100{})
	db.DropTable(&BlocoC.RegC170{})
	db.DropTable(&BlocoC.RegC400{})
	db.DropTable(&BlocoC.RegC405{})
	db.DropTable(&BlocoC.RegC420{})
	db.DropTable(&BlocoC.RegC425{})
	db.DropTable(&BlocoC.RegC460{})
	db.DropTable(&BlocoC.RegC465{})
	db.DropTable(&BlocoC.RegC470{})
	db.DropTable(&BlocoC.RegC490{})
	db.DropTable(&BlocoH.RegH005{})
	db.DropTable(&BlocoH.RegH010{})
	db.DropTable(&Models.Inventario{})
	db.DropTable(&NotaFiscal.Emitente{})
	db.DropTable(&NotaFiscal.Destinatario{})
	db.DropTable(&NotaFiscal.Item{})
	db.DropTable(&NotaFiscal.NotaFiscal{})

}
