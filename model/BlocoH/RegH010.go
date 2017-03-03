package BlocoH

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RegH010 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	CodItem string		`gorm:"type:varchar(60)"`
	Unid string		`gorm:"type:varchar(6)"`
	Qtd float64		`gorm:"type:decimal(19,3)"`
	VlUnit float64		`gorm:"type:decimal(19,6)"`
	VlItem float64		`gorm:"type:decimal(19,2)"`
	IndProp string		`gorm:"type:varchar(1)"`
	CodPart string		`gorm:"type:varchar(60)"`
	TxtCompl string
	CodCta string
	VlItemIr float64	`gorm:"type:decimal(19,2)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`

}

func (RegH010) TableName() string {
	return "reg_h010"
}
