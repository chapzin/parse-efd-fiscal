package parser

import (
	"strings"
	"testing"
)

func TestParseRecords(t *testing.T) {
	doc, err := Parse(strings.NewReader("|0000|019|0|01012024|31012024|EMPRESA|00000000000000|||||||||\n|9999|2|\n"), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if len(doc.Diagnostics) != 0 {
		t.Fatalf("diagnósticos inesperados: %+v", doc.Diagnostics)
	}
	if len(doc.Records) != 2 {
		t.Fatalf("registros = %d", len(doc.Records))
	}
	if doc.Records[0].Code != "0000" || doc.Records[1].Code != "9999" {
		t.Fatalf("códigos inesperados: %+v", doc.Records)
	}
}

func TestParseInvalidDelimiter(t *testing.T) {
	doc, err := Parse(strings.NewReader("0000|sem-delimitador\n"), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if len(doc.Diagnostics) != 1 || doc.Diagnostics[0].Code != "EFD_LINE_DELIMITERS" {
		t.Fatalf("diagnóstico inesperado: %+v", doc.Diagnostics)
	}
}

func TestParseLatin1Auto(t *testing.T) {
	doc, err := Parse(strings.NewReader(string([]byte{'|', '0', '1', '5', '0', '|', 'B', 'A', 'I', 'R', 'R', 'O', ' ', 0xC7, 'A', '|'})), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if got := doc.Records[0].Fields[1]; got != "BAIRRO ÇA" {
		t.Fatalf("campo latin1 = %q", got)
	}
}
