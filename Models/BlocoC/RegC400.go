package BlocoC

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

type RegC400 struct {
	gorm.Model
	Reg    string                `gorm:"type:varchar(4)"`
	CodMod string                `gorm:"type:varchar(2)"`
	EcfMod string                `gorm:"type:varchar(20)"`
	EcfFab string                `gorm:"type:varchar(21)"`
	EcfCx  string                `gorm:"type:varchar(3)"`
	DtIni  time.Time        `gorm:"type:date"`
	DtFin  time.Time        `gorm:"type:date"`
	Cnpj   string                `gorm:"type:varchar(14)"`
}

func (RegC400) TableName() string {
	return "reg_c400"
}

// Implementando Interfaces do Sped RegC400
type RegC400Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

type iRegC400 interface {
	GetRegC400() RegC400
}

func (s RegC400Sped) GetRegC400() RegC400 {
	regC400 := RegC400{
		Reg:    s.Ln[1],
		CodMod: s.Ln[2],
		EcfMod: s.Ln[3],
		EcfFab: s.Ln[4],
		EcfCx:  s.Ln[5],
		DtIni:  s.Reg0000.DtIni,
		DtFin:  s.Reg0000.DtFin,
		Cnpj:   s.Reg0000.Cnpj,
	}
	return regC400
}

func CreateRegC400(read iRegC400) RegC400 {
	return read.GetRegC400()
}
