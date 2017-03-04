package main

import (
	"bufio"
	"os"
	"strings"
	"./SpedError"
	"./SpedExec"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var literalLines []string

func main() {
	 db, err := gorm.Open("mysql","root@/auditoria2?charset=utf8")
	defer db.Close()
	SpedError.CheckErr(err)
	// TODO -- Criar leitura de uma pasta todos arquivos txt e processar os speds
	file, err := os.Open("./sped.txt")
	SpedError.CheckErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// guarda cada linha em indice diferente do slice
	for scanner.Scan() {
		literalLines = append(literalLines, scanner.Text())
	}

	// busca linha
	for _, line := range literalLines {
		line = strings.Replace(line,",",".",-1)
		ln := strings.Split(line, "|")
		SpedExec.TrataLinha(ln[1], line, *db)

	}
}