package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

// RegC170 representa o registro C170 do SPED Fiscal que contém os itens das notas fiscais
type RegC170 struct {
	gorm.Model
	Reg           string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C170"
	NumItem       string    `gorm:"type:varchar(3)"`        // Número sequencial do item
	CodItem       string    `gorm:"type:varchar(60);index"` // Código do item
	DescrCompl    string    `gorm:"type:varchar(255)"`      // Descrição complementar do item
	Qtd           float64   `gorm:"type:decimal(19,5)"`     // Quantidade do item
	Unid          string    `gorm:"type:varchar(6)"`        // Unidade do item
	VlItem        float64   `gorm:"type:decimal(19,2)"`     // Valor total do item
	VlDesc        float64   `gorm:"type:decimal(19,2)"`     // Valor do desconto do item
	IndMov        string    `gorm:"type:varchar(1)"`        // Indicador de movimentação física
	CstIcms       string    `gorm:"type:varchar(3)"`        // Código da situação tributária do ICMS
	Cfop          string    `gorm:"type:varchar(4);index"`  // Código fiscal de operação e prestação
	CodNat        string    `gorm:"type:varchar(10)"`       // Código da natureza da operação
	VlBcIcms      float64   `gorm:"type:decimal(19,2)"`     // Base de cálculo do ICMS
	AliqIcms      float64   `gorm:"type:decimal(6,2)"`      // Alíquota do ICMS
	VlIcms        float64   `gorm:"type:decimal(19,2)"`     // Valor do ICMS
	VlBcIcmsSt    float64   `gorm:"type:decimal(19,2)"`     // Base de cálculo do ICMS ST
	AliqSt        float64   `gorm:"type:decimal(19,2)"`     // Alíquota do ICMS ST
	VlIcmsSt      float64   `gorm:"type:decimal(19,2)"`     // Valor do ICMS ST
	IndApur       string    `gorm:"type:varchar(1)"`        // Indicador de período de apuração do IPI
	CstIpi        string    `gorm:"type:varchar(2)"`        // Código da situação tributária do IPI
	CodEnq        string    `gorm:"type:varchar(3)"`        // Código de enquadramento do IPI
	VlBcIpi       float64   `gorm:"type:decimal(19,2)"`     // Base de cálculo do IPI
	AliqIpi       float64   `gorm:"type:decimal(6,2)"`      // Alíquota do IPI
	VlIpi         float64   `gorm:"type:decimal(19,2)"`     // Valor do IPI
	CstPis        string    `gorm:"type:varchar(2)"`        // Código da situação tributária do PIS
	VlBcPis       float64   `gorm:"type:decimal(19,2)"`     // Base de cálculo do PIS
	AliqPis01     float64   `gorm:"type:decimal(8,4)"`      // Alíquota do PIS (em percentual)
	QuantBcPis    float64   `gorm:"type:decimal(19,3)"`     // Quantidade base de cálculo do PIS
	AliqPis02     float64   `gorm:"type:decimal(8,4)"`      // Alíquota do PIS (em reais)
	VlPis         float64   `gorm:"type:decimal(19,2)"`     // Valor do PIS
	CstCofins     string    `gorm:"type:varchar(2)"`        // Código da situação tributária do COFINS
	VlBcCofins    float64   `gorm:"type:decimal(19,2)"`     // Base de cálculo do COFINS
	AliqCofins01  float64   `gorm:"type:decimal(8,4)"`      // Alíquota do COFINS (em percentual)
	QuantBcCofins float64   `gorm:"type:decimal(19,3)"`     // Quantidade base de cálculo do COFINS
	AliqCofins02  float64   `gorm:"type:decimal(8,4)"`      // Alíquota do COFINS (em reais)
	VlCofins      float64   `gorm:"type:decimal(19,2)"`     // Valor do COFINS
	CodCta        string    `gorm:"type:varchar(255)"`      // Código da conta analítica contábil
	EntradaSaida  string    `gorm:"type:varchar(1)"`        // Indicador de entrada/saída
	NumDoc        string    `gorm:"type:varchar(9);index"`  // Número do documento fiscal
	DtIni         time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin         time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj          string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC170) TableName() string {
	return "reg_c170"
}

