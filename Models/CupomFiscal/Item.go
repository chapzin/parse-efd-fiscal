package cupomfiscal

type CfeItem struct {
	ID        int64  `gorm:"primarykey;autoincrement"`
	IDHeader  int64  `gorm:"notnull,index"`
	UniqueKey string `gorm:"unique,notnull"`
	NItem     int
	CProd     string
	CEAN      string
	XProd     string
	NCM       string
	CFOP      string
	CSTICMS   string
	PICMS     float64
	VICMS     float64
	UCom      string
	QCom      float64
	VUnCom    float64
	VProd     float64
	IndRegra  string
	VItem     float64
	VDesc     float64
	VOutro    float64
}
