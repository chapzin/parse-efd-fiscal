package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/jinzhu/gorm"
)

// RegC400 representa o registro C400 do SPED Fiscal que contém os equipamentos ECF
type RegC400 struct {
	gorm.Model
	Reg    string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C400"
	CodMod string    `gorm:"type:varchar(2)"`        // Código do modelo do documento fiscal
	EcfMod string    `gorm:"type:varchar(20)"`       // Modelo do equipamento
	EcfFab string    `gorm:"type:varchar(21)"`       // Número de fabricação do ECF
	EcfCx  string    `gorm:"type:varchar(3)"`        // Número do caixa atribuído ao ECF
	DtIni  time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin  time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj   string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC400) TableName() string {
	return "reg_c400"
}

// Validações dos campos
func (r *RegC400) Validate() error {
	if r.Reg != "C400" {
		return ErrInvalidRegC400
	}
	if r.CodMod == "" {
		return ErrEmptyCodMod
	}
	if r.EcfMod == "" {
		return ErrEmptyEcfMod
	}
	if r.EcfFab == "" {
		return ErrEmptyEcfFab
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateC400
	}
	return nil
}

// Interface que define o contrato para criar RegC400
type iRegC400 interface {
	GetRegC400() RegC400
}

// RegC400Sped representa a estrutura do arquivo SPED
type RegC400Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

// GetRegC400 converte a linha do SPED em struct RegC400
func (s RegC400Sped) GetRegC400() RegC400 {
	regC400 := RegC400{
		Reg:    s.Ln[1],
		CodMod: s.Ln[2],
		EcfMod: s.Ln[3],
		EcfFab: s.Ln[4],
		EcfCx:  s.Ln[5],
		DtIni:  s.Reg0000.DtIni,
		DtFin:  s.Reg0000.DtFin,
		Cnpj:   s.Reg0000.Cnpj,
	}
	return regC400
}

// CreateRegC400 cria um novo registro RegC400
func CreateRegC400(read iRegC400) (RegC400, error) {
	reg := read.GetRegC400()
	if err := reg.Validate(); err != nil {
		return RegC400{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC400  = errors.New("registro inválido, deve ser C400")
	ErrEmptyCodMod     = errors.New("código do modelo não pode ser vazio")
	ErrEmptyEcfMod     = errors.New("modelo do ECF não pode ser vazio")
	ErrEmptyEcfFab     = errors.New("número de fabricação não pode ser vazio")
	ErrInvalidDateC400 = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *RegC400) GetIdentificacaoECF() string {
	return r.EcfMod + " - " + r.EcfFab + " (CX: " + r.EcfCx + ")"
}

func (r *RegC400) IsECFAtivo(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}
