package BlocoC

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

type RegC405 struct {
	gorm.Model
	Reg       string                `gorm:"type:varchar(4)"`
	DtDoc     time.Time                `gorm:"type:date"`
	Cro       string                `gorm:"type:varchar(3)"`
	Crz       string                `gorm:"type:varchar(6)"`
	NumCooFin string        `gorm:"type:varchar(9)"`
	GtFin     float64                `gorm:"type:decimal(19,2)"`
	VlBrt     float64                `gorm:"type:decimal(19,2)"`
	DtIni     time.Time        `gorm:"type:date"`
	DtFin     time.Time        `gorm:"type:date"`
	Cnpj      string                `gorm:"type:varchar(14)"`
}

func (RegC405) TableName() string {
	return "reg_c405"
}

// Implementando Interface do Sped RegC405
type RegC405Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

type iRegC405 interface {
	GetRegC405() RegC405
}

func (s RegC405Sped) GetRegC405() RegC405 {
	regC405 := RegC405{
		Reg:       s.Ln[1],
		DtDoc:     tools.ConvertData(s.Ln[2]),
		Cro:       s.Ln[3],
		Crz:       s.Ln[4],
		NumCooFin: s.Ln[5],
		GtFin:     tools.ConvFloat(s.Ln[6]),
		VlBrt:     tools.ConvFloat(s.Ln[7]),
		DtIni:     s.Reg0000.DtIni,
		DtFin:     s.Reg0000.DtFin,
		Cnpj:      s.Reg0000.Cnpj,
	}
	return regC405
}

func CreateRegC405(read iRegC405) RegC405 {
	return read.GetRegC405()
}
