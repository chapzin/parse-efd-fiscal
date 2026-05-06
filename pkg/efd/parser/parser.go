// Package parser fornece leitura streaming de arquivos EFD sem dependência de banco.
package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/chapzin/parse-efd-fiscal/pkg/efd/diagnostic"
)

type Encoding string

const (
	EncodingAuto   Encoding = "auto"
	EncodingUTF8   Encoding = "utf-8"
	EncodingLatin1 Encoding = "latin1"
)

type Options struct {
	Encoding     Encoding
	MaxLineBytes int
	Strict       bool
}

type Record struct {
	Line   int
	Code   string
	Fields []string
	Raw    string
}

type Document struct {
	Records     []Record
	Diagnostics []diagnostic.Diagnostic
}

func Parse(r io.Reader, opts Options) (*Document, error) {
	maxLineBytes := opts.MaxLineBytes
	if maxLineBytes <= 0 {
		maxLineBytes = 1024 * 1024
	}
	encoding := opts.Encoding
	if encoding == "" {
		encoding = EncodingAuto
	}

	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 64*1024), maxLineBytes)

	doc := &Document{}
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		raw := decodeLine(scanner.Bytes(), encoding)
		record, diagnostics := parseLine(lineNo, raw)
		doc.Diagnostics = append(doc.Diagnostics, diagnostics...)
		if record.Code != "" {
			doc.Records = append(doc.Records, record)
		}
	}
	if err := scanner.Err(); err != nil {
		d := diagnostic.Diagnostic{
			Severity: diagnostic.SeverityFatal,
			Code:     "EFD_PARSE_SCAN_ERROR",
			Message:  err.Error(),
		}
		doc.Diagnostics = append(doc.Diagnostics, d)
		if opts.Strict {
			return doc, err
		}
	}
	if opts.Strict {
		for _, d := range doc.Diagnostics {
			if d.IsError() {
				return doc, fmt.Errorf("documento EFD contém diagnósticos de erro")
			}
		}
	}
	return doc, nil
}

func parseLine(lineNo int, raw string) (Record, []diagnostic.Diagnostic) {
	trimmed := strings.TrimRight(raw, "\r\n")
	if strings.TrimSpace(trimmed) == "" {
		return Record{}, nil
	}
	fields := strings.Split(trimmed, "|")
	diagnostics := make([]diagnostic.Diagnostic, 0)
	if len(fields) < 3 || fields[0] != "" || fields[len(fields)-1] != "" {
		diagnostics = append(diagnostics, diagnostic.Diagnostic{
			Severity: diagnostic.SeverityError,
			Code:     "EFD_LINE_DELIMITERS",
			Message:  "linha EFD deve iniciar e terminar com delimitador |",
			Line:     lineNo,
			Value:    trimmed,
		})
		return Record{Line: lineNo, Raw: trimmed}, diagnostics
	}
	code := fields[1]
	if code == "" {
		diagnostics = append(diagnostics, diagnostic.Diagnostic{
			Severity: diagnostic.SeverityError,
			Code:     "EFD_EMPTY_RECORD_CODE",
			Message:  "código de registro vazio",
			Line:     lineNo,
		})
	}
	return Record{Line: lineNo, Code: code, Fields: fields[1 : len(fields)-1], Raw: trimmed}, diagnostics
}

func decodeLine(line []byte, encoding Encoding) string {
	line = bytes.TrimRight(line, "\r\n")
	switch encoding {
	case EncodingUTF8:
		return string(line)
	case EncodingLatin1:
		return latin1ToUTF8(line)
	case EncodingAuto, "":
		if utf8.Valid(line) {
			return string(line)
		}
		return latin1ToUTF8(line)
	default:
		if utf8.Valid(line) {
			return string(line)
		}
		return latin1ToUTF8(line)
	}
}

func latin1ToUTF8(line []byte) string {
	runes := make([]rune, 0, len(line))
	for _, b := range line {
		runes = append(runes, rune(b))
	}
	return string(runes)
}
