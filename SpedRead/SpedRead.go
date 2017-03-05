package SpedRead

import (
	"path/filepath"
	"os"
	"bufio"
	"strings"
	"github.com/jinzhu/gorm"
	"github.com/chapzin/parse-efd-fiscal/SpedExec"
	"github.com/chapzin/parse-efd-fiscal/SpedError"
)
// Ler todos os arquivos de uma determinada pasta
func RecursiveSpeds(path string,db gorm.DB){
	filepath.Walk(path, func(sped string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			ext := filepath.Ext(sped)
			if ext == ".txt"{
				// Possivelmente uma goroutines começando aqui
				r := SpedExec.Regs{}
				go ProcessaSped(sped,&r,db)
				// Goroutines finalizando aqui
			}
		}
		return nil
	})
}

func ProcessaSped (sped string,r *SpedExec.Regs,db2 gorm.DB) {

	db, err := gorm.Open("mysql","root@/auditoria2?charset=utf8")
	file, err := os.Open(sped)
	SpedError.CheckErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// guarda cada linha em indice diferente do slice
	for scanner.Scan() {
		ProcessRows(scanner.Text(),r,*db)
	}
}

func ProcessRows (line string,r *SpedExec.Regs ,db gorm.DB){
		if line == "" {
			return
		}
		if line[:1] == "|"{
			line = strings.Replace(line, ",", ".", -1)
			ln := strings.Split(line, "|")
			SpedExec.TrataLinha(ln[1], line,r, db)
		}
}



