package Models

import "github.com/jinzhu/gorm"

// Estrutura de inventário
type Inventario struct {
	gorm.Model
	Codigo       string `gorm:"type:varchar(60);unique_index"`
	Descricao    string
	Tipo         string  `gorm:"type:varchar(2)"`
	UnidInv      string  `gorm:"type:varchar(6)"`
	Ncm          string  `gorm:"type:varchar(8)"`
	InvFinalAno1 float64 `gorm:"type:decimal(19,3)"`
	VlInvAno1    float64 `gorm:"type:decimal(19,3)"`
	// Movimentação Ano 2
	EntradasAno2        float64 `gorm:"type:decimal(19,3)"`
	VlTotalEntradasAno2 float64 `gorm:"type:decimal(19,3)"`
	VlUnitEntAno2       float64 `gorm:"type:decimal(19,3)"`
	SaidasAno2          float64 `gorm:"type:decimal(19,3)"`
	VlTotalSaidasAno2   float64 `gorm:"type:decimal(19,3)"`
	VlUnitSaiAno2       float64 `gorm:"type:decimal(19,3)"`
	MargemAno2          float64 `gorm:"type:decimal(19,3)"`
	InvFinalAno2        float64 `gorm:"type:decimal(19,3)"`
	VlInvAno2           float64 `gorm:"type:decimal(19,3)"`
	DiferencasAno2      float64 `gorm:"type:decimal(19,3)"`
	// Movimentacao Ano 3
	EntradasAno3        float64 `gorm:"type:decimal(19,3)"`
	VlTotalEntradasAno3 float64 `gorm:"type:decimal(19,3)"`
	VlUnitEntAno3       float64 `gorm:"type:decimal(19,3)"`
	SaidasAno3          float64 `gorm:"type:decimal(19,3)"`
	VlTotalSaidasAno3   float64 `gorm:"type:decimal(19,3)"`
	VlUnitSaiAno3       float64 `gorm:"type:decimal(19,3)"`
	MargemAno3          float64 `gorm:"type:decimal(19,3)"`
	InvFinalAno3        float64 `gorm:"type:decimal(19,3)"`
	VlInvAno3           float64 `gorm:"type:decimal(19,3)"`
	DiferencasAno3      float64 `gorm:"type:decimal(19,3)"`
	// Movimentacao Ano 4
	EntradasAno4        float64 `gorm:"type:decimal(19,3)"`
	VlTotalEntradasAno4 float64 `gorm:"type:decimal(19,3)"`
	VlUnitEntAno4       float64 `gorm:"type:decimal(19,3)"`
	SaidasAno4          float64 `gorm:"type:decimal(19,3)"`
	VlTotalSaidasAno4   float64 `gorm:"type:decimal(19,3)"`
	VlUnitSaiAno4       float64 `gorm:"type:decimal(19,3)"`
	MargemAno4          float64 `gorm:"type:decimal(19,3)"`
	InvFinalAno4        float64 `gorm:"type:decimal(19,3)"`
	VlInvAno4           float64 `gorm:"type:decimal(19,3)"`
	DiferencasAno4      float64 `gorm:"type:decimal(19,3)"`
	// Movimentacao Ano 5
	EntradasAno5        float64 `gorm:"type:decimal(19,3)"`
	VlTotalEntradasAno5 float64 `gorm:"type:decimal(19,3)"`
	VlUnitEntAno5       float64 `gorm:"type:decimal(19,3)"`
	SaidasAno5          float64 `gorm:"type:decimal(19,3)"`
	VlTotalSaidasAno5   float64 `gorm:"type:decimal(19,3)"`
	VlUnitSaiAno5       float64 `gorm:"type:decimal(19,3)"`
	MargemAno5          float64 `gorm:"type:decimal(19,3)"`
	InvFinalAno5        float64 `gorm:"type:decimal(19,3)"`
	VlInvAno5           float64 `gorm:"type:decimal(19,3)"`
	DiferencasAno5      float64 `gorm:"type:decimal(19,3)"`
	// Movimentacao Ano 6
	EntradasAno6        float64 `gorm:"type:decimal(19,3)"`
	VlTotalEntradasAno6 float64 `gorm:"type:decimal(19,3)"`
	VlUnitEntAno6       float64 `gorm:"type:decimal(19,3)"`
	SaidasAno6          float64 `gorm:"type:decimal(19,3)"`
	VlTotalSaidasAno6   float64 `gorm:"type:decimal(19,3)"`
	VlUnitSaiAno6       float64 `gorm:"type:decimal(19,3)"`
	MargemAno6          float64 `gorm:"type:decimal(19,3)"`
	InvFinalAno6        float64 `gorm:"type:decimal(19,3)"`
	VlInvAno6           float64 `gorm:"type:decimal(19,3)"`
	DiferencasAno6      float64 `gorm:"type:decimal(19,3)"`
	// Sugestão inventarios
	// Exemplo Inv Final 2011
	SugInvAno1   float64 `gorm:"type:decimal(19,3)"`
	SugVlInvAno1 float64 `gorm:"type:decimal(19,3)"`
	// Exemplo Inv Final 2012
	SugInvAno2   float64 `gorm:"type:decimal(19,3)"`
	SugVlInvAno2 float64 `gorm:"type:decimal(19,3)"`
	// Exemplo Inv Final 2013
	SugInvAno3   float64 `gorm:"type:decimal(19,3)"`
	SugVlInvAno3 float64 `gorm:"type:decimal(19,3)"`
	// Exemplo Inv Final 2014
	SugInvAno4   float64 `gorm:"type:decimal(19,3)"`
	SugVlInvAno4 float64 `gorm:"type:decimal(19,3)"`
	// Exemplo Inv Final 2015
	SugInvAno5   float64 `gorm:"type:decimal(19,3)"`
	SugVlInvAno5 float64 `gorm:"type:decimal(19,3)"`
	// Exemplo Inv Final 2016
	SugInvAno6   float64 `gorm:"type:decimal(19,3)"`
	SugVlInvAno6 float64 `gorm:"type:decimal(19,3)"`
}
