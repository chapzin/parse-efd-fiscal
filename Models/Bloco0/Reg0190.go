package Bloco0

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// Reg0190 representa o registro 0190 do SPED Fiscal que contém a identificação das unidades de medida
type Reg0190 struct {
	gorm.Model
	Reg   string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "0190"
	Unid  string    `gorm:"type:varchar(6);index"`  // Código da unidade de medida
	Descr string    `gorm:"type:varchar(255)"`      // Descrição da unidade de medida
	DtIni time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj  string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte do SPED
}

// TableName retorna o nome da tabela no banco de dados
func (Reg0190) TableName() string {
	return "reg_0190"
}

// Validações dos campos
func (r *Reg0190) Validate() error {
	if r.Reg != "0190" {
		return ErrInvalidReg0190
	}
	if r.Unid == "" {
		return ErrEmptyUnid
	}
	if r.Descr == "" {
		return ErrEmptyDescr
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateRange0190
	}
	return nil
}

// Interface que define o contrato para criar Reg0190
type iReg0190 interface {
	GetReg0190() Reg0190
}

// Reg0190Sped representa a estrutura do arquivo SPED
type Reg0190Sped struct {
	Ln      []string
	Reg0000 Reg0000
}

// GetReg0190 converte a linha do SPED em struct Reg0190
func (s Reg0190Sped) GetReg0190() Reg0190 {
	reg0190 := Reg0190{
		Reg:   s.Ln[1],
		Unid:  s.Ln[2],
		Descr: s.Ln[3],
		DtIni: s.Reg0000.DtIni,
		DtFin: s.Reg0000.DtFin,
		Cnpj:  s.Reg0000.Cnpj,
	}
	return reg0190
}

// Reg0190Xml representa a estrutura para importação via XML
type Reg0190Xml struct {
	Data string
}

// GetReg0190 converte os dados do XML em struct Reg0190
func (x Reg0190Xml) GetReg0190() Reg0190 {
	reg0190 := Reg0190{
		Reg:   "0190",
		Unid:  x.Data,
		Descr: "Importado Xml",
		DtIni: tools.ConvertDataNull(),
		DtFin: tools.ConvertDataNull(),
		Cnpj:  "",
	}
	return reg0190
}

// CreateReg0190 cria um novo registro Reg0190
func CreateReg0190(read iReg0190) (Reg0190, error) {
	reg := read.GetReg0190()
	if err := reg.Validate(); err != nil {
		return Reg0190{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidReg0190       = errors.New("registro inválido, deve ser 0190")
	ErrEmptyUnid            = errors.New("código da unidade não pode ser vazio")
	ErrEmptyDescr           = errors.New("descrição da unidade não pode ser vazia")
	ErrInvalidDateRange0190 = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *Reg0190) IsUnidadeAtiva(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}

func (r *Reg0190) GetUnidadeFormatada() string {
	return r.Unid + " - " + r.Descr
}
