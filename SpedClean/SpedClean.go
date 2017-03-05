package SpedClean

import (
	"time"
	"github.com/jinzhu/gorm"

)

func CleanSpedItems (cnpj string,dtIni time.Time,dtFin time.Time, db gorm.DB){
	db.Exec("DELETE FROM reg_0000 where cnpj =? and dt_ini =? and dt_fin =?",cnpj,dtIni,dtFin)
	db.Exec("DELETE FROM reg_0220 where cnpj =? and dt_ini =? and dt_fin =?",cnpj,dtIni,dtFin)
	db.Exec("DELETE FROM reg_c100 where cnpj =? and dt_ini =? and dt_fin =?",cnpj,dtIni,dtFin)
	db.Exec("DELETE FROM reg_c170 where cnpj =? and dt_ini =? and dt_fin =?",cnpj,dtIni,dtFin)
	db.Exec("DELETE FROM reg_c400 where cnpj =? and dt_ini =? and dt_fin =?",cnpj,dtIni,dtFin)
	db.Exec("DELETE FROM reg_c405 where cnpj =? and dt_ini =? and dt_fin =?",cnpj,dtIni,dtFin)
	db.Exec("DELETE FROM reg_c420 where cnpj =? and dt_ini =? and dt_fin =?",cnpj,dtIni,dtFin)
	db.Exec("DELETE FROM reg_c425 where cnpj =? and dt_ini =? and dt_fin =?",cnpj,dtIni,dtFin)
	db.Exec("DELETE FROM reg_h005 where cnpj =? and dt_ini =? and dt_fin =?",cnpj,dtIni,dtFin)
	db.Exec("DELETE FROM reg_h010 where cnpj =? and dt_ini =? and dt_fin =?",cnpj,dtIni,dtFin)
	//bufio.NewReader(os.Stdin).ReadBytes('\n')


}
