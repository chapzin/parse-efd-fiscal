package Bloco0

import (
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
	"time"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

type Reg0220 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4)"`
	UnidConv string    `gorm:"type:varchar(6)"`
	FatConv  float64   `gorm:"type:decimal(12,6)"`
	UnidCod  string    `gorm:"type:varchar(6)"`
	CodItem  string    `gorm:"type:varchar(60)"`
	DtIni    time.Time `gorm:"type:date"`
	DtFin    time.Time `gorm:"type:date"`
	Cnpj     string    `gorm:"type:varchar(14)"`
	Feito    string    `gorm:"type:varchar(1)"`
}

func (Reg0220) TableName() string {
	return "reg_0220"
}

// Implementando Interface do Sped Reg0220
type Reg0220Sped struct {
	Ln      []string
	Reg0000 Reg0000
	Reg0200 Reg0200
	Digito  string
}

func (s Reg0220Sped) GetReg0220() Reg0220 {
	digitoInt := tools.ConvInt(s.Digito)
	reg0220 := Reg0220{
		Reg:      s.Ln[1],
		UnidConv: s.Ln[2],
		FatConv:  tools.ConvFloat(s.Ln[3]),
		UnidCod:  s.Reg0200.UnidInv,
		CodItem:  tools.AdicionaDigitosCodigo(s.Reg0200.CodItem, digitoInt),
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
		Feito:    "0",
	}
	return reg0220
}

type iReg0220 interface {
	GetReg0220() Reg0220
}

func CreateReg0220(read iReg0220) Reg0220 {
	return read.GetReg0220()
}
