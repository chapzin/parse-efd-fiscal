package main

import (
	"github.com/chapzin/parse-efd-fiscal/SpedError"
	_ "github.com/go-sql-driver/mysql"
//	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"

	"flag"
	"github.com/chapzin/parse-efd-fiscal/SpedDB"
	"github.com/chapzin/parse-efd-fiscal/SpedRead"
)



func main() {
	db, err := gorm.Open("mysql","root@/auditoria2?charset=utf8")
	//db, err := gorm.Open("postgres", "postgresql://chapzin@192.168.99.100:26257/auditoria?sslmode=disable")
	schema := flag.Bool("schema",false, "Recria as tabelas")
	flag.Parse()
	if *schema {
		// Recria o Schema do banco de dados
		SpedDB.Schema(*db)
	}
	defer db.Close()
	SpedError.CheckErr(err)

	// Lendo todos arquivos da pasta speds
	filesSpeds :=SpedRead.RecursiveSpeds("./speds")
	// Pega cada arquivo e ler linha a linha e envia para o banco de dados
	SpedRead.AddAllSpeds(filesSpeds,*db)


}

