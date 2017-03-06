package Bloco0

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
)


type Reg0000 struct {
	gorm.Model
	Reg string	`gorm:"type:varchar(4)"`
	CodVer string	`gorm:"type:varchar(3)"`
	CodFin int
	DtIni time.Time	`gorm:"type:date"`
	DtFin time.Time	`gorm:"type:date"`
	Nome string	`gorm:"type:varchar(100)"`
	Cnpj string	`gorm:"type:varchar(14)"`
	Cpf string	`gorm:"type:varchar(11)"`
	Uf string	`gorm:"type:varchar(2)"`
	Ie string	`gorm:"type:varchar(14)"`
	CodMun string	`gorm:"type:varchar(7)"`
	Im string
	Suframa string	`gorm:"type:varchar(9)"`
	IndPerfil string	`gorm:"type:varchar(1)"`
	IndAtiv int
}

func (Reg0000) TableName() string {
	return "reg_0000"
}

// Implementando Inteface do Sped Reg0000
type Reg0000Sped struct {
	Ln []string
}

func (s Reg0000Sped) GetReg() string{
	return s.Ln[1]
}

func (s Reg0000Sped) GetCodVer() string {
	return s.Ln[2]
}

func (s Reg0000Sped) GetCodFin() int {
	return SpedConvert.ConvInt(s.Ln[3])
}

func (s Reg0000Sped) GetDtIni() time.Time {
	return SpedConvert.ConvertData(s.Ln[4])
}

func (s Reg0000Sped) GetDtFin() time.Time {
	return SpedConvert.ConvertData(s.Ln[5])
}

func (s Reg0000Sped) GetNome() string {
	return s.Ln[6]
}

func (s Reg0000Sped) GetCnpj() string {
	return  s.Ln[7]
}

func (s Reg0000Sped) GetCpf() string {
	return s.Ln[8]
}

func (s Reg0000Sped) GetUf() string  {
	return s.Ln[9]
}

func (s Reg0000Sped) GetIe() string  {
	return s.Ln[10]
}

func (s Reg0000Sped) GetCodMun() string  {
	return s.Ln[11]
}

func (s Reg0000Sped) GetIm() string {
	return s.Ln[12]
}

func (s Reg0000Sped) GetSuframa() string  {
	return s.Ln[13]
}

func (s Reg0000Sped) GetIndPerfil() string  {
	return s.Ln[14]
}

func (s Reg0000Sped) GetIndAtiv() int  {
	return SpedConvert.ConvInt(s.Ln[15])
}

func CreateReg0000 (read iReg0000) Reg0000 {
	reg0000 := Reg0000{
		Reg:		read.GetReg(),
		CodVer:		read.GetCodVer(),
		CodFin:		read.GetCodFin(),
		DtIni:		read.GetDtIni(),
		DtFin:		read.GetDtFin(),
		Nome:		read.GetNome(),
		Cnpj:		read.GetCnpj(),
		Cpf:		read.GetCpf(),
		Uf:		read.GetUf(),
		Ie:		read.GetIe(),
		CodMun:		read.GetCodMun(),
		Im:		read.GetIm(),
		Suframa:	read.GetSuframa(),
		IndPerfil:	read.GetIndPerfil(),
	}
	return reg0000
}

type iReg0000 interface {
	GetReg() string
	GetCodVer() string
	GetCodFin() int
	GetDtIni() time.Time
	GetDtFin() time.Time
	GetNome() string
	GetCnpj() string
	GetCpf() string
	GetUf() string
	GetIe() string
	GetCodMun() string
	GetIm() string
	GetSuframa() string
	GetIndPerfil() string
	GetIndAtiv() int
}
