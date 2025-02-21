package BlocoC

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// RegC100 representa o registro C100 do SPED Fiscal que contém as notas fiscais
type RegC100 struct {
	gorm.Model
	Reg        string    `gorm:"type:varchar(4);index"`         // Texto fixo contendo "C100"
	IndOper    string    `gorm:"type:varchar(1)"`               // Indicador de operação (0-Entrada, 1-Saída)
	IndEmit    string    `gorm:"type:varchar(1)"`               // Indicador do emitente
	CodPart    string    `gorm:"type:varchar(60);index"`        // Código do participante
	CodMod     string    `gorm:"type:varchar(2)"`               // Código do modelo do documento fiscal
	CodSit     string    `gorm:"type:varchar(2)"`               // Código da situação do documento fiscal
	Ser        string    `gorm:"type:varchar(3)"`               // Série do documento fiscal
	NumDoc     string    `gorm:"type:varchar(9);index"`         // Número do documento fiscal
	ChvNfe     string    `gorm:"type:varchar(44);unique_index"` // Chave da NFe
	DtDoc      time.Time `gorm:"type:date;index"`               // Data da emissão do documento fiscal
	DtES       time.Time `gorm:"type:date;index"`               // Data da entrada ou saída
	VlDoc      float64   `gorm:"type:decimal(19,2)"`            // Valor total do documento fiscal
	IndPgto    string    `gorm:"type:varchar(1)"`               // Indicador do tipo de pagamento
	VlDesc     float64   `gorm:"type:decimal(19,2)"`            // Valor total do desconto
	VlAbatNt   float64   `gorm:"type:decimal(19,2)"`            // Valor total do abatimento
	VlMerc     float64   `gorm:"type:decimal(19,2)"`            // Valor total das mercadorias
	IndFrt     string    `gorm:"type:varchar(1)"`               // Indicador do tipo do frete
	VlFrt      float64   `gorm:"type:decimal(19,2)"`            // Valor do frete
	VlSeg      float64   `gorm:"type:decimal(19,2)"`            // Valor do seguro
	VlOutDa    float64   `gorm:"type:decimal(19,2)"`            // Valor de outras despesas
	VlBcIcms   float64   `gorm:"type:decimal(19,2)"`            // Base de cálculo do ICMS
	VlIcms     float64   `gorm:"type:decimal(19,2)"`            // Valor do ICMS
	VlBcIcmsSt float64   `gorm:"type:decimal(19,2)"`            // Base de cálculo do ICMS ST
	VlIcmsSt   float64   `gorm:"type:decimal(19,2)"`            // Valor do ICMS ST
	VlIpi      float64   `gorm:"type:decimal(19,2)"`            // Valor do IPI
	VlPis      float64   `gorm:"type:decimal(19,2)"`            // Valor do PIS
	VlCofins   float64   `gorm:"type:decimal(19,2)"`            // Valor do COFINS
	VlPisSt    float64   `gorm:"type:decimal(19,2)"`            // Valor do PIS ST
	VlCofinsSt float64   `gorm:"type:decimal(19,2)"`            // Valor do COFINS ST
	DtIni      time.Time `gorm:"type:date;index"`               // Data inicial das informações
	DtFin      time.Time `gorm:"type:date;index"`               // Data final das informações
	Cnpj       string    `gorm:"type:varchar(14);index"`        // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (RegC100) TableName() string {
	return "reg_c100"
}

// Validações dos campos
func (r *RegC100) Validate() error {
	if r.Reg != "C100" {
		return ErrInvalidRegC100
	}
	if r.IndOper != "0" && r.IndOper != "1" {
		return ErrInvalidIndOper
	}
	if r.NumDoc == "" {
		return ErrEmptyNumDoc
	}
	if r.DtDoc.After(r.DtFin) {
		return ErrInvalidDateC100
	}
	return nil
}

// Interface que define o contrato para criar RegC100
type iRegC100 interface {
	GetRegC100() RegC100
}

// RegC100Sped representa a estrutura do arquivo SPED
type RegC100Sped struct {
	Ln      []string
	Reg0000 Bloco0.Reg0000
}

// GetRegC100 converte a linha do SPED em struct RegC100
func (s RegC100Sped) GetRegC100() RegC100 {
	regC100 := RegC100{
		Reg:        s.Ln[1],
		IndOper:    s.Ln[2],
		IndEmit:    s.Ln[3],
		CodPart:    s.Ln[4],
		CodMod:     s.Ln[5],
		CodSit:     s.Ln[6],
		Ser:        s.Ln[7],
		NumDoc:     s.Ln[8],
		ChvNfe:     s.Ln[9],
		DtDoc:      tools.ConvertData(s.Ln[10]),
		DtES:       tools.ConvertData(s.Ln[11]),
		VlDoc:      tools.ConvFloat(s.Ln[12]),
		IndPgto:    s.Ln[13],
		VlDesc:     tools.ConvFloat(s.Ln[14]),
		VlAbatNt:   tools.ConvFloat(s.Ln[15]),
		VlMerc:     tools.ConvFloat(s.Ln[16]),
		IndFrt:     s.Ln[17],
		VlFrt:      tools.ConvFloat(s.Ln[18]),
		VlSeg:      tools.ConvFloat(s.Ln[19]),
		VlOutDa:    tools.ConvFloat(s.Ln[20]),
		VlBcIcms:   tools.ConvFloat(s.Ln[21]),
		VlIcms:     tools.ConvFloat(s.Ln[22]),
		VlBcIcmsSt: tools.ConvFloat(s.Ln[23]),
		VlIcmsSt:   tools.ConvFloat(s.Ln[24]),
		VlIpi:      tools.ConvFloat(s.Ln[25]),
		VlPis:      tools.ConvFloat(s.Ln[26]),
		VlCofins:   tools.ConvFloat(s.Ln[27]),
		VlPisSt:    tools.ConvFloat(s.Ln[29]),
		VlCofinsSt: tools.ConvFloat(s.Ln[30]),
		DtIni:      s.Reg0000.DtIni,
		DtFin:      s.Reg0000.DtFin,
		Cnpj:       s.Reg0000.Cnpj,
	}
	return regC100
}

// CreateRegC100 cria um novo registro RegC100
func CreateRegC100(read iRegC100) (RegC100, error) {
	reg := read.GetRegC100()
	if err := reg.Validate(); err != nil {
		return RegC100{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidRegC100  = errors.New("registro inválido, deve ser C100")
	ErrInvalidIndOper  = errors.New("indicador de operação inválido")
	ErrEmptyNumDoc     = errors.New("número do documento não pode ser vazio")
	ErrInvalidDateC100 = errors.New("data do documento não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *RegC100) IsEntrada() bool {
	return r.IndOper == "0"
}

func (r *RegC100) IsSaida() bool {
	return r.IndOper == "1"
}

func (r *RegC100) IsCancelada() bool {
	return r.CodSit == "02"
}

func (r *RegC100) GetValorTotal() float64 {
	return r.VlDoc
}

func (r *RegC100) GetValorLiquido() float64 {
	return r.VlDoc - r.VlDesc - r.VlAbatNt
}
