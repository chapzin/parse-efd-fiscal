package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC470 representa o registro C470 do SPED Fiscal que contém os itens do documento fiscal emitido por ECF
type RegC470 struct {
	gorm.Model
	Reg     string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C470"
	CodItem string    `gorm:"type:varchar(60);index"` // Código do item
	Qtd     float64   `gorm:"type:decimal(19,3)"`     // Quantidade do item
	Qtdcanc float64   `gorm:"type:decimal(19,3)"`     // Quantidade cancelada
	Unid    string    `gorm:"type:varchar(6)"`        // Unidade do item
	VlItem  float64   `gorm:"type:decimal(19,2)"`     // Valor total do item
	CstIcms string    `gorm:"type:varchar(3)"`        // Código da situação tributária do ICMS
	Cfop    string    `gorm:"type:varchar(4);index"`  // Código fiscal de operação e prestação
	AliqIcm float64   `gorm:"type:decimal(6,2)"`      // Alíquota do ICMS
	VlPis   float64   `gorm:"type:decimal(19,2)"`     // Valor do PIS
	VlCofin float64   `gorm:"type:decimal(19,2)"`     // Valor do COFINS
	DtIni   time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin   time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj    string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC470) TableName() string {
	return "reg_c470"
}

// Validações dos campos
func (r *RegC470) Validate() error {
	if r.Reg != "C470" {
		return ErrInvalidRegC470
	}
	if r.CodItem == "" {
		return ErrEmptyCodItemC470
	}
	if r.Qtd <= 0 {
		return ErrInvalidQtdC470
	}
	if r.Unid == "" {
		return ErrEmptyUnidC470
	}
	if r.VlItem < 0 {
		return ErrInvalidVlItemC470
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateC470
	}
	return nil
}

// Interface que define o contrato para criar RegC470
type iRegC470 interface {
	GetRegC470() RegC470
}

// RegC470Sped representa a estrutura do arquivo SPED
type RegC470Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
	Digito  string
}

// GetRegC470 converte a linha do SPED em struct RegC470
func (s RegC470Sped) GetRegC470() RegC470 {
	digitoInt := tools.ConvInt(s.Digito)
	regC470 := RegC470{
		Reg:     s.Ln[1],
		CodItem: tools.AdicionaDigitosCodigo(s.Ln[2], digitoInt),
		Qtd:     tools.ConvFloat(s.Ln[3]),
		Qtdcanc: tools.ConvFloat(s.Ln[4]),
		Unid:    s.Ln[5],
		VlItem:  tools.ConvFloat(s.Ln[6]),
		CstIcms: s.Ln[7],
		Cfop:    s.Ln[8],
		AliqIcm: tools.ConvFloat(s.Ln[9]),
		VlPis:   tools.ConvFloat(s.Ln[10]),
		VlCofin: tools.ConvFloat(s.Ln[11]),
		DtIni:   s.Reg0000.DtIni,
		DtFin:   s.Reg0000.DtFin,
		Cnpj:    s.Reg0000.Cnpj,
	}
	return regC470
}

// CreateRegC470 cria um novo registro RegC470
func CreateRegC470(read iRegC470) (RegC470, error) {
	reg := read.GetRegC470()
	if err := reg.Validate(); err != nil {
		return RegC470{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC470   = errors.New("registro inválido, deve ser C470")
	ErrEmptyCodItemC470 = errors.New("código do item não pode ser vazio")
	ErrInvalidQtdC470   = errors.New("quantidade deve ser maior que zero")
	ErrEmptyUnidC470    = errors.New("unidade não pode ser vazia")
	ErrInvalidVlItemC470 = errors.New("valor do item não pode ser negativo")
	ErrInvalidDateC470   = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *RegC470) GetValorTotal() float64 {
	return r.VlItem
}

func (r *RegC470) GetValorTributos() float64 {
	return r.VlPis + r.VlCofin
}

func (r *RegC470) GetQuantidadeLiquida() float64 {
	return r.Qtd - r.Qtdcanc
}

func (r *RegC470) GetQuantidadeFormatada() string {
	return tools.FloatToString(r.GetQuantidadeLiquida()) + " " + r.Unid
}

func (r *RegC470) IsItemAtivo(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}
