package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/chapzin/parse-efd-fiscal/model/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
)

type RegC170 struct {
	gorm.Model
	Reg           string                `gorm:"type:varchar(4)"`
	NumItem       string                `gorm:"type:varchar(3)"`
	CodItem       string                `gorm:"type:varchar(60)"`
	DescrCompl    string
	Qtd           float64                `gorm:"type:decimal(19,5)"`
	Unid          string                `gorm:"type:varchar(6)"`
	VlItem        float64                `gorm:"type:decimal(19,2)"`
	VlDesc        float64                `gorm:"type:decimal(19,2)"`
	IndMov        string                `gorm:"type:varchar(1)"`
	CstIcms       string                `gorm:"type:varchar(3)"`
	Cfop          string                `gorm:"type:varchar(4)"`
	CodNat        string                `gorm:"type:varchar(10)"`
	VlBcIcms      float64        `gorm:"type:decimal(19,2)"`
	AliqIcms      float64        `gorm:"type:decimal(6,2)"`
	VlIcms        float64                `gorm:"type:decimal(19,2)"`
	VlBcIcmsSt    float64        `gorm:"type:decimal(19,2)"`
	AliqSt        float64                `gorm:"type:decimal(19,2)"`
	VlIcmsSt      float64        `gorm:"type:decimal(19,2)"`
	IndApur       string                `gorm:"type:varchar(1)"`
	CstIpi        string                `gorm:"type:varchar(2)"`
	CodEnq        string                `gorm:"type:varchar(3)"`
	VlBcIpi       float64                `gorm:"type:decimal(19,2)"`
	AliqIpi       float64                `gorm:"type:decimal(6,2)"`
	VlIpi         float64                `gorm:"type:decimal(19,2)"`
	CstPis        string                `gorm:"type:varchar(2)"`
	VlBcPis       float64                `gorm:"type:decimal(19,2)"`
	AliqPis01     float64        `gorm:"type:decimal(8,4)"`
	QuantBcPis    float64        `gorm:"type:decimal(19,3)"`
	AliqPis02     float64        `gorm:"type:decimal(8,4)"`
	VlPis         float64                `gorm:"type:decimal(19,2)"`
	CstCofins     string        `gorm:"type:varchar(2)"`
	VlBcCofins    float64        `gorm:"type:decimal(19,2)"`
	AliqCofins01  float64        `gorm:"type:decimal(8,4)"`
	QuantBcCofins float64        `gorm:"type:decimal(19,3)"`
	AliqCofins02  float64        `gorm:"type:decimal(8,4)"`
	VlCofins      float64        `gorm:"type:decimal(19,2)"`
	CodCta        string
	EntradaSaida  string        `gorm:"type:varchar(1)"` // Se for entrada 0, se for saida 1
	NumDoc        string                `gorm:"type:varchar(9)"`
	DtIni         time.Time        `gorm:"type:date"`
	DtFin         time.Time        `gorm:"type:date"`
	Cnpj          string                `gorm:"type:varchar(14)"`
}

func (RegC170) TableName() string {
	return "reg_c170"
}

// Implementando Interface do SpedRegC170
type RegC170Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
	RegC100 RegC100
}

type iRegC170 interface {
	GetRegC170() RegC170
}

func (s RegC170Sped) GetRegC170() RegC170 {
	regC170 := RegC170{
		Reg:           s.Ln[1],
		NumItem:       s.Ln[2],
		CodItem:       s.Ln[3],
		DescrCompl:    s.Ln[4],
		Qtd:           SpedConvert.ConvFloat(s.Ln[5]),
		Unid:          s.Ln[6],
		VlItem:        SpedConvert.ConvFloat(s.Ln[7]),
		VlDesc:        SpedConvert.ConvFloat(s.Ln[8]),
		IndMov:        s.Ln[9],
		CstIcms:       s.Ln[10],
		Cfop:          s.Ln[11],
		CodNat:        s.Ln[12],
		VlBcIcms:      SpedConvert.ConvFloat(s.Ln[13]),
		AliqIcms:      SpedConvert.ConvFloat(s.Ln[14]),
		VlIcms:        SpedConvert.ConvFloat(s.Ln[15]),
		VlBcIcmsSt:    SpedConvert.ConvFloat(s.Ln[16]),
		AliqSt:        SpedConvert.ConvFloat(s.Ln[17]),
		VlIcmsSt:      SpedConvert.ConvFloat(s.Ln[18]),
		IndApur:       s.Ln[19],
		CstIpi:        s.Ln[20],
		CodEnq:        s.Ln[21],
		VlBcIpi:       SpedConvert.ConvFloat(s.Ln[22]),
		AliqIpi:       SpedConvert.ConvFloat(s.Ln[23]),
		VlIpi:         SpedConvert.ConvFloat(s.Ln[24]),
		CstPis:        s.Ln[25],
		VlBcPis:       SpedConvert.ConvFloat(s.Ln[26]),
		AliqPis01:     SpedConvert.ConvFloat(s.Ln[27]),
		QuantBcPis:    SpedConvert.ConvFloat(s.Ln[28]),
		AliqPis02:     SpedConvert.ConvFloat(s.Ln[29]),
		VlPis:         SpedConvert.ConvFloat(s.Ln[30]),
		CstCofins:     s.Ln[31],
		VlBcCofins:    SpedConvert.ConvFloat(s.Ln[32]),
		AliqCofins01:  SpedConvert.ConvFloat(s.Ln[33]),
		QuantBcCofins: SpedConvert.ConvFloat(s.Ln[34]),
		AliqCofins02:  SpedConvert.ConvFloat(s.Ln[35]),
		VlCofins:      SpedConvert.ConvFloat(s.Ln[36]),
		CodCta:        s.Ln[37],
		EntradaSaida:  s.RegC100.IndOper,
		NumDoc:        s.RegC100.NumDoc,
		DtIni:         s.Reg0000.DtIni,
		DtFin:         s.Reg0000.DtFin,
		Cnpj:          s.Reg0000.Cnpj,
	}
	return regC170
}

func CreateRegC170(read iRegC170) RegC170 {
	return read.GetRegC170()
}
