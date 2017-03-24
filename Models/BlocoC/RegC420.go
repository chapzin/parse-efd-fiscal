package BlocoC

import (
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
	"time"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

type RegC420 struct {
	gorm.Model
	Reg        string  `gorm:"type:varchar(4)"`
	CodTotPar  string  `gorm:"type:varchar(7)"`
	VlrAcumTot float64 `gorm:"type:decimal(19,2)"`
	NrTot      string  `gorm:"type:varchar(2)"`
	DescrNrTot string
	DtIni      time.Time `gorm:"type:date"`
	DtFin      time.Time `gorm:"type:date"`
	Cnpj       string    `gorm:"type:varchar(14)"`
}

func (RegC420) TableName() string {
	return "reg_c420"
}

// Implementando Interface do Sped RegC420
type RegC420Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

type iRegC420 interface {
	GetRegC420() RegC420
}

func (s RegC420Sped) GetRegC420() RegC420 {
	regC420 := RegC420{
		Reg:        s.Ln[1],
		CodTotPar:  s.Ln[2],
		VlrAcumTot: tools.ConvFloat(s.Ln[3]),
		NrTot:      s.Ln[4],
		DescrNrTot: s.Ln[5],
		DtIni:      s.Reg0000.DtIni,
		DtFin:      s.Reg0000.DtFin,
		Cnpj:       s.Reg0000.Cnpj,
	}
	return regC420
}

func CreateRegC420(read iRegC420) RegC420 {
	return read.GetRegC420()
}
