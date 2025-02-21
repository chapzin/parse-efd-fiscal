package SpedDB

import (
	"github.com/chapzin/parse-efd-fiscal/Models"
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoC"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoH"
	cupomfiscal "github.com/chapzin/parse-efd-fiscal/Models/CupomFiscal"
	"github.com/chapzin/parse-efd-fiscal/Models/NotaFiscal"
	"github.com/jinzhu/gorm"
)

func tables() []interface{} {
	tables := []interface{}{
		&Bloco0.Reg0000{},          //nolint:exhaustruct
		&Bloco0.Reg0150{},          //nolint:exhaustruct
		&Bloco0.Reg0190{},          //nolint:exhaustruct
		&Bloco0.Reg0200{},          //nolint:exhaustruct
		&Bloco0.Reg0220{},          //nolint:exhaustruct
		&BlocoC.RegC100{},          //nolint:exhaustruct
		&BlocoC.RegC170{},          //nolint:exhaustruct
		&BlocoC.RegC400{},          //nolint:exhaustruct
		&BlocoC.RegC405{},          //nolint:exhaustruct
		&BlocoC.RegC420{},          //nolint:exhaustruct
		&BlocoC.RegC425{},          //nolint:exhaustruct
		&BlocoC.RegC460{},          //nolint:exhaustruct
		&BlocoC.RegC465{},          //nolint:exhaustruct
		&BlocoC.RegC470{},          //nolint:exhaustruct
		&BlocoC.RegC490{},          //nolint:exhaustruct
		&BlocoC.RegC800{},          //nolint:exhaustruct
		&BlocoC.RegC860{},          //nolint:exhaustruct
		&BlocoC.RegC870{},          //nolint:exhaustruct
		&BlocoC.RegC890{},          //nolint:exhaustruct
		&BlocoH.RegH005{},          //nolint:exhaustruct
		&BlocoH.RegH010{},          //nolint:exhaustruct
		&Models.Inventario{},       //nolint:exhaustruct,misspell
		&NotaFiscal.Emitente{},     //nolint:exhaustruct
		&NotaFiscal.Destinatario{}, //nolint:exhaustruct
		&NotaFiscal.Item{},         //nolint:exhaustruct
		&NotaFiscal.NotaFiscal{},   //nolint:exhaustruct
		&cupomfiscal.CfeItem{},     //nolint:exhaustruct
		&cupomfiscal.CfeHeader{},   //nolint:exhaustruct
	}

	return tables
}

func Create(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(tables()...)

	db.Model(&cupomfiscal.CfeItem{}).AddForeignKey("id_header", "cfe_headers(id)", "RESTRICT", "RESTRICT") //nolint:exhaustruct
}

func Drop(db *gorm.DB) {
	// Drop the tables
	db.DropTable(tables()...)
}
