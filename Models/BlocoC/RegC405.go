package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC405 representa o registro C405 do SPED Fiscal que contém a redução Z
type RegC405 struct {
	gorm.Model
	Reg       string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C405"
	DtDoc     time.Time `gorm:"type:date;index"`        // Data do movimento
	Cro       string    `gorm:"type:varchar(3)"`        // Posição do contador de reinício de operação
	Crz       string    `gorm:"type:varchar(6)"`        // Posição do Contador de Redução Z
	NumCooFin string    `gorm:"type:varchar(9)"`        // Número do Contador de Ordem de Operação do último documento
	GtFin     float64   `gorm:"type:decimal(19,2)"`     // Valor do Grande Total final
	VlBrt     float64   `gorm:"type:decimal(19,2)"`     // Valor da venda bruta
	DtIni     time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin     time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj      string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC405) TableName() string {
	return "reg_c405"
}

// Validações dos campos
func (r *RegC405) Validate() error {
	if r.Reg != "C405" {
		return ErrInvalidRegC405
	}
	if r.Cro == "" {
		return ErrEmptyCro
	}
	if r.Crz == "" {
		return ErrEmptyCrz
	}
	if r.DtDoc.After(r.DtFin) {
		return ErrInvalidDateC405
	}
	if r.VlBrt < 0 {
		return ErrInvalidVlBrt
	}
	return nil
}

// Interface que define o contrato para criar RegC405
type iRegC405 interface {
	GetRegC405() RegC405
}

// RegC405Sped representa a estrutura do arquivo SPED
type RegC405Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

// GetRegC405 converte a linha do SPED em struct RegC405
func (s RegC405Sped) GetRegC405() RegC405 {
	regC405 := RegC405{
		Reg:       s.Ln[1],
		DtDoc:     tools.ConvertData(s.Ln[2]),
		Cro:       s.Ln[3],
		Crz:       s.Ln[4],
		NumCooFin: s.Ln[5],
		GtFin:     tools.ConvFloat(s.Ln[6]),
		VlBrt:     tools.ConvFloat(s.Ln[7]),
		DtIni:     s.Reg0000.DtIni,
		DtFin:     s.Reg0000.DtFin,
		Cnpj:      s.Reg0000.Cnpj,
	}
	return regC405
}

// CreateRegC405 cria um novo registro RegC405
func CreateRegC405(read iRegC405) (RegC405, error) {
	reg := read.GetRegC405()
	if err := reg.Validate(); err != nil {
		return RegC405{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC405  = errors.New("registro inválido, deve ser C405")
	ErrEmptyCro        = errors.New("CRO não pode ser vazio")
	ErrEmptyCrz        = errors.New("CRZ não pode ser vazio")
	ErrInvalidDateC405 = errors.New("data do documento não pode ser posterior à data final")
	ErrInvalidVlBrt    = errors.New("valor bruto não pode ser negativo")
)

// Métodos auxiliares
func (r *RegC405) GetIdentificacaoReducaoZ() string {
	return "CRO: " + r.Cro + " - CRZ: " + r.Crz
}

func (r *RegC405) GetValorBruto() float64 {
	return r.VlBrt
}

func (r *RegC405) IsReducaoZAtiva(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}
