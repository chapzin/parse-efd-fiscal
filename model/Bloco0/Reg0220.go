package Bloco0

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
)


type iReg0220 interface {
	GetReg() string
	GetUnidConv() string
	GetFatConv() float64
	GetDtIni() time.Time
	GetDtFin() time.Time
	GetCnpj() string
}


type Reg0220 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	UnidConv string		`gorm:"type:varchar(6)"`
	FatConv float64		`gorm:"type:decimal(12,6)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`
}

func (Reg0220) TableName() string {
	return "reg_0220"
}


// Implementando Interface do Sped Reg0220
type Reg0220Sped struct {
	Ln []string
	Reg0000 Reg0000
}

func (s Reg0220Sped) GetReg() string {
	return s.Ln[1]
}

func (s Reg0220Sped) GetUnidConv() string {
	return  s.Ln[2]
}

func (s Reg0220Sped) GetFatConv() float64 {
	return SpedConvert.ConvFloat(s.Ln[3])
}

func (s Reg0220Sped) GetDtIni() time.Time {
	return  s.Reg0000.DtIni
}

func (s Reg0220Sped) GetDtFin() time.Time  {
	return s.Reg0000.DtFin
}

func (s Reg0220Sped) GetCnpj() string  {
	return s.Reg0000.Cnpj
}

func CreateReg0220 ( read iReg0220) Reg0220 {
	reg0220 := Reg0220{
		Reg:		read.GetReg(),
		UnidConv:	read.GetUnidConv(),
		FatConv:	read.GetFatConv(),
		DtIni:		read.GetDtIni(),
		DtFin:		read.GetDtFin(),
		Cnpj:		read.GetCnpj(),
	}
	return reg0220
}






