package Bloco0

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// Reg0200 representa o registro 0200 do SPED Fiscal que contém a identificação dos itens (produtos e serviços)
type Reg0200 struct {
	gorm.Model
	Reg        string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "0200"
	CodItem    string    `gorm:"type:varchar(60);index"` // Código do item
	DescrItem  string    `gorm:"type:varchar(255)"`      // Descrição do item
	CodBarra   string    `gorm:"type:varchar(14)"`       // Código de barra do item
	CodAntItem string    `gorm:"type:varchar(60)"`       // Código anterior do item
	UnidInv    string    `gorm:"type:varchar(6)"`        // Unidade de medida do estoque
	TipoItem   string    `gorm:"type:varchar(2)"`        // Tipo do item
	CodNcm     string    `gorm:"type:varchar(8);index"`  // Código NCM
	ExIpi      string    `gorm:"type:varchar(3)"`        // Código EX da TIPI
	CodGen     string    `gorm:"type:varchar(2)"`        // Código do gênero do item
	CodLst     string    `gorm:"type:varchar(5)"`        // Código do serviço (LC 116/2003)
	AliqIcms   float64   `gorm:"type:decimal(6,2)"`      // Alíquota de ICMS aplicável
	DtIni      time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin      time.Time `gorm:"type:date;index"`        // Data final das informações
	Cnpj       string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
}

// TableName retorna o nome da tabela no banco de dados
func (Reg0200) TableName() string {
	return "reg_0200"
}

// Validações dos campos
func (r *Reg0200) Validate() error {
	if r.Reg != "0200" {
		return ErrInvalidReg0200
	}
	if r.CodItem == "" {
		return ErrEmptyCodItem
	}
	if r.DescrItem == "" {
		return ErrEmptyDescrItem
	}
	if r.UnidInv == "" {
		return ErrEmptyUnidInv
	}
	if r.TipoItem == "" {
		return ErrEmptyTipoItem
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateRange0200
	}
	return nil
}

// Interface que define o contrato para criar Reg0200
type iReg0200 interface {
	GetReg0200() Reg0200
}

// Reg0200Sped representa a estrutura do arquivo SPED
type Reg0200Sped struct {
	Ln      []string
	Reg0000 Reg0000
	Digito  string
}

// GetReg0200 converte a linha do SPED em struct Reg0200
func (s Reg0200Sped) GetReg0200() Reg0200 {
	digitoInt := tools.ConvInt(s.Digito)
	codigo := tools.AdicionaDigitosCodigo(s.Ln[2], digitoInt)
	reg0200 := Reg0200{
		Reg:        s.Ln[1],
		CodItem:    codigo,
		DescrItem:  s.Ln[3],
		CodBarra:   s.Ln[4],
		CodAntItem: s.Ln[5],
		UnidInv:    s.Ln[6],
		TipoItem:   s.Ln[7],
		CodNcm:     s.Ln[8],
		ExIpi:      s.Ln[9],
		CodGen:     s.Ln[10],
		CodLst:     s.Ln[11],
		AliqIcms:   tools.ConvFloat(s.Ln[12]),
		DtIni:      s.Reg0000.DtIni,
		DtFin:      s.Reg0000.DtFin,
		Cnpj:       s.Reg0000.Cnpj,
	}
	return reg0200
}

// CreateReg0200 cria um novo registro Reg0200
func CreateReg0200(read iReg0200) (Reg0200, error) {
	reg := read.GetReg0200()
	if err := reg.Validate(); err != nil {
		return Reg0200{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidReg0200       = errors.New("registro inválido, deve ser 0200")
	ErrEmptyCodItem         = errors.New("código do item não pode ser vazio")
	ErrEmptyDescrItem       = errors.New("descrição do item não pode ser vazia")
	ErrEmptyUnidInv         = errors.New("unidade de inventário não pode ser vazia")
	ErrEmptyTipoItem        = errors.New("tipo do item não pode ser vazio")
	ErrInvalidDateRange0200 = errors.New("data inicial não pode ser posterior à data final")
)

// Métodos auxiliares
func (r *Reg0200) IsItemAtivo(data time.Time) bool {
	return !data.Before(r.DtIni) && !data.After(r.DtFin)
}

func (r *Reg0200) IsMercadoriaRevenda() bool {
	return r.TipoItem == "00"
}

func (r *Reg0200) IsMateriaPrima() bool {
	return r.TipoItem == "01"
}

func (r *Reg0200) IsEmbalagem() bool {
	return r.TipoItem == "02"
}

func (r *Reg0200) IsProdutoEmProcesso() bool {
	return r.TipoItem == "03"
}

func (r *Reg0200) IsProdutoAcabado() bool {
	return r.TipoItem == "04"
}

func (r *Reg0200) IsSubproduto() bool {
	return r.TipoItem == "05"
}

func (r *Reg0200) GetDescricaoCompleta() string {
	descr := r.DescrItem
	if r.CodBarra != "" {
		descr += " (EAN: " + r.CodBarra + ")"
	}
	if r.CodNcm != "" {
		descr += " [NCM: " + r.CodNcm + "]"
	}
	return descr
}
