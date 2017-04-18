package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RegC490 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4)"`
	CstIcms  string    `gorm:"type:varchar(3)"`
	Cfop     string    `gorm:"type:varchar(4)"`
	AliqIcms float64   `gorm:"type:decimal(6,2)"`
	VlOpr    float64   `gorm:"type:decimal(19,2)"`
	VlBcIcms float64   `gorm:"type:decimal(19,2)"`
	VlIcms   float64   `gorm:"type:decimal(19,2)"`
	CodObs   string    `gorm:"type:varchar(6)"`
	DtIni    time.Time `gorm:"type:date"`
	DtFin    time.Time `gorm:"type:date"`
	Cnpj     string    `gorm:"type:varchar(14)"`
}

func (RegC490) TableName() string {
	return "reg_C490"
}
