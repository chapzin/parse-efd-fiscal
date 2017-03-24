package BlocoH

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

type RegH005 struct {
	gorm.Model
	Reg    string                `gorm:"type:varchar(4)"`
	DtInv  time.Time        `gorm:"type:date"`
	VlInv  float64                `gorm:"type:decimal(19,2)"`
	MotInv string                `gorm:"type:varchar(2)"`
	DtIni  time.Time        `gorm:"type:date"`
	DtFin  time.Time        `gorm:"type:date"`
	Cnpj   string                `gorm:"type:varchar(14)"`
}

func (RegH005) TableName() string {
	return "reg_h005"
}

// Implementando Interface do Sped RegH005
type RegH005Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

type iRegH005 interface {
	GetRegH005() RegH005
}

func (s RegH005Sped) GetRegH005() RegH005 {
	regH005 := RegH005{
		Reg:    s.Ln[1],
		DtInv:  tools.ConvertData(s.Ln[2]),
		VlInv:  tools.ConvFloat(s.Ln[3]),
		MotInv: s.Ln[4],
		DtIni:  s.Reg0000.DtIni,
		DtFin:  s.Reg0000.DtFin,
		Cnpj:   s.Reg0000.Cnpj,
	}
	return regH005
}

func CreateRegH005(read iRegH005) RegH005 {
	return read.GetRegH005()
}
