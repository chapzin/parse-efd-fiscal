package tools

import (
	"github.com/clbanning/mxj"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const longForm = "2006-01-02"

// Funcao de conversao de arquivo string para inteiro caso venha vazio ele retorna zero
func ConvInt(string string) int {
	integer, err := strconv.Atoi(string)
	if err != nil {
		return 0
	}
	return integer
}

// Funcao que faz parse do xml e suas sub tags
func ConvXml(file string) func(pathTag string, tag string) string {
	xmlFile, err := ioutil.ReadFile(file)
	CheckErr(err)
	return func(pathTag string, tag string) string {
		nfe, errOpenXml := mxj.NewMapXml(xmlFile)
		CheckErr(errOpenXml)
		pathDest := nfe.PathsForKey(pathTag)
		if len(pathDest) > 0 {
			dest, err := nfe.ValuesForPath(pathDest[0])
			CheckErr(err)
			mv := mxj.Map(dest[0].(map[string]interface{}))
			if mv[tag] == nil {
				return ""
			}
			if tag == "dhEmi" {
				mv2 := strings.Split(mv[tag].(string), "T")
				return mv2[0]
			} else {
				return mv[tag].(string)
			}
		}
		return ""
	}
}

// Funcao converte string para float caso venha vazio ele retorna zero
func ConvFloat(string string) float64 {
	float, err := strconv.ParseFloat(string, 64)
	if err != nil {
		return 0
	}
	return float
}

func FloatToString(valor float64) string {
	return strconv.FormatFloat(valor, 'f', 2, 64)
}

func FloatToStringSped(valor float64) string {
	return strings.Replace(strconv.FormatFloat(valor, 'f', 2, 64), ".", ",", -1)
}

func AdicionaDigitosCodigo(codigo string, digitos int) string {
	if digitos != 0 {
		for len(codigo) < digitos {
			codigo = "0" + codigo

		}
	}
	return codigo
}

// Funcao tratando data recebida do arquivo do sped
func dataSpedMysql(dtsped string) string {
	if dtsped != "" {
		dia := dtsped[0:2]
		mes := dtsped[2:4]
		ano := dtsped[4:8]
		dtmysql := ano + "-" + mes + "-" + dia
		return dtmysql
	}
	return ""
}

// Funcao que trata data nula
func ConvertDataNull() time.Time {
	DtIni, _ := time.Parse(longForm, dataSpedMysql("01011960"))
	return DtIni
}

// funcao utilizada para data nula
func ConvertData(string string) time.Time {
	DtIni, err := time.Parse(longForm, dataSpedMysql(string))
	if err != nil {
		return ConvertDataNull()
	}
	return DtIni

}

// Funcao converte data do xml
func ConvertDataXml(string string) time.Time {
	DtIni, err := time.Parse(longForm, string)
	if err != nil {
		return ConvertDataNull()
	}
	return DtIni

}
