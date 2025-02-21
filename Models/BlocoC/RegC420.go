package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC420 representa o registro C420 do SPED Fiscal que contém o registro dos totalizadores parciais da redução Z
type RegC420 struct {
	gorm.Model
	Reg        string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C420"
	CodTotPar  string    `gorm:"type:varchar(7)"`        // Código do totalizador parcial
	VlrAcumTot float64   `gorm:"type:decimal(19,2)"`     // Valor acumulado no totalizador
	NrTot      string    `gorm:"type:varchar(2)"`        // Número do totalizador quando ocorrer mais de uma situação
	DescrNrTot string    `gorm:"type:varchar(255)"`      // Descrição da situação tributária
	DtIni      time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin      time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj       string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC420) TableName() string {
	return "reg_c420"
}

// Validações dos campos
func (r *RegC420) Validate() error {
	if r.Reg != "C420" {
		return ErrInvalidRegC420
	}
	if r.CodTotPar == "" {
		return ErrEmptyCodTotPar
	}
	if r.VlrAcumTot < 0 {
		return ErrInvalidVlrAcumTot
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateC420
	}
	return nil
}

// Interface que define o contrato para criar RegC420
type iRegC420 interface {
	GetRegC420() RegC420
}

// RegC420Sped representa a estrutura do arquivo SPED
type RegC420Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

// GetRegC420 converte a linha do SPED em struct RegC420
func (s RegC420Sped) GetRegC420() RegC420 {
	regC420 := RegC420{
		Reg:        s.Ln[1],
		CodTotPar:  s.Ln[2],
		VlrAcumTot: tools.ConvFloat(s.Ln[3]),
		NrTot:      s.Ln[4],
		DescrNrTot: s.Ln[5],
		DtIni:      s.Reg0000.DtIni,
		DtFin:      s.Reg0000.DtFin,
		Cnpj:       s.Reg0000.Cnpj,
	}
	return regC420
}

// CreateRegC420 cria um novo registro RegC420
func CreateRegC420(read iRegC420) (RegC420, error) {
	reg := read.GetRegC420()
	if err := reg.Validate(); err != nil {
		return RegC420{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC420    = errors.New("registro inválido, deve ser C420")
	ErrEmptyCodTotPar    = errors.New("código do totalizador parcial não pode ser vazio")
	ErrInvalidVlrAcumTot = errors.New("valor acumulado não pode ser negativo")
	ErrInvalidDateC420   = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *RegC420) GetDescricaoCompleta() string {
	descr := r.CodTotPar
	if r.DescrNrTot != "" {
		descr += " - " + r.DescrNrTot
	}
	return descr
}

func (r *RegC420) GetValorAcumulado() float64 {
	return r.VlrAcumTot
}

func (r *RegC420) IsTotalizadorAtivo(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}
