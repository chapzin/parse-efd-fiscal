package BlocoH

import (
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
	"time"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016
// Estrutura do registro H010
type RegH010 struct {
	gorm.Model
	Reg      string  `gorm:"type:varchar(4)"`
	CodItem  string  `gorm:"type:varchar(60)"`
	Unid     string  `gorm:"type:varchar(6)"`
	Qtd      float64 `gorm:"type:decimal(19,3)"`
	VlUnit   float64 `gorm:"type:decimal(19,6)"`
	VlItem   float64 `gorm:"type:decimal(19,2)"`
	IndProp  string  `gorm:"type:varchar(1)"`
	CodPart  string  `gorm:"type:varchar(60)"`
	TxtCompl string
	CodCta   string
	VlItemIr float64   `gorm:"type:decimal(19,2)"`
	DtInv    time.Time `gorm:"type:date"`
	DtIni    time.Time `gorm:"type:date"`
	DtFin    time.Time `gorm:"type:date"`
	Cnpj     string    `gorm:"type:varchar(14)"`
}

// Metodo define nome da tabela no banco de dados
func (RegH010) TableName() string {
	return "reg_h010"
}

// Implementando Interface do Sped RegH010
type RegH010Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
	RegH005 RegH005
	Digito  string
}

type iRegH010 interface {
	GetRegH010() RegH010
}

// Meotodo que popula o H010 do sped fiscal
func (s RegH010Sped) GetRegH010() RegH010 {
	digitoInt := tools.ConvInt(s.Digito)
	regH010 := RegH010{
		Reg:      s.Ln[1],
		CodItem:  tools.AdicionaDigitosCodigo(s.Ln[2], digitoInt),
		Unid:     s.Ln[3],
		Qtd:      tools.ConvFloat(s.Ln[4]),
		VlUnit:   tools.ConvFloat(s.Ln[5]),
		VlItem:   tools.ConvFloat(s.Ln[6]),
		IndProp:  s.Ln[7],
		CodPart:  s.Ln[8],
		TxtCompl: s.Ln[9],
		CodCta:   s.Ln[10],
		VlItemIr: tools.ConvFloat(s.Ln[11]),
		DtInv:    s.RegH005.DtInv,
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
	}
	return regH010
}

// Metodo que cria o registro populado generico
func CreateRegH010(read iRegH010) RegH010 {
	return read.GetRegH010()
}
