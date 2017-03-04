package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
	"database/sql"
)

type RegC170 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	NumItem string		`gorm:"type:varchar(3)"`
	CodItem string		`gorm:"type:varchar(60)"`
	DescrCompl string
	Qtd sql.NullFloat64		`gorm:"type:decimal(19,5)"`
	Unid string		`gorm:"type:varchar(6)"`
	VlItem sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlDesc sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	IndMov string		`gorm:"type:varchar(1)"`
	CstIcms string		`gorm:"type:varchar(3)"`
	Cfop string		`gorm:"type:varchar(4)"`
	CodNat string		`gorm:"type:varchar(10)"`
	VlBcIcms sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	AliqIcms sql.NullFloat64	`gorm:"type:decimal(6,2)"`
	VlIcms sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlBcIcmsSt sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	AliqSt sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlIcmsSt sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	IndApur string		`gorm:"type:varchar(1)"`
	CstIpi string		`gorm:"type:varchar(2)"`
	CodEnq string		`gorm:"type:varchar(3)"`
	VlBcIpi sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	AliqIpi sql.NullFloat64		`gorm:"type:decimal(6,2)"`
	VlIpi sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	CstPis string		`gorm:"type:varchar(2)"`
	VlBcPis sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	AliqPis01 sql.NullFloat64	`gorm:"type:decimal(8,4)"`
	QuantBcPis sql.NullFloat64	`gorm:"type:decimal(19,3)"`
	AliqPis02 sql.NullFloat64	`gorm:"type:decimal(8,4)"`
	VlPis sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	CstCofins string	`gorm:"type:varchar(2)"`
	VlBcCofins sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	AliqCofins01 sql.NullFloat64	`gorm:"type:decimal(8,4)"`
	QuantBcCofins sql.NullFloat64	`gorm:"type:decimal(19,3)"`
	AliqCofins02 sql.NullFloat64	`gorm:"type:decimal(8,4)"`
	VlCofins sql.NullFloat64	`gorm:"type:decimal(19,2)"`
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


