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
)

var id int
var maxid = 24
// Ler todos os arquivos de uma determinada pasta
func RecursiveSpeds(path string) {
	filepath.Walk(path, func(sped string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			ext := filepath.Ext(sped)
			if ext == ".txt" {
				// Possivelmente uma goroutines começando aqui
				r := SpedExec.Regs{}
				id++
				go ProcessaSped(sped, &r)
				wait()
				// Goroutines finalizando aqui
			}
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

func ProcessaSped(sped string, r *SpedExec.Regs) {

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
