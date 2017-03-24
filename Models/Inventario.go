package Model

import "github.com/jinzhu/gorm"

type Inventario struct {
	gorm.Model
	Codigo          string		`gorm:"type:varchar(60)"`
	Descricao       string
	Tipo            string		`gorm:"type:varchar(2)"`
	UnidInv         string		`gorm:"type:varchar(6)"`
	InvInicial      float64		`gorm:"type:decimal(19,3)"`
	VlInvIni	float64		`gorm:"type:decimal(19,3)"`
	Entradas        float64		`gorm:"type:decimal(19,3)"`
	VlTotalEntradas float64		`gorm:"type:decimal(19,3)"`
	VlUnitEnt	float64		`gorm:"type:decimal(19,3)"`
	Saidas          float64		`gorm:"type:decimal(19,3)"`
	VlTotalSaidas   float64		`gorm:"type:decimal(19,3)"`
	VlUnitSai	float64		`gorm:"type:decimal(19,3)"`
	InvFinal        float64		`gorm:"type:decimal(19,3)"`
	VlInvFin	float64		`gorm:"type:decimal(19,3)"`
	Diferencas      float64		`gorm:"type:decimal(19,3)"`
	EstoqueIni	float64		`gorm:"type:decimal(19,3)"`
	VlTotIni	float64		`gorm:"type:decimal(19,3)"`
	EstoqueFin	float64		`gorm:"type:decimal(19,3)"`
	VlTotFin	float64		`gorm:"type:decimal(19,3)"`
	Ano             int
}
