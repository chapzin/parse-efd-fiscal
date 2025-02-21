package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC860 representa o registro C860 do SPED Fiscal que identifica o equipamento SAT-CF-e
type RegC860 struct {
	gorm.Model
	Reg     string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C860"
	CodMod  string    `gorm:"type:varchar(2)"`        // Código do modelo do documento fiscal (59 - CF-e-SAT)
	NrSat   string    `gorm:"type:varchar(9);index"`  // Número de Série do equipamento SAT
	DtDoc   time.Time `gorm:"type:date;index"`        // Data de emissão dos documentos fiscais
	DocIni  string    `gorm:"type:varchar(6)"`        // Número do documento inicial
	DocFim  string    `gorm:"type:varchar(6)"`        // Número do documento final
	DtIni   time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin   time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj    string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC860) TableName() string {
	return "reg_c860"
}

// Validações dos campos
func (r *RegC860) Validate() error {
	if r.Reg != "C860" {
		return ErrInvalidRegC860
	}
	if r.CodMod != "59" {
		return ErrInvalidCodModC860
	}
	if r.NrSat == "" {
		return ErrEmptyNrSatC860
	}
	if r.DocIni == "" {
		return ErrEmptyDocIniC860
	}
	if r.DocFim == "" {
		return ErrEmptyDocFimC860
	}
	if tools.ConvInt(r.DocIni) > tools.ConvInt(r.DocFim) {
		return ErrInvalidDocRange
	}
	if r.DtDoc.After(r.DtFin) || r.DtDoc.Before(r.DtIni) {
		return ErrInvalidDateC860
	}
	return nil
}

// Interface que define o contrato para criar RegC860
type iRegC860 interface {
	GetRegC860() RegC860
}

// RegC860Sped representa a estrutura do arquivo SPED
type RegC860Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

// GetRegC860 converte a linha do SPED em struct RegC860
func (s RegC860Sped) GetRegC860() RegC860 {
	regC860 := RegC860{
		Reg:    s.Ln[1],
		CodMod: s.Ln[2],
		NrSat:  s.Ln[3],
		DtDoc:  tools.ConvertData(s.Ln[4]),
		DocIni: s.Ln[5],
		DocFim: s.Ln[6],
		DtIni:  s.Reg0000.DtIni,
		DtFin:  s.Reg0000.DtFin,
		Cnpj:   s.Reg0000.Cnpj,
	}
	return regC860
}

// CreateRegC860 cria um novo registro RegC860
func CreateRegC860(read iRegC860) (RegC860, error) {
	reg := read.GetRegC860()
	if err := reg.Validate(); err != nil {
		return RegC860{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC860   = errors.New("registro inválido, deve ser C860")
	ErrInvalidCodModC860 = errors.New("código do modelo deve ser 59 (CF-e-SAT)")
	ErrEmptyNrSatC860    = errors.New("número do SAT não pode ser vazio")
	ErrEmptyDocIniC860   = errors.New("número do documento inicial não pode ser vazio")
	ErrEmptyDocFimC860   = errors.New("número do documento final não pode ser vazio")
	ErrInvalidDocRange   = errors.New("documento inicial não pode ser maior que o documento final")
	ErrInvalidDateC860   = errors.New("data do documento deve estar dentro do período de apuração")
)

// Métodos auxiliares
func (r *RegC860) GetIdentificacaoSAT() string {
	return "SAT nº " + r.NrSat + " - Docs: " + r.DocIni + " a " + r.DocFim
}

func (r *RegC860) GetQuantidadeDocumentos() int {
	return tools.ConvInt(r.DocFim) - tools.ConvInt(r.DocIni) + 1
}

func (r *RegC860) IsEquipamentoAtivo(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
} 