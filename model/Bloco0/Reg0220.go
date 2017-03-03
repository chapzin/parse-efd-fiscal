package Bloco0

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Reg0220 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	UnidConv string		`gorm:"type:varchar(6)"`
	FatConv float64		`gorm:"type:decimal(12,6)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`
}

func (Reg0220) TableName() string {
	return "reg_0220"
}