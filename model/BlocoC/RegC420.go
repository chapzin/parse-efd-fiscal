package BlocoC

import (
	"time"
	"github.com/jinzhu/gorm"
)

type RegC420 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	CodTotPar string	`gorm:"type:varchar(7)"`
	VlrAcumTot float64	`gorm:"type:decimal(19,2)"`
	NrTot string		`gorm:"type:varchar(2)"`
	DescrNrTot string
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`
}

func (RegC420) TableName() string {
	return "reg_c420"
}