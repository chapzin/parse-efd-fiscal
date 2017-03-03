package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RegC170 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	NumItem string		`gorm:"type:varchar(3)"`
	CodItem string		`gorm:"type:varchar(60)"`
	DescrCompl string
	Qtd float64		`gorm:"type:decimal(19,5)"`
	Unid string		`gorm:"type:varchar(6)"`
	VlItem float64		`gorm:"type:decimal(19,2)"`
	VlDesc float64		`gorm:"type:decimal(19,2)"`
	IndMov string		`gorm:"type:varchar(1)"`
	CstIcms string		`gorm:"type:varchar(3)"`
	Cfop string		`gorm:"type:varchar(4)"`
	CodNat string		`gorm:"type:varchar(10)"`
	VlBcIcms float64	`gorm:"type:decimal(19,2)"`
	AliqIcms float64	`gorm:"type:decimal(6,2)"`
	VlIcms float64		`gorm:"type:decimal(19,2)"`
	VlBcIcmsSt float64	`gorm:"type:decimal(19,2)"`
	AliqSt float64		`gorm:"type:decimal(19,2)"`
	VlIcmsSt float64	`gorm:"type:decimal(19,2)"`
	IndApur string		`gorm:"type:varchar(1)"`
	CstIpi string		`gorm:"type:varchar(2)"`
	CodEnq string		`gorm:"type:varchar(3)"`
	VlBcIpi float64		`gorm:"type:decimal(19,2)"`
	AliqIpi float64		`gorm:"type:decimal(6,2)"`
	VlIpi float64		`gorm:"type:decimal(19,2)"`
	CstPis string		`gorm:"type:varchar(2)"`
	VlBcPis float64		`gorm:"type:decimal(19,2)"`
	AliqPis01 float64	`gorm:"type:decimal(8,4)"`
	QuantBcPis float64	`gorm:"type:decimal(19,3)"`
	AliqPis02 float64	`gorm:"type:decimal(8,4)"`
	VlPis float64		`gorm:"type:decimal(19,2)"`
	CstCofins string	`gorm:"type:varchar(2)"`
	VlBcCofins float64	`gorm:"type:decimal(19,2)"`
	AliqCofins01 float64	`gorm:"type:decimal(8,4)"`
	QuantBcCofins float64	`gorm:"type:decimal(19,3)"`
	AliqCofins02 string	`gorm:"type:decimal(8,4)"`
	VlCofins float64	`gorm:"type:decimal(19,2)"`
	CodCta string
	EntradaSaida string  	`gorm:"type:varchar(1)"`// Se for entrada 0, se for saida 1
	NumDoc string		`gorm:"type:varchar(9)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`

}

func (RegC170) TableName() string {
	return "reg_c170"
}
