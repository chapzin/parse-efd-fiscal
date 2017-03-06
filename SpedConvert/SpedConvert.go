package SpedConvert

import (
	"strconv"
	"time"
)

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

func DataSpedMysql(dtsped string) string  {
	if dtsped != "" {
		dia := dtsped[0:2]
		mes := dtsped[2:4]
		ano := dtsped[4:8]
		dtmysql := ano+ "-"+ mes + "-" + dia
		return dtmysql
	}
	return ""
}

func ConvertData(string string) time.Time {
	const longForm = "2006-01-02"
	DtIni, err := time.Parse(longForm, DataSpedMysql(string))
	if err !=nil {
		DtIni, _ := time.Parse(longForm, DataSpedMysql("1960-01-01"))
		return DtIni
	}
	return DtIni

}

