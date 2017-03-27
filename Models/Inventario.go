package Models

import "github.com/jinzhu/gorm"

// Estrutura de inventário
type Inventario struct {
	gorm.Model
	Codigo          string `gorm:"type:varchar(60)"`
	Descricao       string
	Tipo            string  `gorm:"type:varchar(2)"`
	UnidInv         string  `gorm:"type:varchar(6)"`
	Ncm             string  `gorm:"type:varchar(8)"`
	InvInicial      float64 `gorm:"type:decimal(19,3)"`
	VlInvIni        float64 `gorm:"type:decimal(19,3)"`
	Entradas        float64 `gorm:"type:decimal(19,3)"`
	VlTotalEntradas float64 `gorm:"type:decimal(19,3)"`
	VlUnitEnt       float64 `gorm:"type:decimal(19,3)"`
	Saidas          float64 `gorm:"type:decimal(19,3)"`
	VlTotalSaidas   float64 `gorm:"type:decimal(19,3)"`
	VlUnitSai       float64 `gorm:"type:decimal(19,3)"`
	Margem          float64 `gorm:"type:decimal(19,3)"`
	InvFinal        float64 `gorm:"type:decimal(19,3)"`
	VlInvFin        float64 `gorm:"type:decimal(19,3)"`
	Diferencas      float64 `gorm:"type:decimal(19,3)"`
	SugInvInicial   float64 `gorm:"type:decimal(19,3)"`
	SugVlInvInicial float64 `gorm:"type:decimal(19,3)"`
	SugInvFinal     float64 `gorm:"type:decimal(19,3)"`
	SugVlInvFinal   float64 `gorm:"type:decimal(19,3)"`
	Ano             int
}
