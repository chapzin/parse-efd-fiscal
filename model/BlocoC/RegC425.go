package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
	"database/sql"
)

type RegC425 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	CodItem string		`gorm:"type:varchar(60)"`
	Qtd sql.NullFloat64		`gorm:"type:decimal(19,3)"`
	Unid string		`gorm:"type:varchar(6)"`
	VlItem sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlPis sql.NullFloat64		`gorm:"type:decimal(19,2)"`
	VlCofins sql.NullFloat64	`gorm:"type:decimal(19,2)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	Cnpj string		`gorm:"type:varchar(14)"`

}

func (RegC425) TableName() string {
	return "reg_c425"
}