package validator

import (
	"strings"
	"testing"

	"github.com/chapzin/parse-efd-fiscal/pkg/efd/parser"
)

func parseForTest(t *testing.T, input string) *parser.Document {
	t.Helper()
	doc, err := parser.Parse(strings.NewReader(input), parser.Options{})
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

func TestValidateValidMinimalDocument(t *testing.T) {
	input := strings.Join([]string{
		"|0000|019|0|01012024|31012024|EMPRESA TESTE|00000000000000||GO|ISENTO|5208707|||A|0|",
		"|9001|0|",
		"|9900|0000|1|",
		"|9900|9001|1|",
		"|9900|9900|5|",
		"|9900|9990|1|",
		"|9900|9999|1|",
		"|9990|6|",
		"|9999|9|",
	}, "\n") + "\n"
	doc := parseForTest(t, input)
	report := Validate(doc, Options{ValidateBlock9: true})
	if report.HasErrors() {
		t.Fatalf("diagnósticos inesperados: %+v", report.Diagnostics)
	}
}

func TestValidateRequiredField(t *testing.T) {
	doc := parseForTest(t, "|0190||UNIDADE|\n")
	report := Validate(doc, Options{})
	if !report.HasErrors() {
		t.Fatalf("esperava erro de campo obrigatório")
	}
	if report.Diagnostics[0].Code != "EFD_REQUIRED_FIELD_EMPTY" {
		t.Fatalf("código inesperado: %+v", report.Diagnostics)
	}
}

func TestValidateFieldType(t *testing.T) {
	doc := parseForTest(t, "|0000|019|X|01012024|31012024|EMPRESA TESTE|00000000000000||GO|ISENTO|5208707|||A|0|\n")
	report := Validate(doc, Options{})
	if !report.HasErrors() {
		t.Fatalf("esperava erro de tipo")
	}
	found := false
	for _, d := range report.Diagnostics {
		if d.Code == "EFD_FIELD_TYPE" && d.Field == "COD_FIN" {
			found = true
		}
	}
	if !found {
		t.Fatalf("erro de tipo COD_FIN não encontrado: %+v", report.Diagnostics)
	}
}

func TestValidateBlock9Mismatch(t *testing.T) {
	input := strings.Join([]string{
		"|0000|019|0|01012024|31012024|EMPRESA TESTE|00000000000000||GO|ISENTO|5208707|||A|0|",
		"|9900|0000|2|",
		"|9999|3|",
	}, "\n") + "\n"
	doc := parseForTest(t, input)
	report := Validate(doc, Options{ValidateBlock9: true})
	if !report.HasErrors() {
		t.Fatalf("esperava erro de contagem bloco 9")
	}
	found := false
	for _, d := range report.Diagnostics {
		if d.Code == "EFD_9900_COUNT_MISMATCH" {
			found = true
		}
	}
	if !found {
		t.Fatalf("erro de contagem não encontrado: %+v", report.Diagnostics)
	}
}
