package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/chapzin/parse-efd-fiscal/model/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

type RegC425 struct {
	gorm.Model
	Reg      string                `gorm:"type:varchar(4)"`
	CodItem  string                `gorm:"type:varchar(60)"`
	Qtd      float64                `gorm:"type:decimal(19,3)"`
	Unid     string                `gorm:"type:varchar(6)"`
	VlItem   float64                `gorm:"type:decimal(19,2)"`
	VlPis    float64                `gorm:"type:decimal(19,2)"`
	VlCofins float64        `gorm:"type:decimal(19,2)"`
	DtIni    time.Time        `gorm:"type:date"`
	DtFin    time.Time        `gorm:"type:date"`
	Cnpj     string                `gorm:"type:varchar(14)"`
}

func (RegC425) TableName() string {
	return "reg_c425"
}

// Implementando Interface do Sped RegC425
type RegC425Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

type iRegC425 interface {
	GetRegC425() RegC425
}

func (s RegC425Sped) GetRegC425() RegC425 {
	regC425 := RegC425{
		Reg:      s.Ln[1],
		CodItem:  s.Ln[2],
		Qtd:      SpedConvert.ConvFloat(s.Ln[3]),
		Unid:     s.Ln[4],
		VlItem:   SpedConvert.ConvFloat(s.Ln[5]),
		VlPis:    SpedConvert.ConvFloat(s.Ln[6]),
		VlCofins: SpedConvert.ConvFloat(s.Ln[7]),
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
	}
	return regC425
}

func CreateRegC425(read iRegC425) RegC425 {
	return read.GetRegC425()
}
