package Bloco0

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// Estrutura criada usando layout Guia Prático EFD-ICMS/IPI – Versão 2.0.20 Atualização: 07/12/2016

// Reg0150 representa o registro 0150 do SPED Fiscal que contém o cadastro de participantes
type Reg0150 struct {
	gorm.Model
	Reg      string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "0150"
	CodPart  string    `gorm:"type:varchar(60);index"` // Código do participante
	Nome     string    `gorm:"type:varchar(100)"`      // Nome do participante
	CodPais  string    `gorm:"type:varchar(5)"`        // Código do país do participante
	Cnpj     string    `gorm:"type:varchar(15);index"` // CNPJ do participante
	Cpf      string    `gorm:"type:varchar(11);index"` // CPF do participante
	Ie       string    `gorm:"type:varchar(14);index"` // Inscrição Estadual do participante
	CodMun   string    `gorm:"type:varchar(7)"`        // Código do município
	Suframa  string    `gorm:"type:varchar(9)"`        // Número de inscrição na Suframa
	Endereco string    `gorm:"type:varchar(60)"`       // Logradouro e endereço do imóvel
	Num      string    `gorm:"type:varchar(10)"`       // Número do imóvel
	Compl    string    `gorm:"type:varchar(60)"`       // Complemento do endereço
	Bairro   string    `gorm:"type:varchar(60)"`       // Bairro do imóvel
	DtIni    time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin    time.Time `gorm:"type:date;index"`        // Data final das informações
	CnpjSped string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte do SPED
}

// TableName retorna o nome da tabela no banco de dados
func (Reg0150) TableName() string {
	return "reg_0150"
}

// Validações dos campos
func (r *Reg0150) Validate() error {
	if r.Reg != "0150" {
		return ErrInvalidReg0150
	}
	if r.CodPart == "" {
		return ErrEmptyCodPart
	}
	if r.Nome == "" {
		return ErrEmptyNome
	}
	if r.Cnpj == "" && r.Cpf == "" {
		return ErrMissingIdentification0150
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateRange0150
	}
	return nil
}

// Interface que define o contrato para criar Reg0150
type iReg0150 interface {
	GetReg0150() Reg0150
}

// Estrutura necessaria para polular o Reg150 pelo sped fiscal
type Reg0150Sped struct {
	Ln      []string
	Reg0000 Reg0000
}

// GetReg0150 converte a linha do SPED em struct Reg0150
func (s Reg0150Sped) GetReg0150() Reg0150 {
	reg0150 := Reg0150{
		Reg:      s.Ln[1],
		CodPart:  s.Ln[2],
		Nome:     s.Ln[3],
		CodPais:  s.Ln[4],
		Cnpj:     s.Ln[5],
		Cpf:      s.Ln[6],
		Ie:       s.Ln[7],
		CodMun:   s.Ln[8],
		Suframa:  s.Ln[9],
		Endereco: s.Ln[10],
		Num:      s.Ln[11],
		Compl:    s.Ln[12],
		Bairro:   s.Ln[13],
		DtIni:    s.Reg0000.DtIni,
		DtFin:    s.Reg0000.DtFin,
		CnpjSped: s.Reg0000.Cnpj,
	}
	return reg0150
}

// Estrutura necessaria para popular o Reg0150 pelo Xml
type Reg0150Xml struct {
	Reader func(pathTag string, tag string) string
}

// GetReg0150 converte os dados do XML em struct Reg0150
func (x Reg0150Xml) GetReg0150() Reg0150 {
	reg0150 := Reg0150{
		Reg:      "0150",
		CodPart:  x.Reader("dest", "CNPJ"),
		Nome:     x.Reader("dest", "xNome"),
		CodPais:  "1058", // TODO: Mover para constante
		Cnpj:     x.Reader("dest", "CNPJ"),
		Cpf:      x.Reader("dest", "CNPJ"),
		Ie:       x.Reader("dest", "IE"),
		CodMun:   x.Reader("enderDest", "cMun"),
		Suframa:  "",
		Endereco: x.Reader("enderDest", "xLgr"),
		Num:      x.Reader("enderDest", "nro"),
		Compl:    x.Reader("enderDest", "xCpl"),
		Bairro:   x.Reader("enderDest", "xBairro"),
		DtIni:    tools.ConvertDataNull(),
		DtFin:    tools.ConvertDataNull(),
		CnpjSped: "Insert Xml",
	}
	return reg0150
}

// CreateReg0150 cria um novo registro Reg0150
func CreateReg0150(read iReg0150) (Reg0150, error) {
	reg := read.GetReg0150()
	if err := reg.Validate(); err != nil {
		return Reg0150{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidReg0150            = errors.New("registro inválido, deve ser 0150")
	ErrEmptyCodPart              = errors.New("código do participante não pode ser vazio")
	ErrEmptyNome                 = errors.New("nome do participante não pode ser vazio")
	ErrMissingIdentification0150 = errors.New("CNPJ ou CPF do participante deve ser informado")
	ErrInvalidDateRange0150      = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *Reg0150) IsParticipanteAtivo(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}

func (r *Reg0150) GetEnderecoCompleto() string {
	endereco := r.Endereco
	if r.Num != "" {
		endereco += ", " + r.Num
	}
	if r.Compl != "" {
		endereco += " - " + r.Compl
	}
	if r.Bairro != "" {
		endereco += ", " + r.Bairro
	}
	return endereco
}

func (r *Reg0150) IsPessoaJuridica() bool {
	return r.Cnpj != ""
}
