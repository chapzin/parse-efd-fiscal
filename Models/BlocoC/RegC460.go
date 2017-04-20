package BlocoC

import (
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
	"time"
)

type RegC460 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4)"`
	CodMod   string    `gorm:"type:varchar(2)"`
	CodSit   string    `gorm:"type:varchar(2)"`
	NumDoc   string    `gorm:"type:varchar(9)"`
	DtDoc    time.Time `gorm:"type:date"`
	VlDoc    float64   `gorm:"type:decimal(19,2)"`
	VlPis    float64   `gorm:"type:decimal(19,2)"`
	VlCofins float64   `gorm:"type:decimal(19,2)"`
	CpfCnpj  string    `gorm:"type:varchar(14)"`
	NomAdq   string    `gorm:"type:varchar(60)"`
	DtIni    time.Time `gorm:"type:date"`
	DtFin    time.Time `gorm:"type:date"`
	Cnpj     string    `gorm:"type:varchar(14)"`
}

func (RegC460) TableName() string {
	return "reg_C460"
}

type RegC460Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

type iRegC460 interface {
	GetRegC460() RegC460
}

func (s RegC460Sped) GetRegC460() RegC460 {
	regC460 := RegC460{
		Reg:      s.Ln[1],
		CodMod:   s.Ln[2],
		NumDoc:   s.Ln[3],
		DtDoc:    tools.ConvertData(s.Ln[4]),
		VlDoc:    tools.ConvFloat(s.Ln[5]),
		VlPis:    tools.ConvFloat(s.Ln[6]),
		VlCofins: tools.ConvFloat(s.Ln[7]),
		CpfCnpj:  s.Ln[8],
		NomAdq:   s.Ln[9],
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
	}
	return regC460
}

// Cria estrutura populada
func CreateRegC460(read iRegC460) RegC460 {
	return read.GetRegC460()
}
