package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
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
	VlDoc float64		`gorm:"type:decimal(19,2)"`
	IndPgto string		`gorm:"type:varchar(1)"`
	VlDesc float64		`gorm:"type:decimal(19,2)"`
	VlAbatNt float64	`gorm:"type:decimal(19,2)"`
	VlMerc float64		`gorm:"type:decimal(19,2)"`
	IndFrt string		`gorm:"type:varchar(1)"`
	VlFrt float64		`gorm:"type:decimal(19,2)"`
	VlSeg float64		`gorm:"type:decimal(19,2)"`
	VlOutDa float64		`gorm:"type:decimal(19,2)"`
	VlBcIcms float64	`gorm:"type:decimal(19,2)"`
	VlIcms float64		`gorm:"type:decimal(19,2)"`
	VlBcIcmsSt float64	`gorm:"type:decimal(19,2)"`
	VlIcmsSt float64	`gorm:"type:decimal(19,2)"`
	VlIpi float64		`gorm:"type:decimal(19,2)"`
	VlPis float64		`gorm:"type:decimal(19,2)"`
	VlCofins float64	`gorm:"type:decimal(19,2)"`
	VlPisSt float64		`gorm:"type:decimal(19,2)"`
	VlCofinsSt float64	`gorm:"type:decimal(19,2)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`
}

func (RegC100) TableName() string {
	return "reg_c100"
}
