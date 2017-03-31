package Bloco0

import (
	"github.com/jinzhu/gorm"
	"time"

	"github.com/chapzin/parse-efd-fiscal/tools"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016
type Reg0200 struct {
	gorm.Model
	Reg        string `gorm:"type:varchar(4)"`
	CodItem    string `gorm:"type:varchar(60)"`
	DescrItem  string
	CodBarra   string
	CodAntItem string    `gorm:"type:varchar(60)"`
	UnidInv    string    `gorm:"type:varchar(6)"`
	TipoItem   string    `gorm:"type:varchar(2)"`
	CodNcm     string    `gorm:"type:varchar(8)"`
	ExIpi      string    `gorm:"type:varchar(3)"`
	CodGen     string    `gorm:"type:varchar(2)"`
	CodLst     string    `gorm:"type:varchar(5)"`
	AliqIcms   float64   `gorm:"type:decimal(6,2)"`
	DtIni      time.Time `gorm:"type:date"`
	DtFin      time.Time `gorm:"type:date"`
	Cnpj       string    `gorm:"type:varchar(14)"`
}

func (Reg0200) TableName() string {
	return "reg_0200"
}

// Implementando Interface do Sped Reg0200
type Reg0200Sped struct {
	Ln      []string
	Reg0000 Reg0000
	Digito  string
}

type iReg0200 interface {
	GetReg0200() Reg0200
}

func (s Reg0200Sped) GetReg0200() Reg0200 {
	digitoInt := tools.ConvInt(s.Digito)
	codigo := tools.AdicionaDigitosCodigo(s.Ln[2], digitoInt)
	reg0200 := Reg0200{
		Reg:        s.Ln[1],
		CodItem:    codigo,
		DescrItem:  s.Ln[3],
		CodBarra:   s.Ln[4],
		CodAntItem: s.Ln[5],
		UnidInv:    s.Ln[6],
		TipoItem:   s.Ln[7],
		CodNcm:     s.Ln[8],
		ExIpi:      s.Ln[9],
		CodGen:     s.Ln[10],
		CodLst:     s.Ln[11],
		AliqIcms:   tools.ConvFloat(s.Ln[12]),
		DtIni:      s.Reg0000.DtIni,
		DtFin:      s.Reg0000.DtFin,
		Cnpj:       s.Reg0000.Cnpj,
	}
	return reg0200
}

func CreateReg0200(read iReg0200) Reg0200 {
	return read.GetReg0200()
}
