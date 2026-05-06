package read

import "testing"

func TestDecodeSpedLineLatin1(t *testing.T) {
	got := decodeSpedLine([]byte{'|', '0', '1', '5', '0', '|', 'B', 'A', 'I', 'R', 'R', 'O', ' ', 0xC7, 'A', '|'})
	want := "|0150|BAIRRO ÇA|"
	if got != want {
		t.Fatalf("decodeSpedLine() = %q, esperado %q", got, want)
	}
}

func TestBelongsToCNPJ(t *testing.T) {
	campos := []string{"", "0000", "COD", "0", "01012024", "31012024", "EMPRESA", "12.345.678/0001-90", ""}
	if !belongsToCNPJ(campos, "12345678000190") {
		t.Fatal("esperava aceitar CNPJ formatado do registro 0000")
	}
	if belongsToCNPJ(campos, "00000000000000") {
		t.Fatal("esperava rejeitar CNPJ diferente")
	}
}
