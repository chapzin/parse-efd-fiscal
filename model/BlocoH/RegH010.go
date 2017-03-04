package BlocoH

import (
	"github.com/jinzhu/gorm"
	"time"
	"database/sql"
)

type RegH010 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	CodItem string		`gorm:"type:varchar(60)"`
	Unid string		`gorm:"type:varchar(6)"`
	Qtd sql.NullFloat64		`gorm:"type:decimal(19,3)"`
	VlUnit sql.NullFloat64		`gorm:"type:decimal(19,6)"`
	VlItem sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	IndProp string		`gorm:"type:varchar(1)"`
	CodPart string		`gorm:"type:varchar(60)"`
	TxtCompl string
	CodCta string
	VlItemIr sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`

}

func (RegH010) TableName() string {
	return "reg_h010"
}
