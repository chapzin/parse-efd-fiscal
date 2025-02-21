package BlocoC

import (
	"errors"
	"strings"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC890 representa o registro C890 do SPED Fiscal que contém o resumo diário de CF-e-SAT por equipamento SAT
type RegC890 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C890"
	CstIcms  string    `gorm:"type:varchar(3)"`        // Código da Situação Tributária do ICMS
	Cfop     string    `gorm:"type:varchar(4);index"`  // Código Fiscal de Operação e Prestação
	AliqIcms float64   `gorm:"type:decimal(6,2)"`      // Alíquota do ICMS
	VlOpr    float64   `gorm:"type:decimal(19,2)"`     // Valor total do CF-e na combinação de CST_ICMS, CFOP e ALÍQUOTA DO ICMS
	VlBcIcms float64   `gorm:"type:decimal(19,2)"`     // Valor acumulado da base de cálculo do ICMS
	VlIcms   float64   `gorm:"type:decimal(19,2)"`     // Valor do ICMS
	CodObs   string    `gorm:"type:varchar(6)"`        // Código da observação do lançamento fiscal (campo 02 do registro 0460)
	DtDoc    time.Time `gorm:"type:date;index"`        // Data do documento fiscal
	NrSat    string    `gorm:"type:varchar(9);index"`  // Número de Série do equipamento SAT
	DtIni    time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin    time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj     string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC890) TableName() string {
	return "reg_c890"
}

// Validações dos campos
func (r *RegC890) Validate() error {
	if r.Reg != "C890" {
		return ErrInvalidRegC890
	}
	if r.CstIcms == "" {
		return ErrEmptyCstIcmsC890
	}
	if r.Cfop == "" {
		return ErrEmptyCfopC890
	}
	if !strings.HasPrefix(r.Cfop, "5") {
		return ErrInvalidCfopC890
	}
	if r.AliqIcms < 0 {
		return ErrInvalidAliqIcmsC890
	}
	if r.VlOpr < 0 {
		return ErrInvalidVlOprC890
	}
	if r.VlBcIcms < 0 {
		return ErrInvalidVlBcIcmsC890
	}
	if r.VlIcms < 0 {
		return ErrInvalidVlIcmsC890
	}
	return nil
}

// Interface que define o contrato para criar RegC890
type iRegC890 interface {
	GetRegC890() RegC890
}

// RegC890Sped representa a estrutura do arquivo SPED
type RegC890Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
	DtDoc   time.Time
	NrSat   string
}

// GetRegC890 converte a linha do SPED em struct RegC890
func (s RegC890Sped) GetRegC890() RegC890 {
	regC890 := RegC890{
		Reg:      s.Ln[1],
		CstIcms:  s.Ln[2],
		Cfop:     s.Ln[3],
		AliqIcms: tools.ConvFloat(s.Ln[4]),
		VlOpr:    tools.ConvFloat(s.Ln[5]),
		VlBcIcms: tools.ConvFloat(s.Ln[6]),
		VlIcms:   tools.ConvFloat(s.Ln[7]),
		CodObs:   s.Ln[8],
		DtDoc:    s.DtDoc, // Deve ser preenchido com a data do C860 relacionado
		NrSat:    s.NrSat, // Deve ser preenchido com o número do SAT do C860 relacionado
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
	}
	return regC890
}

// CreateRegC890 cria um novo registro RegC890
func CreateRegC890(read iRegC890) (RegC890, error) {
	reg := read.GetRegC890()
	if err := reg.Validate(); err != nil {
		return RegC890{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC890      = errors.New("registro inválido, deve ser C890")
	ErrEmptyCstIcmsC890    = errors.New("CST do ICMS não pode ser vazio")
	ErrEmptyCfopC890       = errors.New("CFOP não pode ser vazio")
	ErrInvalidCfopC890     = errors.New("CFOP deve iniciar com 5 (operações internas)")
	ErrInvalidAliqIcmsC890 = errors.New("alíquota do ICMS não pode ser negativa")
	ErrInvalidVlOprC890    = errors.New("valor da operação não pode ser negativo")
	ErrInvalidVlBcIcmsC890 = errors.New("valor da base de cálculo do ICMS não pode ser negativo")
	ErrInvalidVlIcmsC890   = errors.New("valor do ICMS não pode ser negativo")
)

// Métodos auxiliares
func (r *RegC890) GetValorTotal() float64 {
	return r.VlOpr
}

func (r *RegC890) GetValorImposto() float64 {
	return r.VlIcms
}

func (r *RegC890) GetIdentificacaoOperacao() string {
	return "CST: " + r.CstIcms + " - CFOP: " + r.Cfop + " - Alíq: " + tools.FloatToString(r.AliqIcms) + "%"
}

func (r *RegC890) IsOperacaoAtiva(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}
