package SpedConvert

import (
	"strconv"
	"time"
)

const longForm = "2006-01-02"
func ConvInt (string string) int {
	integer, err := strconv.Atoi(string)
	if err != nil {
		return 0
	}
	return integer
}


func ConvFloat (string string) float64 {
	float, err := strconv.ParseFloat(string,64)
	if err != nil {
		return 0
	}
	return float
}

func dataSpedMysql(dtsped string) string  {
	if dtsped != "" {
		dia := dtsped[0:2]
		mes := dtsped[2:4]
		ano := dtsped[4:8]
		dtmysql := ano+ "-"+ mes + "-" + dia
		return dtmysql
	}
	return ""
}


func ConvertDataNull()time.Time{
	DtIni, _ := time.Parse(longForm, dataSpedMysql("1960-01-01"))
	return DtIni
}

func ConvertData(string string) time.Time {
	DtIni, err := time.Parse(longForm, dataSpedMysql(string))
	if err !=nil {
		return ConvertDataNull()
	}
	return DtIni

}

func DataXml(dest []interface{},tag string) string{
	var dest2 string
	for _,v := range dest {
		dests := v.(map[string]interface{})
		dest2 = dests[tag].(string)
	}
	return dest2
}

