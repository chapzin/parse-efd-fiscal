package Bloco0

import (
	"github.com/jinzhu/gorm"
	"time"
)


type Reg0000 struct {
	gorm.Model
	Reg string	`gorm:"type:varchar(4)"`
	CodVer string	`gorm:"type:varchar(3)"`
	CodFin int
	DtIni time.Time	`gorm:"type:date;unique_index"`
	DtFin time.Time	`gorm:"type:date;unique_index"`
	Nome string	`gorm:"type:varchar(100)"`
	Cnpj string	`gorm:"type:varchar(14)"`
	Cpf string	`gorm:"type:varchar(11)"`
	Uf string	`gorm:"type:varchar(2)"`
	Ie string	`gorm:"type:varchar(14)"`
	CodMun string	`gorm:"type:varchar(7)"`
	Im string
	Suframa string	`gorm:"type:varchar(9)"`
	IndPerfil string	`gorm:"type:varchar(1)"`
	IndAtiv int
}

func (Reg0000) TableName() string {
	return "reg_0000"
}