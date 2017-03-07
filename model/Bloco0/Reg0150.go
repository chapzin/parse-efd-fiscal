package Bloco0

import(
	"github.com/jinzhu/gorm"
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
	"time"
)

type Reg0150 struct {
	gorm.Model
	Reg string		`gorm:"type:varchar(4)"`
	CodPart string		`gorm:"type:varchar(60);"`
	Nome string		`gorm:"type:varchar(100)"`
	CodPais string		`gorm:"type:varchar(5)"`
	Cnpj string		`gorm:"type:varchar(15)"`
	Cpf string		`gorm:"type:varchar(11)"`
	Ie string		`gorm:"type:varchar(14)"`
	CodMun string		`gorm:"type:varchar(7)"`
	Suframa string		`gorm:"type:varchar(9)"`
	Endereco string		`gorm:"type:varchar(60)"`
	Num string		`gorm:"type:varchar(10)"`
	Compl string		`gorm:"type:varchar(60)"`
	Bairro string		`gorm:"type:varchar(60)"`
	DtIni time.Time 	`gorm:"type:date"`
	DtFin time.Time 	`gorm:"type:date"`
	CnpjSped string		`gorm:"type:varchar(14)"`

}

func (Reg0150) TableName() string {
	return "reg_0150"
}

type Reg0150Sped struct {
	Ln []string
	Reg0000 Reg0000
}

func (s Reg0150Sped) GetReg() string {
	return s.Ln[1]
}

func (s Reg0150Sped) GetCodPart() string {
	return s.Ln[2]
}

func (s Reg0150Sped) GetNome() string {
	return s.Ln[3]
}

func (s Reg0150Sped) GetCodPais() string {
	return s.Ln[4]
}

func (s Reg0150Sped) GetCnpj() string {
	return s.Ln[5]
}

func (s Reg0150Sped) GetCpf() string {
	return s.Ln[6]
}

func (s Reg0150Sped) GetIe() string {
	return s.Ln[7]
}

func (s Reg0150Sped) GetCodMun() string {
	return s.Ln[8]
}

func (s Reg0150Sped) GetSuframa() string {
	return s.Ln[9]
}

func (s Reg0150Sped) GetEndereco() string {
	return s.Ln[10]
}

func (s Reg0150Sped) GetNum() string {
	return s.Ln[11]
}

func (s Reg0150Sped) GetCompl() string {
	return s.Ln[12]
}

func (s Reg0150Sped) GetBairro() string {
	return s.Ln[13]
}

func (s Reg0150Sped) GetDtIni() time.Time {
	return s.Reg0000.DtIni
}

func (s Reg0150Sped) GetDtFin() time.Time {
	return s.Reg0000.DtFin
}

func (s Reg0150Sped) GetCnpjSped() string  {
	return s.Reg0000.Cnpj
}


type iReg0150 interface {
	GetReg() string
	GetCodPart() string
	GetNome() string
	GetCodPais() string
	GetCnpj() string
	GetCpf() string
	GetIe() string
	GetCodMun() string
	GetSuframa() string
	GetEndereco() string
	GetNum() string
	GetCompl() string
	GetBairro() string
	GetDtIni() time.Time
	GetDtFin() time.Time
	GetCnpjSped() string
}

type Reg0150Xml struct {
	// para usar o cliente e o endereco deve ser mapeado como abaixo
	//nfe, _ := mxj.NewMapXml(xmlFile) << vindo de um ioutil.ReadFile
	//	cliente, err := nfe.ValuesForKey("dest")
	//endereco, err := nfe.ValuesForKey("enderDest")

	Cliente []interface{}
	Endereco []interface{}
}

func (x Reg0150Xml) GetReg() string {
	return "0000"
}

func (x Reg0150Xml) GetCodPart() string {
	cnpj := SpedConvert.DataXml(x.Cliente,"CNPJ")
	return cnpj
}

func (x Reg0150Xml) GetNome() string {
	nome := SpedConvert.DataXml(x.Cliente,"xNome")
	return nome
}

func (x Reg0150Xml) GetCodPais() string {
	return "1058"
}

func (x Reg0150Xml) GetCnpj() string {
	cnpj := SpedConvert.DataXml(x.Cliente,"CNPJ")
	return cnpj
}

func (x Reg0150Xml) GetCpf() string {
	cpf := SpedConvert.DataXml(x.Cliente,"CNPJ")
	return cpf
}

func (x Reg0150Xml) GetIe() string {
	ie := SpedConvert.DataXml(x.Cliente,"IE")
	return ie
}

func (x Reg0150Xml) GetCodMun() string {
	codMun := SpedConvert.DataXml(x.Endereco,"cMun")
	return codMun
}

func (x Reg0150Xml) GetSuframa() string {
	return ""
}

func (x Reg0150Xml) GetEndereco() string {
	endereco := SpedConvert.DataXml(x.Endereco,"xLgr")
	return endereco
}

func (x Reg0150Xml) GetNum() string  {
	num := SpedConvert.DataXml(x.Endereco,"nro")
	return num
}

func (x Reg0150Xml) GetCompl() string  {
	compl := SpedConvert.DataXml(x.Endereco,"xCpl")
	return compl
}

func (x Reg0150Xml) GetBairro() string  {
	bairro := SpedConvert.DataXml(x.Endereco,"xBairro")
	return bairro
}

func (x Reg0150Xml) GetDtIni() time.Time {
	return SpedConvert.ConvertDataNull()
}

func (x Reg0150Xml) GetDtFin() time.Time {
	return SpedConvert.ConvertDataNull()
}

func (x Reg0150Xml) GetCnpjSped() string {
	return ""
}


func CreateReg0150 (read iReg0150) Reg0150 {
	reg0150 := Reg0150{
		Reg:		read.GetReg(),
		CodPart:	read.GetCodPart(),
		Nome:		read.GetNome(),
		CodPais:	read.GetCodPais(),
		Cnpj:		read.GetCnpj(),
		Cpf:		read.GetCpf(),
		Ie:		read.GetIe(),
		CodMun:		read.GetCodMun(),
		Suframa:	read.GetSuframa(),
		Endereco:	read.GetSuframa(),
		Num:		read.GetNum(),
		Compl:		read.GetCompl(),
		Bairro:		read.GetBairro(),
		DtIni:		read.GetDtIni(),
		DtFin:		read.GetDtFin(),
		CnpjSped:	read.GetCnpjSped(),
	}
	return  reg0150
}
