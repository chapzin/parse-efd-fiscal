// Package types contém tipos fiscais fortes reutilizáveis pela biblioteca EFD.
package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var onlyDigitsRe = regexp.MustCompile(`\D`)

// CNPJ representa um CNPJ normalizado com 14 dígitos.
type CNPJ string

func ParseCNPJ(value string) (CNPJ, error) {
	digits := onlyDigits(value)
	if digits == "" {
		return "", nil
	}
	if len(digits) != 14 {
		return "", fmt.Errorf("CNPJ deve conter 14 dígitos: %q", value)
	}
	return CNPJ(digits), nil
}

func MustCNPJ(value string) CNPJ {
	cnpj, err := ParseCNPJ(value)
	if err != nil {
		panic(err)
	}
	return cnpj
}

func (c CNPJ) String() string { return string(c) }

func (c CNPJ) Formatted() string {
	value := string(c)
	if len(value) != 14 {
		return value
	}
	return fmt.Sprintf("%s.%s.%s/%s-%s", value[0:2], value[2:5], value[5:8], value[8:12], value[12:14])
}

func (c CNPJ) Valid() bool {
	return len(c) == 14 && allDigits(string(c))
}

// CPF representa um CPF normalizado com 11 dígitos.
type CPF string

func ParseCPF(value string) (CPF, error) {
	digits := onlyDigits(value)
	if digits == "" {
		return "", nil
	}
	if len(digits) != 11 {
		return "", fmt.Errorf("CPF deve conter 11 dígitos: %q", value)
	}
	return CPF(digits), nil
}

func (c CPF) String() string { return string(c) }

// UF representa uma unidade federativa brasileira.
type UF string

var validUF = map[UF]struct{}{
	"AC": {}, "AL": {}, "AP": {}, "AM": {}, "BA": {}, "CE": {}, "DF": {}, "ES": {}, "GO": {}, "MA": {},
	"MT": {}, "MS": {}, "MG": {}, "PA": {}, "PB": {}, "PR": {}, "PE": {}, "PI": {}, "RJ": {}, "RN": {},
	"RS": {}, "RO": {}, "RR": {}, "SC": {}, "SP": {}, "SE": {}, "TO": {}, "EX": {},
}

func ParseUF(value string) (UF, error) {
	uf := UF(strings.ToUpper(strings.TrimSpace(value)))
	if uf == "" {
		return "", nil
	}
	if _, ok := validUF[uf]; !ok {
		return "", fmt.Errorf("UF inválida: %q", value)
	}
	return uf, nil
}

func (u UF) String() string { return string(u) }

// CFOP representa um Código Fiscal de Operações e Prestações.
type CFOP string

func ParseCFOP(value string) (CFOP, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", nil
	}
	if len(value) != 4 || !allDigits(value) {
		return "", fmt.Errorf("CFOP deve conter 4 dígitos: %q", value)
	}
	return CFOP(value), nil
}

func (c CFOP) String() string { return string(c) }

func (c CFOP) Entrada() bool {
	return strings.HasPrefix(string(c), "1") || strings.HasPrefix(string(c), "2") || strings.HasPrefix(string(c), "3")
}

func (c CFOP) Saida() bool {
	return strings.HasPrefix(string(c), "5") || strings.HasPrefix(string(c), "6") || strings.HasPrefix(string(c), "7")
}

// NCM representa uma Nomenclatura Comum do Mercosul com 8 dígitos.
type NCM string

func ParseNCM(value string) (NCM, error) {
	value = onlyDigits(value)
	if value == "" {
		return "", nil
	}
	if len(value) != 8 {
		return "", fmt.Errorf("NCM deve conter 8 dígitos: %q", value)
	}
	return NCM(value), nil
}

func (n NCM) String() string { return string(n) }

// ChaveNFe representa chave de acesso NFe/NFCe/CFe de 44 dígitos.
type ChaveNFe string

func ParseChaveNFe(value string) (ChaveNFe, error) {
	digits := onlyDigits(value)
	if digits == "" {
		return "", nil
	}
	if len(digits) != 44 {
		return "", fmt.Errorf("chave NFe deve conter 44 dígitos: %q", value)
	}
	return ChaveNFe(digits), nil
}

func (c ChaveNFe) String() string { return string(c) }

// DataSped representa uma data no formato DDMMAAAA usado em arquivos SPED.
type DataSped struct {
	time.Time
}

func ParseDataSped(value string) (DataSped, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return DataSped{}, nil
	}
	if len(value) != 8 || !allDigits(value) {
		return DataSped{}, fmt.Errorf("data SPED deve estar em DDMMAAAA: %q", value)
	}
	t, err := time.Parse("02012006", value)
	if err != nil {
		return DataSped{}, fmt.Errorf("data SPED inválida %q: %w", value, err)
	}
	return DataSped{Time: t}, nil
}

func (d DataSped) String() string {
	if d.Time.IsZero() {
		return ""
	}
	return d.Time.Format("02012006")
}

func (d DataSped) ISO() string {
	if d.Time.IsZero() {
		return ""
	}
	return d.Time.Format("2006-01-02")
}

// Decimal é uma representação decimal exata simples para valores fiscais.
// Mantém coeficiente inteiro e escala decimal para evitar float64 em novos modelos.
type Decimal struct {
	coef  int64
	scale int32
}

func ParseDecimal(value string) (Decimal, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return Decimal{}, nil
	}
	value = strings.ReplaceAll(value, ",", ".")
	negative := strings.HasPrefix(value, "-")
	value = strings.TrimPrefix(value, "-")
	parts := strings.Split(value, ".")
	if len(parts) > 2 || parts[0] == "" || !allDigits(parts[0]) {
		return Decimal{}, fmt.Errorf("decimal inválido: %q", value)
	}
	scale := int32(0)
	digits := parts[0]
	if len(parts) == 2 {
		if !allDigits(parts[1]) {
			return Decimal{}, fmt.Errorf("decimal inválido: %q", value)
		}
		scale = int32(len(parts[1]))
		digits += parts[1]
	}
	coef, err := strconv.ParseInt(digits, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("decimal fora do limite: %w", err)
	}
	if negative {
		coef = -coef
	}
	return Decimal{coef: coef, scale: scale}, nil
}

func (d Decimal) String() string {
	negative := d.coef < 0
	coef := d.coef
	if negative {
		coef = -coef
	}
	digits := strconv.FormatInt(coef, 10)
	if d.scale == 0 {
		if negative {
			return "-" + digits
		}
		return digits
	}
	for len(digits) <= int(d.scale) {
		digits = "0" + digits
	}
	pos := len(digits) - int(d.scale)
	out := digits[:pos] + "." + digits[pos:]
	if negative {
		return "-" + out
	}
	return out
}

func (d Decimal) Coef() int64  { return d.coef }
func (d Decimal) Scale() int32 { return d.scale }

func onlyDigits(value string) string { return onlyDigitsRe.ReplaceAllString(value, "") }

func allDigits(value string) bool {
	if value == "" {
		return false
	}
	for _, r := range value {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
