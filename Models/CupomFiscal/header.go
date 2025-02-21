package cupomfiscal

import (
	"time"
)

type CfeHeader struct {
	ID                  int64
	Numero              int64
	Chave               string `gorm:"unique"`
	CUf                 string
	Mod                 string
	NserieSAT           string
	NCFe                string
	DEmi                time.Time
	CDV                 string
	TpAmb               string
	CNPJ                string
	NumeroCaixa         string
	CnpjEmitente        string
	Emitente            string
	XFant               string
	CnpjCpfDestinatario string
	Destinatario        string
	VlICMS              float64
	VlProd              float64
	VlDesc              float64
	VlPIS               float64
	VlCOFINS            float64
	VlPISST             float64
	VlCOFINSST          float64
	VlOutro             float64
	VlCFe               float64
	VlCFeLei12741       float64
	CMP                 string
	VMP                 float64
	VTroco              float64
}
