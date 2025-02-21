package BlocoC

import (
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/jinzhu/gorm"
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

type RegC465Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

type iRegC465 interface {
	GetRegC465() RegC465
}

func (s RegC465Sped) GetRegC465() RegC465 {
	regC465 := RegC465{
		Reg:    s.Ln[1],
		ChvCfe: s.Ln[2],
		NumCcF: s.Ln[3],
		DtIni:  s.Reg0000.DtIni,
		DtFin:  s.Reg0000.DtFin,
		Cnpj:   s.Reg0000.Cnpj,
	}
	return regC465
}

// Cria estrutura populada
func CreateRegC465(read iRegC465) RegC465 {
	return read.GetRegC465()
}
