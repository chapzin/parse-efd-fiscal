package tools

import (
	"testing"
	"time"
)

func TestConvInt(t *testing.T) {
	v := ConvInt("5")
	if v != 5 {
		t.Error("Esperado um retorno inteiro ", v)
	}
}

func TestFloatToString(t *testing.T) {
	v := FloatToString(3.5)
	if v != "3.5" {
		t.Error("Esperado um retorno string ", v)
	}
}

func TestConvertDataNull(t *testing.T) {
	v := ConvertDataNull()
	correto, _ := time.Parse(longForm, dataSpedMysql("01011960"))
	if v != correto {
		t.Error("Nao foi retornado a data de 1960 como esperado")
	}

}

func TestConvertData(t *testing.T) {
	v := ConvertData("24121991")
	correto, _ := time.Parse(longForm, "1991-12-24")
	if v != correto {
		t.Error("Data nao convertida como esperado")
	}
}

func TestConvertDataXml(t *testing.T) {
	v := ConvertDataXml("2017-03-25")
	correto, _ := time.Parse(longForm, "2017-03-25")
	if v != correto {
		t.Error("Data de xml nao foi convertido corretamente")
	}
}

func TestConvFloat(t *testing.T) {
	v := ConvFloat("1.99")
	if v != 1.99 {
		t.Error("Conversao de string pra float nao foi feita")
	}
}
