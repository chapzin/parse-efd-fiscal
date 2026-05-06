package layout

import "testing"

func TestDefaultCatalogContainsEssentialRecords(t *testing.T) {
	catalog := MustDefaultCatalog()
	for _, code := range []string{"0000", "0150", "0190", "0200", "0220", "C100", "C170", "H005", "H010", "9001", "9900", "9990", "9999"} {
		if _, ok := catalog.Record(code); !ok {
			t.Fatalf("registro %s não encontrado no catálogo", code)
		}
	}
}

func TestRecordFieldByIndex(t *testing.T) {
	catalog := MustDefaultCatalog()
	rec, _ := catalog.Record("0000")
	field, ok := rec.FieldByIndex(7)
	if !ok {
		t.Fatal("campo 7 não encontrado")
	}
	if field.Name != "CNPJ" || field.Type != FieldCNPJ {
		t.Fatalf("campo inesperado: %+v", field)
	}
}
