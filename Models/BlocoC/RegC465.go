package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/jinzhu/gorm"
)

// RegC465 representa o registro C465 do SPED Fiscal que contém o complemento do cupom fiscal eletrônico emitido por ECF
type RegC465 struct {
	gorm.Model
	Reg    string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C465"
	ChvCfe string    `gorm:"type:varchar(44);index"` // Chave do cupom fiscal eletrônico
	NumCcf string    `gorm:"type:varchar(9)"`        // Número do contador de cupom fiscal
	DtIni  time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin  time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj   string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC465) TableName() string {
	return "reg_c465"
}

// Validações dos campos
func (r *RegC465) Validate() error {
	if r.Reg != "C465" {
		return ErrInvalidRegC465
	}
	if r.ChvCfe == "" {
		return ErrEmptyChvCfe
	}
	if len(r.ChvCfe) != 44 {
		return ErrInvalidChvCfe
	}
	if r.NumCcf == "" {
		return ErrEmptyNumCcf
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateC465
	}
	return nil
}

// Interface que define o contrato para criar RegC465
type iRegC465 interface {
	GetRegC465() RegC465
}

// RegC465Sped representa a estrutura do arquivo SPED
type RegC465Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

// GetRegC465 converte a linha do SPED em struct RegC465
func (s RegC465Sped) GetRegC465() RegC465 {
	regC465 := RegC465{
		Reg:    s.Ln[1],
		ChvCfe: s.Ln[2],
		NumCcf: s.Ln[3],
		DtIni:  s.Reg0000.DtIni,
		DtFin:  s.Reg0000.DtFin,
		Cnpj:   s.Reg0000.Cnpj,
	}
	return regC465
}

// CreateRegC465 cria um novo registro RegC465
func CreateRegC465(read iRegC465) (RegC465, error) {
	reg := read.GetRegC465()
	if err := reg.Validate(); err != nil {
		return RegC465{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC465  = errors.New("registro inválido, deve ser C465")
	ErrEmptyChvCfe     = errors.New("chave do CFe não pode ser vazia")
	ErrInvalidChvCfe   = errors.New("chave do CFe deve ter 44 caracteres")
	ErrEmptyNumCcf     = errors.New("número do CCF não pode ser vazio")
	ErrInvalidDateC465 = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *RegC465) GetIdentificacaoCFe() string {
	return "CFe: " + r.ChvCfe + " (CCF: " + r.NumCcf + ")"
}

func (r *RegC465) IsCFeAtivo(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}
