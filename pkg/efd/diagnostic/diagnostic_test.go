package diagnostic

import "testing"

func TestReportHasErrors(t *testing.T) {
	report := Report{Diagnostics: []Diagnostic{
		{Severity: SeverityWarning, Code: "WARN"},
		{Severity: SeverityError, Code: "ERR"},
	}}
	if !report.HasErrors() {
		t.Fatal("esperava erro no relatório")
	}
	if got := report.CountBySeverity(SeverityWarning); got != 1 {
		t.Fatalf("warnings = %d", got)
	}
}
