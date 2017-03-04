package BlocoC

import (
	"time"
	"github.com/jinzhu/gorm"
)

type RegC405 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	DtDoc time.Time		`gorm:"type:date"`
	Cro string		`gorm:"type:varchar(3)"`
	Crz string		`gorm:"type:varchar(6)"`
	NumCooFin string	`gorm:"type:varchar(9)"`
	GtFin float64		`gorm:"type:decimal(19,2)"`
	VlBrt float64		`gorm:"type:decimal(19,2)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`
}

func (RegC405) TableName() string {
	return "reg_c405"
}