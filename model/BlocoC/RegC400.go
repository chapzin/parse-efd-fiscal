package BlocoC

import (
	"time"
	"github.com/jinzhu/gorm"
)

type RegC400 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	CodMod string		`gorm:"type:varchar(2)"`
	EcfMod string		`gorm:"type:varchar(20)"`
	EcfFab string		`gorm:"type:varchar(21)"`
	EcfCx string		`gorm:"type:varchar(3)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`
}

func (RegC400) TableName() string {
	return "reg_c400"
}