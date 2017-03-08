package SpedConvert

import (
	"strconv"
	"time"
	"io/ioutil"
	"github.com/chapzin/parse-efd-fiscal/SpedError"
	"github.com/clbanning/mxj"
)

const longForm = "2006-01-02"

func ConvInt(string string) int {
	integer, err := strconv.Atoi(string)
	if err != nil {
		return 0
	}
	return integer
}



func ConvXml(file string) func(pathTag string, tag string) string {
	xmlFile, err := ioutil.ReadFile(file)
	SpedError.CheckErr(err)
	return func (pathTag string, tag string) string {
		nfe, errOpenXml := mxj.NewMapXml(xmlFile)
		SpedError.CheckErr(errOpenXml)
		pathDest := nfe.PathsForKey(pathTag)
		dest, err := nfe.ValuesForPath(pathDest[0])
		SpedError.CheckErr(err)
		mv := mxj.Map(dest[0].(map[string]interface{}))
		return mv[tag].(string)
	}
}

func ConvFloat(string string) float64 {
	float, err := strconv.ParseFloat(string, 64)
	if err != nil {
		return 0
	}
	return float
}

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

func ConvertDataNull() time.Time {
	DtIni, _ := time.Parse(longForm, dataSpedMysql("1960-01-01"))
	return DtIni
}

func ConvertData(string string) time.Time {
	DtIni, err := time.Parse(longForm, dataSpedMysql(string))
	if err != nil {
		return ConvertDataNull()
	}
	return DtIni

}


