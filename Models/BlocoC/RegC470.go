package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RegC470 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4)"`
	CodItem  string    `gorm:"type:varchar(60)"`
	Qtd      float64   `gorm:"type:decimal(19,2)"`
	QtdCanc  float64   `gorm:"type:decimal(19,2)"`
	Unid     string    `gorm:"type:varchar(6)"`
	VlItem   float64   `gorm:"type:decimal(19,2)"`
	CstIcms  string    `gorm:"type:varchar(3)"`
	Cfop     string    `gorm:"type:varchar(4)"`
	AliqIcms float64   `gorm:"type:decimal(6,2)"`
	VlPis    float64   `gorm:"type:decimal(19,2)"`
	VlCofins float64   `gorm:"type:decimal(19,2)"`
	DtIni    time.Time `gorm:"type:date"`
	DtFin    time.Time `gorm:"type:date"`
	Cnpj     string    `gorm:"type:varchar(14)"`
}

func (RegC470) TableName() string {
	return "reg_C470"
}