// Validações dos campos
func (r *RegC170) Validate() error {
	if r.Reg != "C170" {
		return ErrInvalidRegC170
	}
	if r.CodItem == "" {
		return ErrEmptyCodItem
	}
	if r.Qtd <= 0 {
		return ErrInvalidQtd
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateC170
	}
	return nil
}

// Implementando Interface do SpedRegC170
type RegC170Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
	RegC100 RegC100
	Digito  string
}

type iRegC170 interface {
	GetRegC170() RegC170
}

// GetRegC170 converte a linha do SPED em struct RegC170
func (s RegC170Sped) GetRegC170() RegC170 {
	digitoInt := tools.ConvInt(s.Digito)
	regC170 := RegC170{
		Reg:           s.Ln[1],
		NumItem:       s.Ln[2],
		CodItem:       tools.AdicionaDigitosCodigo(s.Ln[3], digitoInt),
		DescrCompl:    s.Ln[4],
		Qtd:           tools.ConvFloat(s.Ln[5]),
		Unid:          s.Ln[6],
		VlItem:        tools.ConvFloat(s.Ln[7]),
		VlDesc:        tools.ConvFloat(s.Ln[8]),
		IndMov:        s.Ln[9],
		CstIcms:       s.Ln[10],
		Cfop:          s.Ln[11],
		CodNat:        s.Ln[12],
		VlBcIcms:      tools.ConvFloat(s.Ln[13]),
		AliqIcms:      tools.ConvFloat(s.Ln[14]),
		VlIcms:        tools.ConvFloat(s.Ln[15]),
		VlBcIcmsSt:    tools.ConvFloat(s.Ln[16]),
		AliqSt:        tools.ConvFloat(s.Ln[17]),
		VlIcmsSt:      tools.ConvFloat(s.Ln[18]),
		IndApur:       s.Ln[19],
		CstIpi:        s.Ln[20],
		CodEnq:        s.Ln[21],
		VlBcIpi:       tools.ConvFloat(s.Ln[22]),
		AliqIpi:       tools.ConvFloat(s.Ln[23]),
		VlIpi:         tools.ConvFloat(s.Ln[24]),
		CstPis:        s.Ln[25],
		VlBcPis:       tools.ConvFloat(s.Ln[26]),
		AliqPis01:     tools.ConvFloat(s.Ln[27]),
		QuantBcPis:    tools.ConvFloat(s.Ln[28]),
		AliqPis02:     tools.ConvFloat(s.Ln[29]),
		VlPis:         tools.ConvFloat(s.Ln[30]),
		CstCofins:     s.Ln[31],
		VlBcCofins:    tools.ConvFloat(s.Ln[32]),
		AliqCofins01:  tools.ConvFloat(s.Ln[33]),
		QuantBcCofins: tools.ConvFloat(s.Ln[34]),
		AliqCofins02:  tools.ConvFloat(s.Ln[35]),
		VlCofins:      tools.ConvFloat(s.Ln[36]),
		CodCta:        s.Ln[37],
		EntradaSaida:  s.RegC100.IndOper,
		NumDoc:        s.RegC100.NumDoc,
		DtIni:         s.Reg0000.DtIni,
		DtFin:         s.Reg0000.DtFin,
		Cnpj:          s.Reg0000.Cnpj,
	}
	return regC170
}

// CreateRegC170 cria um novo registro RegC170
func CreateRegC170(read iRegC170) (RegC170, error) {
	reg := read.GetRegC170()
	if err := reg.Validate(); err != nil {
		return RegC170{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC170  = errors.New("registro inválido, deve ser C170")
	ErrEmptyCodItem    = errors.New("código do item não pode ser vazio")
	ErrInvalidQtd      = errors.New("quantidade deve ser maior que zero")
	ErrInvalidDateC170 = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *RegC170) GetValorTotal() float64 {
	return r.VlItem
}

func (r *RegC170) GetValorLiquido() float64 {
	return r.VlItem - r.VlDesc
}

func (r *RegC170) IsEntrada() bool {
	return r.EntradaSaida == "0"
}

func (r *RegC170) IsSaida() bool {
	return r.EntradaSaida == "1"
}
