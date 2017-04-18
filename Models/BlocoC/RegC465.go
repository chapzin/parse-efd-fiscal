package BlocoC

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RegC465 struct {
	gorm.Model
	Reg    string    `gorm:"type:varchar(4)"`
	ChvCfe string    `gorm:"type:varchar(44)"`
	NumCcF string    `gorm:"type:varchar(9)"`
	DtIni  time.Time `gorm:"type:date"`
	DtFin  time.Time `gorm:"type:date"`
	Cnpj   string    `gorm:"type:varchar(14)"`
}

func (RegC465) TableName() string {
	return "reg_C465"
}
