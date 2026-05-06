// Package diagnostic define diagnósticos estruturados para parsing e validação EFD.
package diagnostic

// Severity classifica o impacto de um diagnóstico.
type Severity string

const (
	SeverityInfo    Severity = "info"
	SeverityWarning Severity = "warning"
	SeverityError   Severity = "error"
	SeverityFatal   Severity = "fatal"
)

// Evidence aponta a origem de uma constatação fiscal ou estrutural.
type Evidence struct {
	Source string `json:"source"`
	Line   int    `json:"line,omitempty"`
	Field  string `json:"field,omitempty"`
	Value  string `json:"value,omitempty"`
}

// Diagnostic representa um erro, aviso ou informação auditável.
type Diagnostic struct {
	Severity   Severity   `json:"severity"`
	Code       string     `json:"code"`
	Message    string     `json:"message"`
	RecordCode string     `json:"record_code,omitempty"`
	Line       int        `json:"line,omitempty"`
	Field      string     `json:"field,omitempty"`
	Value      string     `json:"value,omitempty"`
	Evidence   []Evidence `json:"evidence,omitempty"`
	Suggestion string     `json:"suggestion,omitempty"`
}

func (d Diagnostic) IsError() bool {
	return d.Severity == SeverityError || d.Severity == SeverityFatal
}

// Report agrupa diagnósticos e fornece contadores úteis para CLI/API.
type Report struct {
	Diagnostics []Diagnostic `json:"diagnostics"`
}

func (r Report) HasErrors() bool {
	for _, d := range r.Diagnostics {
		if d.IsError() {
			return true
		}
	}
	return false
}

func (r Report) CountBySeverity(severity Severity) int {
	count := 0
	for _, d := range r.Diagnostics {
		if d.Severity == severity {
			count++
		}
	}
	return count
}
