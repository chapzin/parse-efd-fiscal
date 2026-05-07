// Package efd fornece a API pública de alto nível para parse e validação EFD ICMS/IPI.
package efd

import (
	"context"
	"io"

	"github.com/chapzin/parse-efd-fiscal/pkg/efd/diagnostic"
	"github.com/chapzin/parse-efd-fiscal/pkg/efd/layout"
	"github.com/chapzin/parse-efd-fiscal/pkg/efd/parser"
	"github.com/chapzin/parse-efd-fiscal/pkg/efd/validator"
)

// ParseOptions configura a leitura de arquivos EFD.
type ParseOptions = parser.Options

// ValidateOptions configura validações estruturais e de catálogo.
type ValidateOptions = validator.Options

// Document representa um arquivo EFD parseado sem dependência de banco de dados.
type Document = parser.Document

// Record representa uma linha/registo EFD parseado.
type Record = parser.Record

// Report agrupa diagnósticos de validação.
type Report = diagnostic.Report

// Diagnostic representa erro, aviso ou informação auditável.
type Diagnostic = diagnostic.Diagnostic

// Catalog descreve registros e campos suportados.
type Catalog = layout.Catalog

// DefaultCatalog retorna o catálogo mínimo padrão da biblioteca.
func DefaultCatalog() (Catalog, error) {
	return layout.DefaultCatalog()
}

// Parse lê um arquivo EFD de forma streaming e retorna um documento em memória.
func Parse(r io.Reader, opts ParseOptions) (*Document, error) {
	return parser.Parse(r, opts)
}

// Validate valida um documento parseado usando as opções informadas.
func Validate(doc *Document, opts ValidateOptions) Report {
	return validator.Validate(doc, opts)
}

// ParseAndValidate lê e valida um arquivo EFD em uma única chamada.
func ParseAndValidate(ctx context.Context, r io.Reader, parseOpts ParseOptions, validateOpts ValidateOptions) (*Document, Report, error) {
	select {
	case <-ctx.Done():
		return nil, Report{}, ctx.Err()
	default:
	}

	doc, err := Parse(r, parseOpts)
	if err != nil {
		return doc, Validate(doc, validateOpts), err
	}

	select {
	case <-ctx.Done():
		return doc, Report{}, ctx.Err()
	default:
	}

	report := Validate(doc, validateOpts)
	return doc, report, nil
}
