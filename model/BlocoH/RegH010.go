package BlocoH

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/chapzin/parse-efd-fiscal/model/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
)

type RegH010 struct {
	gorm.Model
	Reg      string                `gorm:"type:varchar(4)"`
	CodItem  string                `gorm:"type:varchar(60)"`
	Unid     string                `gorm:"type:varchar(6)"`
	Qtd      float64                `gorm:"type:decimal(19,3)"`
	VlUnit   float64                `gorm:"type:decimal(19,6)"`
	VlItem   float64                `gorm:"type:decimal(19,2)"`
	IndProp  string                `gorm:"type:varchar(1)"`
	CodPart  string                `gorm:"type:varchar(60)"`
	TxtCompl string
	CodCta   string
	VlItemIr float64        `gorm:"type:decimal(19,2)"`
	DtIni    time.Time        `gorm:"type:date"`
	DtFin    time.Time        `gorm:"type:date"`
	Cnpj     string                `gorm:"type:varchar(14)"`
}

func (RegH010) TableName() string {
	return "reg_h010"
}

// Implementando Interface do Sped RegH010
type RegH010Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

type iRegH010 interface {
	GetRegH010() RegH010
}

func (s RegH010Sped) GetRegH010() RegH010 {
	regH010 := RegH010{
		Reg:      s.Ln[1],
		CodItem:  s.Ln[2],
		Unid:     s.Ln[3],
		Qtd:      SpedConvert.ConvFloat(s.Ln[4]),
		VlUnit:   SpedConvert.ConvFloat(s.Ln[5]),
		VlItem:   SpedConvert.ConvFloat(s.Ln[6]),
		IndProp:  s.Ln[7],
		CodPart:  s.Ln[8],
		TxtCompl: s.Ln[9],
		CodCta:   s.Ln[10],
		VlItemIr: SpedConvert.ConvFloat(s.Ln[11]),
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
	}
	return regH010
}

func CreateRegH010(read iRegH010) RegH010 {
	return read.GetRegH010()
}
