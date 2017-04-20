package BlocoC

import (
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
	"time"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016
// Estrutura do registro C100
type RegC100 struct {
	gorm.Model
	Reg        string    `gorm:"type:varchar(4)"`
	IndOper    string    `gorm:"type:varchar(1)"`
	IndEmit    string    `gorm:"type:varchar(1)"`
	CodPart    string    `gorm:"type:varchar(60)"`
	CodMod     string    `gorm:"type:varchar(2)"`
	CodSit     string    `gorm:"type:varchar(2)"`
	Ser        string    `gorm:"type:varchar(3)"`
	NumDoc     string    `gorm:"type:varchar(9)"`
	ChvNfe     string    `gorm:"type:varchar(44);unique_index"`
	DtDoc      time.Time `gorm:"type:date"`
	DtES       time.Time `gorm:"type:date"`
	VlDoc      float64   `gorm:"type:decimal(19,2)"`
	IndPgto    string    `gorm:"type:varchar(1)"`
	VlDesc     float64   `gorm:"type:decimal(19,2)"`
	VlAbatNt   float64   `gorm:"type:decimal(19,2)"`
	VlMerc     float64   `gorm:"type:decimal(19,2)"`
	IndFrt     string    `gorm:"type:varchar(1)"`
	VlFrt      float64   `gorm:"type:decimal(19,2)"`
	VlSeg      float64   `gorm:"type:decimal(19,2)"`
	VlOutDa    float64   `gorm:"type:decimal(19,2)"`
	VlBcIcms   float64   `gorm:"type:decimal(19,2)"`
	VlIcms     float64   `gorm:"type:decimal(19,2)"`
	VlBcIcmsSt float64   `gorm:"type:decimal(19,2)"`
	VlIcmsSt   float64   `gorm:"type:decimal(19,2)"`
	VlIpi      float64   `gorm:"type:decimal(19,2)"`
	VlPis      float64   `gorm:"type:decimal(19,2)"`
	VlCofins   float64   `gorm:"type:decimal(19,2)"`
	VlPisSt    float64   `gorm:"type:decimal(19,2)"`
	VlCofinsSt float64   `gorm:"type:decimal(19,2)"`
	DtIni      time.Time `gorm:"type:date"`
	DtFin      time.Time `gorm:"type:date"`
	Cnpj       string    `gorm:"type:varchar(14)"`
}

// Metodo do nome da tabela
func (RegC100) TableName() string {
	return "reg_c100"
}

// Implementando Interface do Sped RegC100
type RegC100Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

type iRegC100 interface {
	GetRegC100() RegC100
}

// Metodo popula strutura com conteudo do sped fiscal
func (s RegC100Sped) GetRegC100() RegC100 {
	regC100 := RegC100{
		Reg:        s.Ln[1],
		IndOper:    s.Ln[2],
		IndEmit:    s.Ln[3],
		CodPart:    s.Ln[4],
		CodMod:     s.Ln[5],
		CodSit:     s.Ln[6],
		Ser:        s.Ln[7],
		NumDoc:     s.Ln[8],
		ChvNfe:     s.Ln[9],
		DtDoc:      tools.ConvertData(s.Ln[10]),
		DtES:       tools.ConvertData(s.Ln[11]),
		VlDoc:      tools.ConvFloat(s.Ln[12]),
		IndPgto:    s.Ln[13],
		VlDesc:     tools.ConvFloat(s.Ln[14]),
		VlAbatNt:   tools.ConvFloat(s.Ln[15]),
		VlMerc:     tools.ConvFloat(s.Ln[16]),
		IndFrt:     s.Ln[17],
		VlFrt:      tools.ConvFloat(s.Ln[18]),
		VlSeg:      tools.ConvFloat(s.Ln[19]),
		VlOutDa:    tools.ConvFloat(s.Ln[20]),
		VlBcIcms:   tools.ConvFloat(s.Ln[21]),
		VlIcms:     tools.ConvFloat(s.Ln[22]),
		VlBcIcmsSt: tools.ConvFloat(s.Ln[23]),
		VlIcmsSt:   tools.ConvFloat(s.Ln[24]),
		VlIpi:      tools.ConvFloat(s.Ln[25]),
		VlPis:      tools.ConvFloat(s.Ln[26]),
		VlCofins:   tools.ConvFloat(s.Ln[27]),
		VlPisSt:    tools.ConvFloat(s.Ln[29]),
		VlCofinsSt: tools.ConvFloat(s.Ln[30]),
		DtIni:      s.Reg0000.DtIni,
		DtFin:      s.Reg0000.DtFin,
		Cnpj:       s.Reg0000.Cnpj,
	}
	return regC100
}

// Cria estrutura populada
func CreateRegC100(read iRegC100) RegC100 {
	return read.GetRegC100()
}
