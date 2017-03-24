package NotaFiscal

import "github.com/jinzhu/gorm"

type Destinatario struct {
	gorm.Model
	CNPJ    string
	XNome   string
	XLgr    string
	Nro     string
	XCpl    string
	XBairro string
	CMun    string
	XMun    string
	Uf      string
	Cep     string
	CPais   string
	XPais   string
	Fone    string
	Ie      string
}
