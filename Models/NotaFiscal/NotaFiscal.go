package NotaFiscal

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Estrutura da nota fiscal eletronica
type NotaFiscal struct {
	gorm.Model
	NNF            string
	ChNFe          string `gorm:"unique"`
	NatOp          string
	IndPag         string
	Mod            string
	Serie          string
	DEmi           time.Time
	TpNF           string
	TpImp          string
	TpEmis         string
	CDV            string
	TpAmb          string
	FinNFe         string
	ProcEmi        string
	Emitente       Emitente
	EmitenteID     int
	Destinatario   Destinatario
	DestinatarioID int
	Itens          []Item
}
