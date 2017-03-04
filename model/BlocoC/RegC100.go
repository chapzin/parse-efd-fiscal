package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
	"database/sql"
)

type RegC100 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	IndOper string		`gorm:"type:varchar(1)"`
	IndEmit string		`gorm:"type:varchar(1)"`
	CodPart string		`gorm:"type:varchar(60)"`
	CodMod string		`gorm:"type:varchar(2)"`
	CodSit string		`gorm:"type:varchar(2)"`
	Ser string		`gorm:"type:varchar(3)"`
	NumDoc string		`gorm:"type:varchar(9)"`
	ChvNfe string		`gorm:"type:varchar(44);unique_index"`
	DtDoc time.Time 	`gorm:"type:date"`
	DtES time.Time 		`gorm:"type:date"`
	VlDoc sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	IndPgto string		`gorm:"type:varchar(1)"`
	VlDesc sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlAbatNt sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	VlMerc sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	IndFrt string		`gorm:"type:varchar(1)"`
	VlFrt sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlSeg sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlOutDa sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlBcIcms sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	VlIcms sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlBcIcmsSt sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	VlIcmsSt sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	VlIpi sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlPis sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlCofins sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	VlPisSt sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlCofinsSt sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`
}

func (RegC100) TableName() string {
	return "reg_c100"
}
