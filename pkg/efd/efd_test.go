package efd

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestParseAndValidate(t *testing.T) {
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

	doc, report, err := ParseAndValidate(context.Background(), strings.NewReader(input), ParseOptions{}, ValidateOptions{ValidateBlock9: true})
	if err != nil {
		t.Fatal(err)
	}
	if doc == nil || len(doc.Records) != 9 {
		t.Fatalf("documento inesperado: %+v", doc)
	}
	if report.HasErrors() {
		t.Fatalf("diagnósticos inesperados: %+v", report.Diagnostics)
	}
}

func TestParseAndValidateCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, _, err := ParseAndValidate(ctx, strings.NewReader(""), ParseOptions{}, ValidateOptions{})
	if err == nil {
		t.Fatal("esperava erro de contexto cancelado")
	}
}

func TestSerializePublicAPI(t *testing.T) {
	doc, err := Parse(strings.NewReader("|0000|019|0|\n"), ParseOptions{})
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := Serialize(&out, doc, SerializeOptions{}); err != nil {
		t.Fatal(err)
	}
	if out.String() != "|0000|019|0|\n" {
		t.Fatalf("Serialize API = %q", out.String())
	}
}
