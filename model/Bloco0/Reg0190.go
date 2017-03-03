package Bloco0

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Reg0190 struct {
	gorm.Model
	Reg string	`gorm:"type:varchar(4)"`
	Unid string	`gorm:"type:varchar(6);unique_index"`
	Descr string
	DtIni time.Time `gorm:"type:date"`
	DtFin time.Time `gorm:"type:date"`
	Cnpj string	`gorm:"type:varchar(14)"`
}

func (Reg0190) TableName() string {
	return "reg_0190"
}