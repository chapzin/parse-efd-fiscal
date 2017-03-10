package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Emitente struct {
	gorm.Model
	CNPJ string
	XNome string
	XLgr string
	Nro string
	XCpl string
	XBairro string
	CMun string
	XMun string
	Uf string
	Cep string
	CPais string
	XPais string
	Fone string
	Ie string
}

type Destinatario struct {
	gorm.Model
	CNPJ string
	XNome string
	XLgr string
	Nro string
	XCpl string
	XBairro string
	CMun string
	XMun string
	Uf string
	Cep string
	CPais string
	XPais string
	Fone string
	Ie string
}

type Item struct {
	gorm.Model
	Codigo string
	Ean string
	Descricao string
	Ncm string
	Cfop string
	Unid string
	Qtd float64
	VUnit float64
	VTotal float64
	DtEmit time.Time
	NotaFiscalID uint
}

type NotaFiscal struct {
	gorm.Model
	NNF string
	ChNFe string
	NatOp string
	IndPag string
	Mod string
	Serie string
	DEmi time.Time
	TpNF string
	TpImp string
	TpEmis string
	CDV string
	TpAmb string
	FinNFe string
	ProcEmi string
	Emitente Emitente
	EmitenteID int
	Destinatario Destinatario
	DestinatarioID int
	Itens []Item
}
