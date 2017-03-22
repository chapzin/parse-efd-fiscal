package Bloco0

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/chapzin/parse-efd-fiscal/tools"
)
// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

type Reg0000 struct {
	gorm.Model
	Reg       string        `gorm:"type:varchar(4)"`
	CodVer    string        `gorm:"type:varchar(3)"`
	CodFin    int
	DtIni     time.Time        `gorm:"type:date"`
	DtFin     time.Time        `gorm:"type:date"`
	Nome      string        `gorm:"type:varchar(100)"`
	Cnpj      string        `gorm:"type:varchar(14)"`
	Cpf       string        `gorm:"type:varchar(11)"`
	Uf        string        `gorm:"type:varchar(2)"`
	Ie        string        `gorm:"type:varchar(14)"`
	CodMun    string        `gorm:"type:varchar(7)"`
	Im        string
	Suframa   string       `gorm:"type:varchar(9)"`
	IndPerfil string        `gorm:"type:varchar(1)"`
	IndAtiv   int
}

func (Reg0000) TableName() string {
	return "reg_0000"
}

// Implementando Inteface do Sped Reg0000
type Reg0000Sped struct {
	Ln []string
}

type iReg0000 interface {
	GetReg0000() Reg0000
}

func (s Reg0000Sped) GetReg0000() Reg0000 {
	reg0000 := Reg0000{
		Reg:       s.Ln[1],
		CodVer:    s.Ln[2],
		CodFin:    tools.ConvInt(s.Ln[3]),
		DtIni:     tools.ConvertData(s.Ln[4]),
		DtFin:     tools.ConvertData(s.Ln[5]),
		Nome:      s.Ln[6],
		Cnpj:      s.Ln[7],
		Cpf:       s.Ln[8],
		Uf:        s.Ln[9],
		Ie:        s.Ln[10],
		CodMun:    s.Ln[11],
		Im:        s.Ln[12],
		Suframa:   s.Ln[13],
		IndPerfil: s.Ln[14],
		IndAtiv:   tools.ConvInt(s.Ln[15]),
	}

	return reg0000
}

func CreateReg0000(read iReg0000) Reg0000 {
	return read.GetReg0000()
}
