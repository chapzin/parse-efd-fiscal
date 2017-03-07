package Bloco0

import (
	"github.com/jinzhu/gorm"
	"time"

)

type Reg0200 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	CodItem string		`gorm:"type:varchar(60)"`
	DescrItem string
	CodBarra string
	CodAntItem string	`gorm:"type:varchar(60)"`
	UnidInv string		`gorm:"type:varchar(6)"`
	TipoItem string		`gorm:"type:varchar(2)"`
	CodNcm string		`gorm:"type:varchar(8)"`
	ExIpi string		`gorm:"type:varchar(3)"`
	CodGen string		`gorm:"type:varchar(2)"`
	CodLst string		`gorm:"type:varchar(5)"`
	AliqIcms float64	`gorm:"type:decimal(6,2)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`
}

func (Reg0200) TableName() string {
	return "reg_0200"
}