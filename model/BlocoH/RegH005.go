package BlocoH

import (
	"time"
	"github.com/jinzhu/gorm"
)

type RegH005 struct {
	gorm.Model
	Reg 	string		`gorm:"type:varchar(4)"`
	DtInv 	time.Time	`gorm:"type:date"`
	VlInv 	float64		`gorm:"type:decimal(19,2)"`
	MotInv 	string		`gorm:"type:varchar(2)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`
}

func (RegH005) TableName() string {
	return "reg_h005"
}