package SpedRead

import (
	"path/filepath"
	"os"
	"github.com/chapzin/parse-efd-fiscal/SpedError"
	"bufio"
	"strings"
	"github.com/chapzin/parse-efd-fiscal/SpedExec"
	"github.com/jinzhu/gorm"
)
// Ler todos os arquivos de uma determinada pasta
func RecursiveSpeds(path string) []string{
	fileList := []string{}
	filepath.Walk(path, func(sped string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			fileList = append(fileList, sped)
		}
		return nil
	})
	return fileList
}
// Pega todos os arquivos vindo de uma string e abre um a um, lendo linha por linha
func AddAllSpeds(files []string,db gorm.DB) {
	for _, sped := range files{
		ext :=string(sped[len(sped)-3:])
		if ext == "txt"{
			var literalLines []string
			file, err := os.Open(sped)
			SpedError.CheckErr(err)
			defer file.Close()
			scanner := bufio.NewScanner(file)
			// guarda cada linha em indice diferente do slice
			for scanner.Scan() {
				literalLines = append(literalLines, scanner.Text())
			}
			for _, line := range literalLines {
				if line[:1] == "|"{
					line = strings.Replace(line, ",", ".", -1)
					ln := strings.Split(line, "|")
					SpedExec.TrataLinha(ln[1], line, db)
				}
			}
		}
	}




}
