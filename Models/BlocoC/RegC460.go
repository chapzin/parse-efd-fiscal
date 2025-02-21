package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC460 representa o registro C460 do SPED Fiscal que contém o documento fiscal emitido por ECF
type RegC460 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C460"
	CodMod   string    `gorm:"type:varchar(2)"`        // Código do modelo do documento fiscal
	CodSit   string    `gorm:"type:varchar(2)"`        // Código da situação do documento fiscal
	NumDoc   string    `gorm:"type:varchar(9);index"`  // Número do documento fiscal
	DtDoc    time.Time `gorm:"type:date;index"`        // Data da emissão do documento
	VlDoc    float64   `gorm:"type:decimal(19,2)"`     // Valor total do documento fiscal
	VlPis    float64   `gorm:"type:decimal(19,2)"`     // Valor do PIS
	VlCofins float64   `gorm:"type:decimal(19,2)"`     // Valor do COFINS
	CpfCnpj  string    `gorm:"type:varchar(14);index"` // CPF ou CNPJ do adquirente
	NomAdq   string    `gorm:"type:varchar(60)"`       // Nome do adquirente
	DtIni    time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin    time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj     string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC460) TableName() string {
	return "reg_c460"
}

// Validações dos campos
func (r *RegC460) Validate() error {
	if r.Reg != "C460" {
		return ErrInvalidRegC460
	}
	if r.CodMod == "" {
		return ErrEmptyCodModC460
	}
	if r.NumDoc == "" {
		return ErrEmptyNumDocC460
	}
	if r.VlDoc < 0 {
		return ErrInvalidVlDocC460
	}
	if r.DtDoc.After(r.DtFin) {
		return ErrInvalidDateC460
	}
	return nil
}

// Interface que define o contrato para criar RegC460
type iRegC460 interface {
	GetRegC460() RegC460
}

// RegC460Sped representa a estrutura do arquivo SPED
type RegC460Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

// GetRegC460 converte a linha do SPED em struct RegC460
func (s RegC460Sped) GetRegC460() RegC460 {
	regC460 := RegC460{
		Reg:      s.Ln[1],
		CodMod:   s.Ln[2],
		CodSit:   s.Ln[3],
		NumDoc:   s.Ln[4],
		DtDoc:    tools.ConvertData(s.Ln[5]),
		VlDoc:    tools.ConvFloat(s.Ln[6]),
		VlPis:    tools.ConvFloat(s.Ln[7]),
		VlCofins: tools.ConvFloat(s.Ln[8]),
		CpfCnpj:  s.Ln[9],
		NomAdq:   s.Ln[10],
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		Cnpj:     s.Reg0000.Cnpj,
	}
	return regC460
}

// CreateRegC460 cria um novo registro RegC460
func CreateRegC460(read iRegC460) (RegC460, error) {
	reg := read.GetRegC460()
	if err := reg.Validate(); err != nil {
		return RegC460{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC460   = errors.New("registro inválido, deve ser C460")
	ErrEmptyCodModC460  = errors.New("código do modelo não pode ser vazio")
	ErrEmptyNumDocC460  = errors.New("número do documento não pode ser vazio")
	ErrInvalidVlDocC460 = errors.New("valor do documento não pode ser negativo")
	ErrInvalidDateC460  = errors.New("data do documento não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *RegC460) GetValorTotal() float64 {
	return r.VlDoc
}

func (r *RegC460) GetValorTributos() float64 {
	return r.VlPis + r.VlCofins
}

func (r *RegC460) IsCancelado() bool {
	return r.CodSit == "02"
}

func (r *RegC460) GetIdentificacaoCliente() string {
	if r.NomAdq != "" {
		return r.NomAdq + " (" + r.CpfCnpj + ")"
	}
	return r.CpfCnpj
}
