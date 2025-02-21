package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC800 representa o registro C800 do SPED Fiscal que contém os documentos emitidos por equipamento SAT-CF-e
type RegC800 struct {
	gorm.Model
	Reg        string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "C800"
	CodMod     string    `gorm:"type:varchar(2)"`        // Código do modelo do documento fiscal (59 - CF-e-SAT)
	CodSit     string    `gorm:"type:varchar(2)"`        // Código da situação do documento fiscal
	NumCfe     string    `gorm:"type:varchar(6);index"`  // Número do Cupom Fiscal Eletrônico
	DtDoc      time.Time `gorm:"type:date;index"`        // Data da emissão do Cupom Fiscal Eletrônico
	VlCfe      float64   `gorm:"type:decimal(19,2)"`     // Valor total do Cupom Fiscal Eletrônico
	VlPis      float64   `gorm:"type:decimal(19,2)"`     // Valor total do PIS
	VlCofins   float64   `gorm:"type:decimal(19,2)"`     // Valor total da COFINS
	CnpjCpf    string    `gorm:"type:varchar(14)"`       // CNPJ ou CPF do destinatário
	NrSat      string    `gorm:"type:varchar(9);index"`  // Número de Série do equipamento SAT
	ChvCfe     string    `gorm:"type:varchar(44);index"` // Chave do Cupom Fiscal Eletrônico
	VlDesc     float64   `gorm:"type:decimal(19,2)"`     // Valor total de descontos
	VlMerc     float64   `gorm:"type:decimal(19,2)"`     // Valor total das mercadorias e serviços
	VlOutDa    float64   `gorm:"type:decimal(19,2)"`     // Valor total de outras despesas acessórias
	VlIcms     float64   `gorm:"type:decimal(19,2)"`     // Valor do ICMS
	VlPisSt    float64   `gorm:"type:decimal(19,2)"`     // Valor total do PIS retido por substituição tributária
	VlCofinsSt float64   `gorm:"type:decimal(19,2)"`     // Valor total da COFINS retido por substituição tributária
	DtIni      time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin      time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj       string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC800) TableName() string {
	return "reg_c800"
}

// Validações dos campos
func (r *RegC800) Validate() error {
	if r.Reg != "C800" {
		return ErrInvalidRegC800
	}
	if r.CodMod != "59" {
		return ErrInvalidCodModC800
	}
	if !tools.Contains([]string{"00", "01", "02", "03"}, r.CodSit) {
		return ErrInvalidCodSitC800
	}
	if r.NumCfe == "" {
		return ErrEmptyNumCfeC800
	}
	if r.NrSat == "" {
		return ErrEmptyNrSatC800
	}
	if r.ChvCfe == "" || len(r.ChvCfe) != 44 {
		return ErrInvalidChvCfeC800
	}
	if r.VlCfe < 0 {
		return ErrInvalidVlCfeC800
	}
	if r.DtDoc.After(r.DtFin) {
		return ErrInvalidDateC800
	}
	return nil
}

// Interface que define o contrato para criar RegC800
type iRegC800 interface {
	GetRegC800() RegC800
}

// RegC800Sped representa a estrutura do arquivo SPED
type RegC800Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

// GetRegC800 converte a linha do SPED em struct RegC800
func (s RegC800Sped) GetRegC800() RegC800 {
	regC800 := RegC800{
		Reg:        s.Ln[1],
		CodMod:     s.Ln[2],
		CodSit:     s.Ln[3],
		NumCfe:     s.Ln[4],
		DtDoc:      tools.ConvertData(s.Ln[5]),
		VlCfe:      tools.ConvFloat(s.Ln[6]),
		VlPis:      tools.ConvFloat(s.Ln[7]),
		VlCofins:   tools.ConvFloat(s.Ln[8]),
		CnpjCpf:    s.Ln[9],
		NrSat:      s.Ln[10],
		ChvCfe:     s.Ln[11],
		VlDesc:     tools.ConvFloat(s.Ln[12]),
		VlMerc:     tools.ConvFloat(s.Ln[13]),
		VlOutDa:    tools.ConvFloat(s.Ln[14]),
		VlIcms:     tools.ConvFloat(s.Ln[15]),
		VlPisSt:    tools.ConvFloat(s.Ln[16]),
		VlCofinsSt: tools.ConvFloat(s.Ln[17]),
		DtIni:      s.Reg0000.DtIni,
		DtFin:      s.Reg0000.DtFin,
		Cnpj:       s.Reg0000.Cnpj,
	}
	return regC800
}

// CreateRegC800 cria um novo registro RegC800
func CreateRegC800(read iRegC800) (RegC800, error) {
	reg := read.GetRegC800()
	if err := reg.Validate(); err != nil {
		return RegC800{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC800    = errors.New("registro inválido, deve ser C800")
	ErrInvalidCodModC800 = errors.New("código do modelo deve ser 59 (CF-e-SAT)")
	ErrInvalidCodSitC800 = errors.New("código da situação deve ser 00, 01, 02 ou 03")
	ErrEmptyNumCfeC800   = errors.New("número do CF-e não pode ser vazio")
	ErrEmptyNrSatC800    = errors.New("número do SAT não pode ser vazio")
	ErrInvalidChvCfeC800 = errors.New("chave do CF-e inválida")
	ErrInvalidVlCfeC800  = errors.New("valor do CF-e não pode ser negativo")
	ErrInvalidDateC800   = errors.New("data do documento não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *RegC800) GetValorTotal() float64 {
	return r.VlCfe
}

func (r *RegC800) GetValorLiquido() float64 {
	return r.VlCfe - r.VlDesc
}

func (r *RegC800) GetValorTributos() float64 {
	return r.VlIcms + r.VlPis + r.VlCofins + r.VlPisSt + r.VlCofinsSt
}

func (r *RegC800) IsCancelado() bool {
	return r.CodSit == "02" || r.CodSit == "03"
}

func (r *RegC800) GetIdentificacaoSAT() string {
	return "SAT nº " + r.NrSat + " - CF-e nº " + r.NumCfe
}

func (r *RegC800) IsCFeAtivo(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}
