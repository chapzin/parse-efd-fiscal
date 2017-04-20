package BlocoC

import (
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
	"time"
)

type RegC490 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4)"`
	CstIcms  string    `gorm:"type:varchar(3)"`
	Cfop     string    `gorm:"type:varchar(4)"`
	AliqIcms float64   `gorm:"type:decimal(6,2)"`
	VlOpr    float64   `gorm:"type:decimal(19,2)"`
	VlBcIcms float64   `gorm:"type:decimal(19,2)"`
	VlIcms   float64   `gorm:"type:decimal(19,2)"`
	CodObs   string    `gorm:"type:varchar(6)"`
	DtIni    time.Time `gorm:"type:date"`
	DtFin    time.Time `gorm:"type:date"`
	Cnpj     string    `gorm:"type:varchar(14)"`
}

// Metodo do nome da tabela
func (RegC490) TableName() string {
	return "reg_C490"
}

// Implementando Interface do Sped RegC490
type RegC490Sped struct {
	Ln     []string
	Reg000 Bloco0.Reg0000
}

type iRegC490 interface {
	GetRegC490() RegC490
}

func (s RegC490Sped) GetRegC490() RegC490 {
	regC490 := RegC490{
		Reg:      s.Ln[1],
		CstIcms:  s.Ln[2],
		Cfop:     s.Ln[3],
		AliqIcms: tools.ConvFloat(s.Ln[4]),
		VlOpr:    tools.ConvFloat(s.Ln[5]),
		VlBcIcms: tools.ConvFloat(s.Ln[6]),
		VlIcms:   tools.ConvFloat(s.Ln[7]),
		CodObs:   s.Ln[8],
		DtIni:    s.Reg000.DtIni,
		DtFin:    s.Reg000.DtFin,
		Cnpj:     s.Reg000.Cnpj,
	}
	return regC490
}

// Cria estrutura populada
func CreateRegC490(read iRegC490) RegC490 {
	return read.GetRegC490()
}
