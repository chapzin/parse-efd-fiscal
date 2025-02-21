package BlocoC

import (
	"errors"
	"strings"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC870 representa o registro C870 do SPED Fiscal que contém os itens do resumo diário dos documentos CF-e-SAT
type RegC870 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C870"
	CodItem  string    `gorm:"type:varchar(60);index"` // Código do item
	Qtd      float64   `gorm:"type:decimal(19,5)"`     // Quantidade do item
	Unid     string    `gorm:"type:varchar(6)"`        // Unidade do item
	CstIcms  string    `gorm:"type:varchar(3)"`        // Código da Situação Tributária do ICMS
	Cfop     string    `gorm:"type:varchar(4);index"`  // Código Fiscal de Operação e Prestação
	DtIni    time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin    time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj     string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC870) TableName() string {
	return "reg_c870"
}

// Validações dos campos
func (r *RegC870) Validate() error {
	if r.Reg != "C870" {
		return ErrInvalidRegC870
	}
	if r.CodItem == "" {
		return ErrEmptyCodItemC870
	}
	if r.Qtd <= 0 {
		return ErrInvalidQtdC870
	}
	if r.Unid == "" {
		return ErrEmptyUnidC870
	}
	if r.CstIcms == "" {
		return ErrEmptyCstIcmsC870
	}
	if r.Cfop == "" {
		return ErrEmptyCfopC870
	}
	if !strings.HasPrefix(r.Cfop, "5") {
		return ErrInvalidCfopC870
	}
	return nil
}

// Interface que define o contrato para criar RegC870
type iRegC870 interface {
	GetRegC870() RegC870
}

// RegC870Sped representa a estrutura do arquivo SPED
type RegC870Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
	Digito  string
}

// GetRegC870 converte a linha do SPED em struct RegC870
func (s RegC870Sped) GetRegC870() RegC870 {
	digitoInt := tools.ConvInt(s.Digito)
	regC870 := RegC870{
		Reg:     s.Ln[1],
		CodItem: tools.AdicionaDigitosCodigo(s.Ln[2], digitoInt),
		Qtd:     tools.ConvFloat(s.Ln[3]),
		Unid:    s.Ln[4],
		CstIcms: s.Ln[5],
		Cfop:    s.Ln[6],
		DtIni:   s.Reg0000.DtIni,
		DtFin:   s.Reg0000.DtFin,
		Cnpj:    s.Reg0000.Cnpj,
	}
	return regC870
}

// CreateRegC870 cria um novo registro RegC870
func CreateRegC870(read iRegC870) (RegC870, error) {
	reg := read.GetRegC870()
	if err := reg.Validate(); err != nil {
		return RegC870{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC870   = errors.New("registro inválido, deve ser C870")
	ErrEmptyCodItemC870 = errors.New("código do item não pode ser vazio")
	ErrInvalidQtdC870   = errors.New("quantidade deve ser maior que zero")
	ErrEmptyUnidC870    = errors.New("unidade não pode ser vazia")
	ErrEmptyCstIcmsC870 = errors.New("CST do ICMS não pode ser vazio")
	ErrEmptyCfopC870    = errors.New("CFOP não pode ser vazio")
	ErrInvalidCfopC870  = errors.New("CFOP deve iniciar com 5")
)

// Métodos auxiliares
func (r *RegC870) GetQuantidadeFormatada() string {
	return tools.FloatToString(r.Qtd) + " " + r.Unid
}

func (r *RegC870) GetIdentificacaoItem() string {
	return r.CodItem + " - CST: " + r.CstIcms + " - CFOP: " + r.Cfop
}

func (r *RegC870) IsItemAtivo(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
} 