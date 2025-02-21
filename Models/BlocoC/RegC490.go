package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC490 representa o registro C490 do SPED Fiscal que contém o resumo do movimento diário por ECF
type RegC490 struct {
	gorm.Model
	Reg       string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C490"
	CstIcms   string    `gorm:"type:varchar(3)"`        // Código da situação tributária do ICMS
	Cfop      string    `gorm:"type:varchar(4);index"`  // Código fiscal de operação e prestação
	AliqIcms  float64   `gorm:"type:decimal(6,2)"`      // Alíquota do ICMS
	VlOpr     float64   `gorm:"type:decimal(19,2)"`     // Valor da operação correspondente à combinação CST_ICMS, CFOP e alíquota do ICMS
	VlBcIcms  float64   `gorm:"type:decimal(19,2)"`     // Valor acumulado da base de cálculo do ICMS
	VlIcms    float64   `gorm:"type:decimal(19,2)"`     // Valor acumulado do ICMS
	CodObs    string    `gorm:"type:varchar(6)"`        // Código da observação fiscal
	DtIni     time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin     time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj      string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC490) TableName() string {
	return "reg_c490"
}

// Validações dos campos
func (r *RegC490) Validate() error {
	if r.Reg != "C490" {
		return ErrInvalidRegC490
	}
	if r.CstIcms == "" {
		return ErrEmptyCstIcmsC490
	}
	if r.Cfop == "" {
		return ErrEmptyCfopC490
	}
	if r.VlOpr < 0 {
		return ErrInvalidVlOprC490
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateC490
	}
	return nil
}

// Interface que define o contrato para criar RegC490
type iRegC490 interface {
	GetRegC490() RegC490
}

// RegC490Sped representa a estrutura do arquivo SPED
type RegC490Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

// GetRegC490 converte a linha do SPED em struct RegC490
func (s RegC490Sped) GetRegC490() RegC490 {
	regC490 := RegC490{
		Reg:      s.Ln[1],
		CstIcms:  s.Ln[2],
		Cfop:     s.Ln[3],
		AliqIcms: tools.ConvFloat(s.Ln[4]),
		VlOpr:    tools.ConvFloat(s.Ln[5]),
		VlBcIcms: tools.ConvFloat(s.Ln[6]),
		VlIcms:   tools.ConvFloat(s.Ln[7]),
		CodObs:   s.Ln[8],
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
	}
	return regC490
}

// CreateRegC490 cria um novo registro RegC490
func CreateRegC490(read iRegC490) (RegC490, error) {
	reg := read.GetRegC490()
	if err := reg.Validate(); err != nil {
		return RegC490{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC490   = errors.New("registro inválido, deve ser C490")
	ErrEmptyCstIcmsC490 = errors.New("CST do ICMS não pode ser vazio")
	ErrEmptyCfopC490    = errors.New("CFOP não pode ser vazio")
	ErrInvalidVlOprC490 = errors.New("valor da operação não pode ser negativo")
	ErrInvalidDateC490  = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *RegC490) GetValorTotal() float64 {
	return r.VlOpr
}

func (r *RegC490) GetValorImposto() float64 {
	return r.VlIcms
}

func (r *RegC490) GetIdentificacaoOperacao() string {
	return "CST: " + r.CstIcms + " - CFOP: " + r.Cfop + " - Alíq: " + tools.FloatToString(r.AliqIcms) + "%"
}

func (r *RegC490) IsOperacaoAtiva(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}
