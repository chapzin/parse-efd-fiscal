package BlocoH

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/chapzin/parse-efd-fiscal/model/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
)

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
		DtInv:  SpedConvert.ConvertData(s.Ln[2]),
		VlInv:  SpedConvert.ConvFloat(s.Ln[3]),
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
