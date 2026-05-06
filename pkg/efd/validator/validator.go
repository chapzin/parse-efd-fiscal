// Package validator valida documentos EFD com base em catálogo de layout.
package validator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chapzin/parse-efd-fiscal/pkg/efd/diagnostic"
	"github.com/chapzin/parse-efd-fiscal/pkg/efd/layout"
	"github.com/chapzin/parse-efd-fiscal/pkg/efd/parser"
	"github.com/chapzin/parse-efd-fiscal/pkg/efd/types"
)

type Options struct {
	Catalog        layout.Catalog
	AllowUnknown   bool
	ValidateBlock9 bool
}

func Validate(doc *parser.Document, opts Options) diagnostic.Report {
	catalog := opts.Catalog
	if catalog.Records == nil {
		catalog = layout.MustDefaultCatalog()
	}
	report := diagnostic.Report{Diagnostics: append([]diagnostic.Diagnostic{}, doc.Diagnostics...)}
	counts := map[string]int{}
	for _, record := range doc.Records {
		counts[record.Code]++
		spec, ok := catalog.Record(record.Code)
		if !ok {
			if !opts.AllowUnknown {
				report.Diagnostics = append(report.Diagnostics, diag(diagnostic.SeverityWarning, "EFD_UNKNOWN_RECORD", record, "", fmt.Sprintf("registro %s ainda não está no catálogo", record.Code)))
			}
			continue
		}
		report.Diagnostics = append(report.Diagnostics, validateRecord(record, spec)...)
	}
	if opts.ValidateBlock9 {
		report.Diagnostics = append(report.Diagnostics, validateBlock9(doc, counts)...)
	}
	return report
}

func validateRecord(record parser.Record, spec layout.RecordSpec) []diagnostic.Diagnostic {
	diagnostics := make([]diagnostic.Diagnostic, 0)
	fieldCount := len(record.Fields)
	if fieldCount < spec.MinFields || (spec.MaxFields > 0 && fieldCount > spec.MaxFields) {
		diagnostics = append(diagnostics, diag(diagnostic.SeverityError, "EFD_FIELD_COUNT", record, "", fmt.Sprintf("registro %s tem %d campos; esperado entre %d e %d", record.Code, fieldCount, spec.MinFields, spec.MaxFields)))
	}
	for _, field := range spec.Fields {
		if field.Index < 1 || field.Index > fieldCount {
			if field.Required {
				diagnostics = append(diagnostics, diag(diagnostic.SeverityError, "EFD_REQUIRED_FIELD_MISSING", record, field.Name, "campo obrigatório ausente"))
			}
			continue
		}
		value := record.Fields[field.Index-1]
		if field.Required && strings.TrimSpace(value) == "" {
			diagnostics = append(diagnostics, diag(diagnostic.SeverityError, "EFD_REQUIRED_FIELD_EMPTY", record, field.Name, "campo obrigatório vazio"))
			continue
		}
		if value == "" {
			continue
		}
		if field.Literal != "" && value != field.Literal {
			diagnostics = append(diagnostics, diag(diagnostic.SeverityError, "EFD_FIXED_FIELD", record, field.Name, fmt.Sprintf("valor fixo esperado %q, obtido %q", field.Literal, value)))
		}
		if field.MaxLen > 0 && len([]rune(value)) > field.MaxLen {
			diagnostics = append(diagnostics, diag(diagnostic.SeverityError, "EFD_FIELD_MAX_LENGTH", record, field.Name, fmt.Sprintf("campo excede tamanho máximo %d", field.MaxLen)))
		}
		if err := validateFieldType(field.Type, value); err != nil {
			diagnostics = append(diagnostics, diag(diagnostic.SeverityError, "EFD_FIELD_TYPE", record, field.Name, err.Error()))
		}
	}
	return diagnostics
}

func validateFieldType(fieldType layout.FieldType, value string) error {
	switch fieldType {
	case layout.FieldFixed, layout.FieldString:
		return nil
	case layout.FieldInteger:
		_, err := strconv.Atoi(value)
		return err
	case layout.FieldDecimal:
		_, err := types.ParseDecimal(value)
		return err
	case layout.FieldDate:
		_, err := types.ParseDataSped(value)
		return err
	case layout.FieldCNPJ:
		_, err := types.ParseCNPJ(value)
		return err
	case layout.FieldCPF:
		_, err := types.ParseCPF(value)
		return err
	case layout.FieldUF:
		_, err := types.ParseUF(value)
		return err
	case layout.FieldCFOP:
		_, err := types.ParseCFOP(value)
		return err
	case layout.FieldNCM:
		_, err := types.ParseNCM(value)
		return err
	case layout.FieldChNFe:
		_, err := types.ParseChaveNFe(value)
		return err
	default:
		return nil
	}
}

func validateBlock9(doc *parser.Document, counts map[string]int) []diagnostic.Diagnostic {
	diagnostics := make([]diagnostic.Diagnostic, 0)
	for _, record := range doc.Records {
		switch record.Code {
		case "9900":
			if len(record.Fields) < 3 {
				continue
			}
			code := record.Fields[1]
			reported, err := strconv.Atoi(record.Fields[2])
			if err != nil {
				diagnostics = append(diagnostics, diag(diagnostic.SeverityError, "EFD_9900_COUNT_TYPE", record, "QTD_REG_BLC", err.Error()))
				continue
			}
			actual := counts[code]
			if actual != reported {
				diagnostics = append(diagnostics, diag(diagnostic.SeverityError, "EFD_9900_COUNT_MISMATCH", record, "QTD_REG_BLC", fmt.Sprintf("registro 9900 informa %d ocorrência(s) de %s, mas foram encontradas %d", reported, code, actual)))
			}
		case "9999":
			if len(record.Fields) >= 2 {
				reported, err := strconv.Atoi(record.Fields[1])
				if err == nil && reported != len(doc.Records) {
					diagnostics = append(diagnostics, diag(diagnostic.SeverityError, "EFD_9999_TOTAL_MISMATCH", record, "QTD_LIN", fmt.Sprintf("9999 informa %d linhas, mas foram encontrados %d registros", reported, len(doc.Records))))
				}
			}
		}
	}
	return diagnostics
}

func diag(severity diagnostic.Severity, code string, record parser.Record, field string, message string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{Severity: severity, Code: code, Message: message, RecordCode: record.Code, Line: record.Line, Field: field}
}
