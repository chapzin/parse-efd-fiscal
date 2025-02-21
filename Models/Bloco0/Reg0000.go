package Bloco0

import (
	"errors"
	"time"

	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

// Reg0000 representa o registro 0000 do SPED Fiscal que contém informações de identificação do contribuinte
type Reg0000 struct {
	gorm.Model
	Reg       string    `gorm:"type:varchar(4);index"`  // Texto fixo contendo "0000"
	CodVer    string    `gorm:"type:varchar(3)"`        // Código da versão do layout
	CodFin    int       `gorm:"type:int"`               // Código da finalidade do arquivo
	DtIni     time.Time `gorm:"type:date;index"`        // Data inicial das informações
	DtFin     time.Time `gorm:"type:date;index"`        // Data final das informações
	Nome      string    `gorm:"type:varchar(100)"`      // Nome empresarial do contribuinte
	Cnpj      string    `gorm:"type:varchar(14);index"` // CNPJ do contribuinte
	Cpf       string    `gorm:"type:varchar(11)"`       // CPF do contribuinte
	Uf        string    `gorm:"type:varchar(2)"`        // Sigla da unidade da federação
	Ie        string    `gorm:"type:varchar(14);index"` // Inscrição Estadual do contribuinte
	CodMun    string    `gorm:"type:varchar(7)"`        // Código do município
	Im        string    `gorm:"type:varchar(15)"`       // Inscrição Municipal
	Suframa   string    `gorm:"type:varchar(9)"`        // Inscrição SUFRAMA
	IndPerfil string    `gorm:"type:varchar(1)"`        // Perfil de apresentação do arquivo fiscal
	IndAtiv   int       `gorm:"type:int"`               // Indicador de tipo de atividade
}

// TableName retorna o nome da tabela no banco de dados
func (Reg0000) TableName() string {
	return "reg_0000"
}

// Reg0000Sped representa a estrutura do arquivo SPED
type Reg0000Sped struct {
	Ln []string
}

// Interface que define o contrato para criar Reg0000
type iReg0000 interface {
	GetReg0000() Reg0000
}

// Validações dos campos
func (r *Reg0000) Validate() error {
	if r.Reg != "0000" {
		return ErrInvalidReg
	}
	if r.CodVer == "" {
		return ErrEmptyCodVer
	}
	if r.DtIni.After(r.DtFin) {
		return ErrInvalidDateRange
	}
	if r.Cnpj == "" && r.Cpf == "" {
		return ErrMissingIdentification
	}
	return nil
}

// GetReg0000 converte a linha do SPED em struct Reg0000
func (s Reg0000Sped) GetReg0000() Reg0000 {
	reg0000 := Reg0000{
		Reg:       s.Ln[1],
		CodVer:    s.Ln[2],
		CodFin:    tools.ConvInt(s.Ln[3]),
		DtIni:     tools.ConvertData(s.Ln[4]),
		DtFin:     tools.ConvertData(s.Ln[5]),
		Nome:      s.Ln[6],
		Cnpj:      s.Ln[7],
		Cpf:       s.Ln[8],
		Uf:        s.Ln[9],
		Ie:        s.Ln[10],
		CodMun:    s.Ln[11],
		Im:        s.Ln[12],
		Suframa:   s.Ln[13],
		IndPerfil: s.Ln[14],
		IndAtiv:   tools.ConvInt(s.Ln[15]),
	}

	return reg0000
}

// CreateReg0000 cria um novo registro Reg0000
func CreateReg0000(read iReg0000) (Reg0000, error) {
	reg := read.GetReg0000()
	if err := reg.Validate(); err != nil {
		return Reg0000{}, err
	}
	return reg, nil
}

// Constantes de erro
var (
	ErrInvalidReg            = errors.New("registro inválido, deve ser 0000")
	ErrEmptyCodVer           = errors.New("código da versão não pode ser vazio")
	ErrInvalidDateRange      = errors.New("data inicial não pode ser posterior à data final")
	ErrMissingIdentification = errors.New("CNPJ ou CPF deve ser informado")
)

// Métodos auxiliares
func (r *Reg0000) IsAtivo() bool {
	return r.IndAtiv == 0
}

func (r *Reg0000) IsPerfil(perfil string) bool {
	return r.IndPerfil == perfil
}

func (r *Reg0000) GetPeriodo() string {
	return r.DtIni.Format("200601")
}
