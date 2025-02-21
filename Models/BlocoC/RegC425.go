package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC425 representa o registro C425 do SPED Fiscal que contém os itens do documento fiscal emitido por ECF
type RegC425 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C425"
	CodItem  string    `gorm:"type:varchar(60);index"` // Código do item
	Qtd      float64   `gorm:"type:decimal(19,3)"`     // Quantidade do item
	Unid     string    `gorm:"type:varchar(6)"`        // Unidade do item
	VlItem   float64   `gorm:"type:decimal(19,2)"`     // Valor total do item
	VlPis    float64   `gorm:"type:decimal(19,2)"`     // Valor do PIS
	VlCofins float64   `gorm:"type:decimal(19,2)"`     // Valor do COFINS
	DtIni    time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin    time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj     string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC425) TableName() string {
	return "reg_c425"
}

// Validações dos campos
func (r *RegC425) Validate() error {
	if r.Reg != "C425" {
		return ErrInvalidRegC425
	}
	if r.CodItem == "" {
		return ErrEmptyCodItemC425
	}
	if r.Qtd <= 0 {
		return ErrInvalidQtdC425
	}
	if r.Unid == "" {
		return ErrEmptyUnidC425
	}
	if r.VlItem < 0 {
		return ErrInvalidVlItemC425
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateC425
	}
	return nil
}

// Interface que define o contrato para criar RegC425
type iRegC425 interface {
	GetRegC425() RegC425
}

// RegC425Sped representa a estrutura do arquivo SPED
type RegC425Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
	Digito  string
}

// GetRegC425 converte a linha do SPED em struct RegC425
func (s RegC425Sped) GetRegC425() RegC425 {
	digitoInt := tools.ConvInt(s.Digito)
	regC425 := RegC425{
		Reg:      s.Ln[1],
		CodItem:  tools.AdicionaDigitosCodigo(s.Ln[2], digitoInt),
		Qtd:      tools.ConvFloat(s.Ln[3]),
		Unid:     s.Ln[4],
		VlItem:   tools.ConvFloat(s.Ln[5]),
		VlPis:    tools.ConvFloat(s.Ln[6]),
		VlCofins: tools.ConvFloat(s.Ln[7]),
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
	}
	return regC425
}

// CreateRegC425 cria um novo registro RegC425
func CreateRegC425(read iRegC425) (RegC425, error) {
	reg := read.GetRegC425()
	if err := reg.Validate(); err != nil {
		return RegC425{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC425    = errors.New("registro inválido, deve ser C425")
	ErrEmptyCodItemC425  = errors.New("código do item não pode ser vazio")
	ErrInvalidQtdC425    = errors.New("quantidade deve ser maior que zero")
	ErrEmptyUnidC425     = errors.New("unidade não pode ser vazia")
	ErrInvalidVlItemC425 = errors.New("valor do item não pode ser negativo")
	ErrInvalidDateC425   = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *RegC425) GetValorTotal() float64 {
	return r.VlItem
}

func (r *RegC425) GetValorTributos() float64 {
	return r.VlPis + r.VlCofins
}

func (r *RegC425) GetQuantidadeFormatada() string {
	return tools.FloatToString(r.Qtd) + " " + r.Unid
}

func (r *RegC425) IsItemAtivo(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}
