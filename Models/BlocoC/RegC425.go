package BlocoC

import (
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
	"time"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

type RegC425 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4)"`
	CodItem  string    `gorm:"type:varchar(60)"`
	Qtd      float64   `gorm:"type:decimal(19,3)"`
	Unid     string    `gorm:"type:varchar(6)"`
	VlItem   float64   `gorm:"type:decimal(19,2)"`
	VlPis    float64   `gorm:"type:decimal(19,2)"`
	VlCofins float64   `gorm:"type:decimal(19,2)"`
	DtIni    time.Time `gorm:"type:date"`
	DtFin    time.Time `gorm:"type:date"`
	Cnpj     string    `gorm:"type:varchar(14)"`
}

func (RegC425) TableName() string {
	return "reg_c425"
}

// Implementando Interface do Sped RegC425
type RegC425Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
	Digito  string
}

type iRegC425 interface {
	GetRegC425() RegC425
}

func (s RegC425Sped) GetRegC425() RegC425 {
	digitoInt := tools.ConvInt(s.Digito)
	regC425 := RegC425{
		Reg:      s.Ln[1],
		CodItem:  tools.AdicionaDigitosCodigo(s.Ln[2], digitoInt),
		Qtd:      tools.ConvFloat(s.Ln[3]),
		Unid:     s.Ln[4],
		VlItem:   tools.ConvFloat(s.Ln[5]),
		VlPis:    tools.ConvFloat(s.Ln[6]),
		VlCofins: tools.ConvFloat(s.Ln[7]),
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
	}
	return regC425
}

func CreateRegC425(read iRegC425) RegC425 {
	return read.GetRegC425()
}
