package BlocoC

import (
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
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

type RegC470Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
	Digito  string
}

type iRegC470 interface {
	GetRegC470() RegC470
}

func (s RegC470Sped) GetRegC470() RegC470 {
	digitoInt := tools.ConvInt(s.Digito)
	regC470 := RegC470{
		Reg:      s.Ln[1],
		CodItem:  tools.AdicionaDigitosCodigo(s.Ln[2], digitoInt),
		Qtd:      tools.ConvFloat(s.Ln[3]),
		QtdCanc:  tools.ConvFloat(s.Ln[4]),
		Unid:     s.Ln[5],
		VlItem:   tools.ConvFloat(s.Ln[6]),
		CstIcms:  s.Ln[7],
		Cfop:     s.Ln[8],
		AliqIcms: tools.ConvFloat(s.Ln[9]),
		VlPis:    tools.ConvFloat(s.Ln[10]),
		VlCofins: tools.ConvFloat(s.Ln[11]),
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
	}
	return regC470
}

// Cria estrutura populada
func CreateRegC470(read iRegC470) RegC470 {
	return read.GetRegC470()
}
