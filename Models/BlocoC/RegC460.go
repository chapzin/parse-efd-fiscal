package BlocoC

import (
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/jinzhu/gorm"
	"time"
)

type RegC460 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4)"`
	CodMod   string    `gorm:"type:varchar(2)"`
	CodSit   string    `gorm:"type:varchar(2)"`
	NumDoc   string    `gorm:"type:varchar(9)"`
	DtDoc    time.Time `gorm:"type:date"`
	VlDoc    float64   `gorm:"type:decimal(19,2)"`
	VlPis    float64   `gorm:"type:decimal(19,2)"`
	VlCofins float64   `gorm:"type:decimal(19,2)"`
	CpfCnpj  string    `gorm:"type:varchar(14)"`
	NomAdq   string    `gorm:"type:varchar(60)"`
	DtIni    time.Time `gorm:"type:date"`
	DtFin    time.Time `gorm:"type:date"`
	Cnpj     string    `gorm:"type:varchar(14)"`
}

func (RegC460) TableName() string {
	return "reg_C460"
}

type RegC460Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}
