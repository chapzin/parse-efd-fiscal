// Package layout define catálogos versionados de registros/campos EFD.
package layout

import "fmt"

type FieldType string

const (
	FieldFixed   FieldType = "fixed"
	FieldString  FieldType = "string"
	FieldInteger FieldType = "integer"
	FieldDecimal FieldType = "decimal"
	FieldDate    FieldType = "date"
	FieldCNPJ    FieldType = "cnpj"
	FieldCPF     FieldType = "cpf"
	FieldUF      FieldType = "uf"
	FieldCFOP    FieldType = "cfop"
	FieldNCM     FieldType = "ncm"
	FieldChNFe   FieldType = "chave_nfe"
)

type FieldSpec struct {
	Index    int
	Name     string
	Type     FieldType
	Required bool
	Literal  string
	MaxLen   int
}

type RecordSpec struct {
	Code        string
	Name        string
	Block       string
	MinFields   int
	MaxFields   int
	Implemented bool
	Fields      []FieldSpec
}

type Catalog struct {
	ID           string
	GuideVersion string
	Description  string
	Records      map[string]RecordSpec
}

func (c Catalog) Record(code string) (RecordSpec, bool) {
	rec, ok := c.Records[code]
	return rec, ok
}

func (r RecordSpec) FieldByIndex(index int) (FieldSpec, bool) {
	for _, field := range r.Fields {
		if field.Index == index {
			return field, true
		}
	}
	return FieldSpec{}, false
}

func MustDefaultCatalog() Catalog {
	catalog, err := DefaultCatalog()
	if err != nil {
		panic(err)
	}
	return catalog
}

func DefaultCatalog() (Catalog, error) {
	catalog := Catalog{
		ID:           "legacy-minimum",
		GuideVersion: "2.0.20+roadmap",
		Description:  "Catálogo mínimo inicial baseado nos registros já implementados no projeto e nos registros de controle do Bloco 9.",
		Records:      map[string]RecordSpec{},
	}
	for _, record := range []RecordSpec{
		reg0000(), reg0150(), reg0190(), reg0200(), reg0220(),
		regC100(), regC170(), regH005(), regH010(),
		control("9001", "Abertura do Bloco 9", "9", 2),
		control("9900", "Registros do arquivo", "9", 3),
		control("9990", "Encerramento do Bloco 9", "9", 2),
		control("9999", "Encerramento do arquivo digital", "9", 2),
	} {
		if _, exists := catalog.Records[record.Code]; exists {
			return Catalog{}, fmt.Errorf("registro duplicado no catálogo: %s", record.Code)
		}
		catalog.Records[record.Code] = record
	}
	return catalog, nil
}

func fixed(index int, code string) FieldSpec {
	return FieldSpec{Index: index, Name: "REG", Type: FieldFixed, Required: true, Literal: code, MaxLen: len(code)}
}

func field(index int, name string, typ FieldType, required bool, maxLen int) FieldSpec {
	return FieldSpec{Index: index, Name: name, Type: typ, Required: required, MaxLen: maxLen}
}

func control(code, name, block string, fields int) RecordSpec {
	return RecordSpec{Code: code, Name: name, Block: block, MinFields: fields, MaxFields: fields, Implemented: true, Fields: []FieldSpec{fixed(1, code)}}
}

func reg0000() RecordSpec {
	return RecordSpec{Code: "0000", Name: "Abertura do arquivo digital e identificação da entidade", Block: "0", MinFields: 15, MaxFields: 15, Implemented: true, Fields: []FieldSpec{
		fixed(1, "0000"), field(2, "COD_VER", FieldString, true, 3), field(3, "COD_FIN", FieldInteger, true, 1), field(4, "DT_INI", FieldDate, true, 8), field(5, "DT_FIN", FieldDate, true, 8), field(6, "NOME", FieldString, true, 100), field(7, "CNPJ", FieldCNPJ, false, 14), field(8, "CPF", FieldCPF, false, 11), field(9, "UF", FieldUF, true, 2), field(10, "IE", FieldString, true, 14), field(11, "COD_MUN", FieldString, true, 7), field(12, "IM", FieldString, false, 15), field(13, "SUFRAMA", FieldString, false, 9), field(14, "IND_PERFIL", FieldString, true, 1), field(15, "IND_ATIV", FieldInteger, true, 1),
	}}
}

func reg0150() RecordSpec {
	return RecordSpec{Code: "0150", Name: "Tabela de cadastro do participante", Block: "0", MinFields: 13, MaxFields: 13, Implemented: true, Fields: []FieldSpec{
		fixed(1, "0150"), field(2, "COD_PART", FieldString, true, 60), field(3, "NOME", FieldString, true, 100), field(4, "COD_PAIS", FieldString, false, 5), field(5, "CNPJ", FieldCNPJ, false, 14), field(6, "CPF", FieldCPF, false, 11), field(7, "IE", FieldString, false, 14), field(8, "COD_MUN", FieldString, false, 7), field(9, "SUFRAMA", FieldString, false, 9), field(10, "END", FieldString, false, 60), field(11, "NUM", FieldString, false, 10), field(12, "COMPL", FieldString, false, 60), field(13, "BAIRRO", FieldString, false, 60),
	}}
}

func reg0190() RecordSpec {
	return RecordSpec{Code: "0190", Name: "Identificação das unidades de medida", Block: "0", MinFields: 3, MaxFields: 3, Implemented: true, Fields: []FieldSpec{fixed(1, "0190"), field(2, "UNID", FieldString, true, 6), field(3, "DESCR", FieldString, true, 60)}}
}

