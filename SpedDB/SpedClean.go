package SpedDB

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Remove todos as tabelas dos registros
func CleanSpedItems(cnpj string, dtIni time.Time, dtFin time.Time, db *gorm.DB) {
	// Inicia todos os registros necessários
	regs := []string{"reg_0000", "reg_0150", "reg_0190", "reg_0200", "reg_0220", "reg_c100",
		"reg_c170", "reg_c400", "reg_c405",
		"reg_c420", "reg_c425", "reg_h005", "reg_h010"}

	for _, element := range regs {
		db.Exec("DELETE FROM "+element+" where cnpj =? and dt_ini =? and dt_fin =?", cnpj, dtIni, dtFin)
		if element == "reg_0150" {
			db.Exec("DELETE FROM "+element+" where cnpj_sped =? and dt_ini =? and dt_fin =?", cnpj, dtIni, dtFin)

		}
	}

}
