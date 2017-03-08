package SpedRead

import (
	"path/filepath"
	"os"
	"bufio"
	"strings"
	"github.com/jinzhu/gorm"
	"github.com/chapzin/parse-efd-fiscal/SpedExec"
	"github.com/chapzin/parse-efd-fiscal/SpedError"
	"time"
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
	"github.com/chapzin/parse-efd-fiscal/model/Bloco0"
)

var id int
var maxid = 100
// Ler todos os arquivos de uma determinada pasta
func RecursiveSpeds(path string) {
	filepath.Walk(path, func(file string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			ext := filepath.Ext(file)
			if ext == ".txt" {
				SpedError.CheckErr(err)
				// Possivelmente uma goroutines começando aqui
				r := SpedExec.Regs{}
				id++
				go InsertSped(file, &r)
				// Goroutines finalizando aqui
			}

			if ext == ".xml" {
				id++
				go InsertXml(file)
			}
			wait()
		}
		return nil
	})
}

func wait() {
	for {
		if id >= maxid {
			time.Sleep(1 * time.Second)
		} else {
			return
		}
	}

}

func InsertXml(xml string) {
	db, err := gorm.Open("mysql", "root@/auditoria2?charset=utf8")
	SpedError.CheckErr(err)
	reader := SpedConvert.ConvXml(xml)
	// Inserindo Reg0150 do xml
	reg0150Xml := Bloco0.Reg0150Xml{reader}
	reg0150 := Bloco0.CreateReg0150(reg0150Xml)
	db.NewRecord(reg0150)
	db.Create(&reg0150)
}

func InsertSped(sped string, r *SpedExec.Regs) {
	db, err := gorm.Open("mysql", "root@/auditoria2?charset=utf8")
	file, err := os.Open(sped)
	SpedError.CheckErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// guarda cada linha em indice diferente do slice
	for scanner.Scan() {
		ProcessRows(scanner.Text(), r, *db)
	}
	id--
}

func ProcessRows(line string, r *SpedExec.Regs, db gorm.DB) {
	if line == "" {
		return
	}
	if line[:1] == "|" {
		line = strings.Replace(line, ",", ".", -1)
		ln := strings.Split(line, "|")
		SpedExec.TrataLinha(ln[1], line, r, db)
	}
}
