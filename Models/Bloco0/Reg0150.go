package Bloco0

import (
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
	"time"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

type Reg0150 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4)"`
	CodPart  string    `gorm:"type:varchar(60);"`
	Nome     string    `gorm:"type:varchar(100)"`
	CodPais  string    `gorm:"type:varchar(5)"`
	Cnpj     string    `gorm:"type:varchar(15)"`
	Cpf      string    `gorm:"type:varchar(11)"`
	Ie       string    `gorm:"type:varchar(14)"`
	CodMun   string    `gorm:"type:varchar(7)"`
	Suframa  string    `gorm:"type:varchar(9)"`
	Endereco string    `gorm:"type:varchar(60)"`
	Num      string    `gorm:"type:varchar(10)"`
	Compl    string    `gorm:"type:varchar(60)"`
	Bairro   string    `gorm:"type:varchar(60)"`
	DtIni    time.Time `gorm:"type:date"`
	DtFin    time.Time `gorm:"type:date"`
	CnpjSped string    `gorm:"type:varchar(14)"`
}

func (Reg0150) TableName() string {
	return "reg_0150"
}

type iReg0150 interface {
	GetReg0150() Reg0150
}

type Reg0150Sped struct {
	Ln      []string
	Reg0000 Reg0000
}

func (s Reg0150Sped) GetReg0150() Reg0150 {
	reg0150 := Reg0150{
		Reg:      s.Ln[1],
		CodPart:  s.Ln[2],
		Nome:     s.Ln[3],
		CodPais:  s.Ln[4],
		Cnpj:     s.Ln[5],
		Cpf:      s.Ln[6],
		Ie:       s.Ln[7],
		CodMun:   s.Ln[8],
		Suframa:  s.Ln[9],
		Endereco: s.Ln[10],
		Num:      s.Ln[11],
		Compl:    s.Ln[12],
		Bairro:   s.Ln[13],
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		CnpjSped: s.Reg0000.Cnpj,
	}
	return reg0150
}

type Reg0150Xml struct {
	Reader func(pathTag string, tag string) string
}

func (x Reg0150Xml) GetReg0150() Reg0150 {
	reg0150 := Reg0150{
		Reg:      "0150",
		CodPart:  x.Reader("dest", "CNPJ"), //SpedConvert.DataXml(x.Cliente, "CNPJ"),
		Nome:     x.Reader("dest", "xNome"),
		CodPais:  "1058", //Importante separar esse numero em uma constante em outro arquivo, talvez. Assim, vai ser constante geral para o projeto
		Cnpj:     x.Reader("dest", "CNPJ"),
		Cpf:      x.Reader("dest", "CNPJ"),
		Ie:       x.Reader("dest", "IE"),
		CodMun:   x.Reader("enderDest", "cMun"),
		Suframa:  "",
		Endereco: x.Reader("enderDest", "xLgr"),
		Num:      x.Reader("enderDest", "nro"),
		Compl:    x.Reader("enderDest", "xCpl"),
		Bairro:   x.Reader("enderDest", "xBairro"),
		DtIni:    tools.ConvertDataNull(),
		DtFin:    tools.ConvertDataNull(),
		CnpjSped: "Insert Xml",
	}
	return reg0150
}

func CreateReg0150(read iReg0150) Reg0150 {
	return read.GetReg0150()
}
