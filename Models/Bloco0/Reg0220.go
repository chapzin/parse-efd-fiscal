package Bloco0

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// Reg0220 representa o registro 0220 do SPED Fiscal que contém os fatores de conversão de unidades
type Reg0220 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "0220"
	UnidConv string    `gorm:"type:varchar(6);index"`  // Unidade comercial a ser convertida
	FatConv  float64   `gorm:"type:decimal(12,6)"`     // Fator de conversão
	UnidCod  string    `gorm:"type:varchar(6);index"`  // Unidade de medida do inventário
	CodItem  string    `gorm:"type:varchar(60);index"` // Código do item
	DtIni    time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin    time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj     string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
	Feito    string    `gorm:"type:varchar(1)"`        // Indicador de processamento
}

// TableName retorna o nome da tabela no banco de dados
func (Reg0220) TableName() string {
	return "reg_0220"
}

// Validações dos campos
func (r *Reg0220) Validate() error {
	if r.Reg != "0220" {
		return ErrInvalidReg0220
	}
	if r.UnidConv == "" {
		return ErrEmptyUnidConv
	}
	if r.UnidCod == "" {
		return ErrEmptyUnidCod
	}
	if r.FatConv <= 0 {
		return ErrInvalidFatConv
	}
	if r.CodItem == "" {
		return ErrEmptyCodItem0220
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateRange0220
	}
	return nil
}

// Interface que define o contrato para criar Reg0220
type iReg0220 interface {
	GetReg0220() Reg0220
}

// Reg0220Sped representa a estrutura do arquivo SPED
type Reg0220Sped struct {
	Ln      []string
	Reg0000 Reg0000
	Reg0200 Reg0200
	Digito  string
}

// GetReg0220 converte a linha do SPED em struct Reg0220
func (s Reg0220Sped) GetReg0220() Reg0220 {
	digitoInt := tools.ConvInt(s.Digito)
	reg0220 := Reg0220{
		Reg:      s.Ln[1],
		UnidConv: s.Ln[2],
		FatConv:  tools.ConvFloat(s.Ln[3]),
		UnidCod:  s.Reg0200.UnidInv,
		CodItem:  tools.AdicionaDigitosCodigo(s.Reg0200.CodItem, digitoInt),
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
		Feito:    "0",
	}
	return reg0220
}

// CreateReg0220 cria um novo registro Reg0220
func CreateReg0220(read iReg0220) (Reg0220, error) {
	reg := read.GetReg0220()
	if err := reg.Validate(); err != nil {
		return Reg0220{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidReg0220       = errors.New("registro inválido, deve ser 0220")
	ErrEmptyUnidConv        = errors.New("unidade de conversão não pode ser vazia")
	ErrEmptyUnidCod         = errors.New("unidade de código não pode ser vazia")
	ErrInvalidFatConv       = errors.New("fator de conversão deve ser maior que zero")
	ErrEmptyCodItem0220     = errors.New("código do item não pode ser vazio")
	ErrInvalidDateRange0220 = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *Reg0220) IsConversaoAtiva(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}

func (r *Reg0220) GetFatorConversao() float64 {
	return r.FatConv
}

func (r *Reg0220) ConvertQuantidade(qtd float64) float64 {
	return qtd * r.FatConv
}

func (r *Reg0220) GetDescricaoConversao() string {
	return r.UnidCod + " -> " + r.UnidConv + " (x" + tools.FloatToString(r.FatConv) + ")"
}

func (r *Reg0220) IsProcessado() bool {
	return r.Feito == "1"
}

func (r *Reg0220) SetProcessado() {
	r.Feito = "1"
}
