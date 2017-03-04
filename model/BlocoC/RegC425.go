package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RegC425 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	CodItem string		`gorm:"type:varchar(60)"`
	Qtd float64		`gorm:"type:decimal(19,3)"`
	Unid string		`gorm:"type:varchar(6)"`
	VlItem float64		`gorm:"type:decimal(19,2)"`
	VlPis float64		`gorm:"type:decimal(19,2)"`
	VlCofins float64	`gorm:"type:decimal(19,2)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`

}

func (RegC425) TableName() string {
	return "reg_c425"
}