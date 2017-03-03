package limpaSped

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"bufio"
	"os"
	_ "github.com/go-sql-driver/mysql"

)

var db, err = sql.Open("mysql","root@/auditoria?charset=utf8")
func LimpaSped(bloco string,dtIni string,dtFin string, Cnpj string){
	bl := "reg_"+bloco
	_, errReg :=db.Query("Delete from "+bl+" where dt_ini=? and dt_fin=? and cnpj=?",dataSpedMysql(dtIni),dataSpedMysql(dtFin),Cnpj)
	checkErr(errReg)
}

func checkErr(err error) {
	if err != nil {
		logrus.Warn(err)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func dataSpedMysql(dtsped string) string  {
	dia := dtsped[0:2]
	mes := dtsped[2:4]
	ano := dtsped[4:8]
	dtmysql := ano+ "-"+ mes + "-" + dia
	return dtmysql
}