func reg0200() RecordSpec {
	return RecordSpec{Code: "0200", Name: "Tabela de identificação do item", Block: "0", MinFields: 13, MaxFields: 13, Implemented: true, Fields: []FieldSpec{fixed(1, "0200"), field(2, "COD_ITEM", FieldString, true, 60), field(3, "DESCR_ITEM", FieldString, true, 255), field(4, "COD_BARRA", FieldString, false, 60), field(5, "COD_ANT_ITEM", FieldString, false, 60), field(6, "UNID_INV", FieldString, true, 6), field(7, "TIPO_ITEM", FieldString, true, 2), field(8, "COD_NCM", FieldNCM, false, 8), field(9, "EX_IPI", FieldString, false, 3), field(10, "COD_GEN", FieldString, false, 2), field(11, "COD_LST", FieldString, false, 5), field(12, "ALIQ_ICMS", FieldDecimal, false, 8), field(13, "CEST", FieldString, false, 7)}}
}

func reg0220() RecordSpec {
	return RecordSpec{Code: "0220", Name: "Fatores de conversão de unidades", Block: "0", MinFields: 3, MaxFields: 4, Implemented: true, Fields: []FieldSpec{fixed(1, "0220"), field(2, "UNID_CONV", FieldString, true, 6), field(3, "FAT_CONV", FieldDecimal, true, 20), field(4, "COD_BARRA", FieldString, false, 60)}}
}

func regC100() RecordSpec {
	return RecordSpec{Code: "C100", Name: "Documento fiscal", Block: "C", MinFields: 29, MaxFields: 29, Implemented: true, Fields: []FieldSpec{fixed(1, "C100"), field(2, "IND_OPER", FieldString, true, 1), field(3, "IND_EMIT", FieldString, true, 1), field(4, "COD_PART", FieldString, false, 60), field(5, "COD_MOD", FieldString, true, 2), field(6, "COD_SIT", FieldString, true, 2), field(7, "SER", FieldString, false, 3), field(8, "NUM_DOC", FieldString, true, 9), field(9, "CHV_NFE", FieldChNFe, false, 44), field(10, "DT_DOC", FieldDate, true, 8), field(11, "DT_E_S", FieldDate, false, 8), field(12, "VL_DOC", FieldDecimal, true, 20), field(13, "IND_PGTO", FieldString, false, 1), field(14, "VL_DESC", FieldDecimal, false, 20), field(15, "VL_ABAT_NT", FieldDecimal, false, 20), field(16, "VL_MERC", FieldDecimal, false, 20), field(17, "IND_FRT", FieldString, false, 1), field(18, "VL_FRT", FieldDecimal, false, 20), field(19, "VL_SEG", FieldDecimal, false, 20), field(20, "VL_OUT_DA", FieldDecimal, false, 20), field(21, "VL_BC_ICMS", FieldDecimal, false, 20), field(22, "VL_ICMS", FieldDecimal, false, 20), field(23, "VL_BC_ICMS_ST", FieldDecimal, false, 20), field(24, "VL_ICMS_ST", FieldDecimal, false, 20), field(25, "VL_IPI", FieldDecimal, false, 20), field(26, "VL_PIS", FieldDecimal, false, 20), field(27, "VL_COFINS", FieldDecimal, false, 20), field(28, "VL_PIS_ST", FieldDecimal, false, 20), field(29, "VL_COFINS_ST", FieldDecimal, false, 20)}}
}

func regC170() RecordSpec {
	return RecordSpec{Code: "C170", Name: "Itens do documento", Block: "C", MinFields: 38, MaxFields: 38, Implemented: true, Fields: []FieldSpec{fixed(1, "C170"), field(2, "NUM_ITEM", FieldInteger, true, 3), field(3, "COD_ITEM", FieldString, true, 60), field(4, "DESCR_COMPL", FieldString, false, 255), field(5, "QTD", FieldDecimal, true, 20), field(6, "UNID", FieldString, true, 6), field(7, "VL_ITEM", FieldDecimal, true, 20), field(8, "VL_DESC", FieldDecimal, false, 20), field(9, "IND_MOV", FieldString, true, 1), field(10, "CST_ICMS", FieldString, true, 3), field(11, "CFOP", FieldCFOP, true, 4), field(12, "NAT_BC_CRED", FieldString, false, 2), field(13, "VL_BC_ICMS", FieldDecimal, false, 20), field(14, "ALIQ_ICMS", FieldDecimal, false, 8), field(15, "VL_ICMS", FieldDecimal, false, 20)}}
}

func regH005() RecordSpec {
	return RecordSpec{Code: "H005", Name: "Totais do inventário", Block: "H", MinFields: 4, MaxFields: 4, Implemented: true, Fields: []FieldSpec{fixed(1, "H005"), field(2, "DT_INV", FieldDate, true, 8), field(3, "VL_INV", FieldDecimal, true, 20), field(4, "MOT_INV", FieldString, true, 2)}}
}

func regH010() RecordSpec {
	return RecordSpec{Code: "H010", Name: "Inventário", Block: "H", MinFields: 10, MaxFields: 10, Implemented: true, Fields: []FieldSpec{fixed(1, "H010"), field(2, "COD_ITEM", FieldString, true, 60), field(3, "UNID", FieldString, true, 6), field(4, "QTD", FieldDecimal, true, 20), field(5, "VL_UNIT_ITEM", FieldDecimal, true, 20), field(6, "VL_ITEM", FieldDecimal, true, 20), field(7, "IND_PROP", FieldString, true, 1), field(8, "COD_PART", FieldString, false, 60), field(9, "TXT_COMPL", FieldString, false, 255), field(10, "COD_CTA", FieldString, false, 60)}}
}
