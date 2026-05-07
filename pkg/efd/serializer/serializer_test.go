package serializer

import (
	"bytes"
	"strings"
	"testing"

	"github.com/chapzin/parse-efd-fiscal/pkg/efd/parser"
)

func TestSerializePreservesRecords(t *testing.T) {
	doc := &parser.Document{Records: []parser.Record{
		{Code: "0000", Fields: []string{"0000", "019", "0"}},
		{Code: "9999", Fields: []string{"9999", "2"}},
	}}
	var out bytes.Buffer
	if err := Serialize(&out, doc, Options{}); err != nil {
		t.Fatal(err)
	}
	want := "|0000|019|0|\n|9999|2|\n"
	if out.String() != want {
		t.Fatalf("Serialize() = %q, esperado %q", out.String(), want)
	}
}

func TestSerializeRebuildsBlock9(t *testing.T) {
	input := strings.Join([]string{
		"|0000|019|0|01012024|31012024|EMPRESA TESTE|00000000000000||GO|ISENTO|5208707|||A|0|",
		"|0190|UN|UNIDADE|",
		"|9999|999|",
	}, "\n") + "\n"
	doc, err := parser.Parse(strings.NewReader(input), parser.Options{})
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := Serialize(&out, doc, Options{RebuildBlock9: true}); err != nil {
		t.Fatal(err)
	}
	got := out.String()
	for _, line := range []string{
		"|9001|0|",
		"|9900|0000|1|",
		"|9900|0190|1|",
		"|9900|9001|1|",
		"|9900|9900|6|",
		"|9900|9990|1|",
		"|9900|9999|1|",
		"|9990|9|",
		"|9999|11|",
	} {
		if !strings.Contains(got, line+"\n") {
			t.Fatalf("linha %q não encontrada em:\n%s", line, got)
		}
	}
}

func TestSerializeCRLF(t *testing.T) {
	doc := &parser.Document{Records: []parser.Record{{Code: "9999", Fields: []string{"9999", "1"}}}}
	var out bytes.Buffer
	if err := Serialize(&out, doc, Options{LineEnding: LineEndingCRLF}); err != nil {
		t.Fatal(err)
	}
	if out.String() != "|9999|1|\r\n" {
		t.Fatalf("line ending inesperado: %q", out.String())
	}
}
