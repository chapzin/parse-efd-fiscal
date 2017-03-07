package Bloco0

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
)



type Reg0190 struct {
	gorm.Model
	Reg string	`gorm:"type:varchar(4)"`
	Unid string	`gorm:"type:varchar(6);unique_index"`
	Descr string
	DtIni time.Time `gorm:"type:date"`
	DtFin time.Time `gorm:"type:date"`
	Cnpj string	`gorm:"type:varchar(14)"`
}

func (Reg0190) TableName() string {
	return "reg_0190"
}

// Implementando Interface do Sped Reg0190
type Reg0190Sped struct {
	Ln []string
	Reg0000 Reg0000
}

func (s Reg0190Sped) GetReg() string {
	return s.Ln[1]
}

func (s Reg0190Sped) GetUnid() string {
	return s.Ln[2]
}

func (s Reg0190Sped) GetDescr() string {
	return  s.Ln[3]
}

func (s Reg0190Sped) GetDtIni() time.Time {
	return s.Reg0000.DtIni
}

func (s Reg0190Sped) GetDtFin() time.Time {
	return s.Reg0000.DtFin
}

func (s Reg0190Sped) GetCnpj() string  {
	return s.Reg0000.Cnpj
}

type Reg0190Xml struct {
	Data string
}

func (x Reg0190Xml ) GetReg() string {
	return "0190"
}

func (x Reg0190Xml) GetUnid() string  {
	return x.Data
}

func (x Reg0190Xml) GetDescr() string {
	return "Importado Xml"
}

func (x Reg0190Xml) GetDtIni() time.Time {
	return SpedConvert.ConvertDataNull()
}

func (x Reg0190Xml) GetDtFin() time.Time {
	return SpedConvert.ConvertDataNull()
}

func (x Reg0190Xml) GetCnpj() string {
	return ""
}


type iReg0190 interface {
	GetReg() string
	GetUnid() string
	GetDescr() string
	GetDtIni() time.Time
	GetDtFin() time.Time
	GetCnpj() string

}

func CreateReg0190 (read iReg0190) Reg0190 {
	reg0190 := Reg0190{
		Reg:		read.GetReg(),
		Unid:		read.GetUnid(),
		Descr:		read.GetDescr(),
		DtIni:		read.GetDtIni(),
		DtFin:		read.GetDtFin(),
		Cnpj:		read.GetCnpj(),
	}
	return reg0190
}


