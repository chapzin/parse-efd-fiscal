package types

import "testing"

func TestParseCNPJ(t *testing.T) {
	cnpj, err := ParseCNPJ("12.345.678/0001-90")
	if err != nil {
		t.Fatal(err)
	}
	if cnpj.String() != "12345678000190" {
		t.Fatalf("CNPJ normalizado = %q", cnpj.String())
	}
	if cnpj.Formatted() != "12.345.678/0001-90" {
		t.Fatalf("CNPJ formatado = %q", cnpj.Formatted())
	}
}

func TestParseCFOP(t *testing.T) {
	cfop, err := ParseCFOP("5102")
	if err != nil {
		t.Fatal(err)
	}
	if !cfop.Saida() || cfop.Entrada() {
		t.Fatalf("CFOP 5102 deveria ser saída")
	}
}

func TestParseDataSped(t *testing.T) {
	data, err := ParseDataSped("31012024")
	if err != nil {
		t.Fatal(err)
	}
	if data.ISO() != "2024-01-31" || data.String() != "31012024" {
		t.Fatalf("data inesperada: ISO=%s SPED=%s", data.ISO(), data.String())
	}
}

func TestParseDecimal(t *testing.T) {
	decimal, err := ParseDecimal("123,45")
	if err != nil {
		t.Fatal(err)
	}
	if decimal.String() != "123.45" || decimal.Coef() != 12345 || decimal.Scale() != 2 {
		t.Fatalf("decimal inesperado: %s coef=%d scale=%d", decimal.String(), decimal.Coef(), decimal.Scale())
	}
}
