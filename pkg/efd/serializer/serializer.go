// Package serializer escreve documentos EFD e pode reconstruir o Bloco 9.
package serializer

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/chapzin/parse-efd-fiscal/pkg/efd/parser"
)

type LineEnding string

const (
	LineEndingLF   LineEnding = "\n"
	LineEndingCRLF LineEnding = "\r\n"
)

type Options struct {
	LineEnding    LineEnding
	RebuildBlock9 bool
}

func Serialize(w io.Writer, doc *parser.Document, opts Options) error {
	lineEnding := string(opts.LineEnding)
	if lineEnding == "" {
		lineEnding = string(LineEndingLF)
	}
	records := doc.Records
	if opts.RebuildBlock9 {
		records = RebuildBlock9(records)
	}
	for _, record := range records {
		line := FormatRecord(record)
		if _, err := io.WriteString(w, line+lineEnding); err != nil {
			return err
		}
	}
	return nil
}

func FormatRecord(record parser.Record) string {
	if len(record.Fields) == 0 {
		return "||"
	}
	return "|" + strings.Join(record.Fields, "|") + "|"
}

func RebuildBlock9(records []parser.Record) []parser.Record {
	withoutBlock9 := make([]parser.Record, 0, len(records))
	for _, record := range records {
		if !isBlock9(record.Code) {
			withoutBlock9 = append(withoutBlock9, record)
		}
	}

	counts := map[string]int{}
	for _, record := range withoutBlock9 {
		counts[record.Code]++
	}

	block9 := buildBlock9(counts, len(withoutBlock9))
	out := make([]parser.Record, 0, len(withoutBlock9)+len(block9))
	out = append(out, withoutBlock9...)
	out = append(out, block9...)
	for i := range out {
		out[i].Line = i + 1
	}
	return out
}

func buildBlock9(counts map[string]int, nonBlock9Count int) []parser.Record {
	blockCounts := make(map[string]int, len(counts)+4)
	for code, count := range counts {
		blockCounts[code] = count
	}
	blockCounts["9001"] = 1
	blockCounts["9990"] = 1
	blockCounts["9999"] = 1
	// 9900 conta todos os registros informados no próprio 9900, incluindo 9900.
	blockCounts["9900"] = len(blockCounts) + 1

	codes := make([]string, 0, len(blockCounts))
	for code := range blockCounts {
		codes = append(codes, code)
	}
	sort.Strings(codes)

	records := make([]parser.Record, 0, len(codes)+3)
	records = append(records, newRecord("9001", "0"))
	for _, code := range codes {
		records = append(records, newRecord("9900", code, fmt.Sprintf("%d", blockCounts[code])))
	}
	qtdLin9 := len(records) + 2 // inclui 9990 e 9999.
	records = append(records, newRecord("9990", fmt.Sprintf("%d", qtdLin9)))
	records = append(records, newRecord("9999", fmt.Sprintf("%d", nonBlock9Count+qtdLin9)))
	return records
}

func newRecord(code string, values ...string) parser.Record {
	fields := make([]string, 0, len(values)+1)
	fields = append(fields, code)
	fields = append(fields, values...)
	return parser.Record{Code: code, Fields: fields, Raw: "|" + strings.Join(fields, "|") + "|"}
}

func isBlock9(code string) bool {
	return strings.HasPrefix(code, "9")
}
