package SpedExec

import (
	"fmt"
	"strings"

	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoC"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoH"
	"github.com/chapzin/parse-efd-fiscal/SpedDB"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Regs struct {
	RegC100 BlocoC.RegC100
	Reg0000 Bloco0.Reg0000
	Reg0200 Bloco0.Reg0200
	RegH005 BlocoH.RegH005
	Digito  string
}

func TrataLinha(ln1 string, linha string, r *Regs, db *gorm.DB) error {
	switch ln1 {
	case "0000":
		ln := strings.Split(linha, "|")
		reg0000Sped := Bloco0.Reg0000Sped{Ln: ln}
		reg0000, err := Bloco0.CreateReg0000(reg0000Sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro 0000: %v", err)
		}
		r.Reg0000 = reg0000

		// Caso já exista informacoes da movimentacao dos produtos referente ao sped que está sendo importado os dados são deletados
		SpedDB.CleanSpedItems(r.Reg0000.Cnpj, r.Reg0000.DtIni, r.Reg0000.DtFin, db)
		db.NewRecord(r.Reg0000)
		db.Create(&r.Reg0000)
	case "0001":
		//fmt.Println(linha)
	case "0005":
		//fmt.Println(linha)
	case "0015":
		//fmt.Println(linha)
	case "0100":
		//fmt.Println(linha)
	case "0150":
		ln := strings.Split(linha, "|")
		reg0150sped := Bloco0.Reg0150Sped{Ln: ln, Reg0000: r.Reg0000}
		reg0150, err := Bloco0.CreateReg0150(reg0150sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro 0150: %v", err)
		}
		db.NewRecord(reg0150)
		db.Create(&reg0150)
	case "0190":
		ln := strings.Split(linha, "|")
		reg0190sped := Bloco0.Reg0190Sped{Ln: ln, Reg0000: r.Reg0000}
		reg0190, err := Bloco0.CreateReg0190(reg0190sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro 0190: %v", err)
		}
		db.NewRecord(reg0190)
		db.Create(&reg0190)
	case "0200":
		ln := strings.Split(linha, "|")
		reg0200Sped := Bloco0.Reg0200Sped{Ln: ln, Reg0000: r.Reg0000, Digito: r.Digito}
		reg0200, err := Bloco0.CreateReg0200(reg0200Sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro 0200: %v", err)
		}
		r.Reg0200 = reg0200
		db.NewRecord(r.Reg0200)
		db.Create(&r.Reg0200)
	case "0205":
		//fmt.Println(linha)
	case "0206":
		//fmt.Println(linha)
	case "0210":
		//fmt.Println(linha)
	case "0220":
		ln := strings.Split(linha, "|")
		reg0220sped := Bloco0.Reg0220Sped{Ln: ln, Reg0000: r.Reg0000, Reg0200: r.Reg0200, Digito: r.Digito}
		reg0220, err := Bloco0.CreateReg0220(reg0220sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro 0220: %v", err)
		}
		db.NewRecord(reg0220)
		db.Create(&reg0220)
	case "0300":
		//fmt.Println(linha)
	case "0305":
		//fmt.Println(linha)
	case "0400":
		//fmt.Println(linha)
	case "0450":
		//fmt.Println(linha)
	case "0460":
		ln := strings.Split(linha, "|")
		regC460sped := BlocoC.RegC460Sped{Ln: ln, Reg0000: r.Reg0000}
		regC460, err := BlocoC.CreateRegC460(regC460sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C460: %v", err)
		}
		db.NewRecord(regC460)
		db.Create(&regC460)
	case "0500":
		//fmt.Println(linha)
	case "0600":
		//fmt.Println(linha)
	case "0990":
		//fmt.Println(linha)
	case "C001":
		//fmt.Println(linha)
	case "C100":
		ln := strings.Split(linha, "|")
		regC100sped := BlocoC.RegC100Sped{Ln: ln, Reg0000: r.Reg0000}
		regC100, err := BlocoC.CreateRegC100(regC100sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C100: %v", err)
		}
		r.RegC100 = regC100
		db.NewRecord(r.RegC100)
		db.Create(&r.RegC100)
	case "C101":
		//fmt.Println(linha)
	case "C105":
		//fmt.Println(linha)
	case "C110":
		//fmt.Println(linha)
	case "C111":
		//fmt.Println(linha)
	case "C112":
		//fmt.Println(linha)
	case "C113":
		//fmt.Println(linha)
	case "C114":
		//fmt.Println(linha)
	case "C115":
		//fmt.Println(linha)
	case "C116":
		//fmt.Println(linha)
	case "C120":
		//fmt.Println(linha)
	case "C130":
		//fmt.Println(linha)
	case "C140":
		//fmt.Println(linha)
	case "C141":
		//fmt.Println(linha)
	case "C160":
		//fmt.Println(linha)
	case "C165":
		//fmt.Println(linha)
	case "C170":
		ln := strings.Split(linha, "|")
		regC170sped := BlocoC.RegC170Sped{
			Ln:      ln,
			Reg0000: r.Reg0000,
			RegC100: r.RegC100,
			Digito:  r.Digito,
		}
		regC170, err := BlocoC.CreateRegC170(regC170sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C170: %v", err)
		}
		db.NewRecord(regC170)
		db.Create(&regC170)
	case "C171":
		//fmt.Println(linha)
	case "C172":
		//fmt.Println(linha)
	case "C173":
		//fmt.Println(linha)
	case "C174":
		//fmt.Println(linha)
	case "C175":
		//fmt.Println(linha)
	case "C176":
		//fmt.Println(linha)
	case "C177":
		//fmt.Println(linha)
	case "C178":
		//fmt.Println(linha)
	case "C179":
		//fmt.Println(linha)
	case "C190":
		//fmt.Println(linha)
	case "C195":
		//fmt.Println(linha)
	case "C197":
		//fmt.Println(linha)
	case "C300":
		//fmt.Println(linha)
	case "C310":
		//fmt.Println(linha)
	case "C320":
		//fmt.Println(linha)
	case "C321":
		//fmt.Println(linha)
	case "C350":
		//fmt.Println(linha)
	case "C370":
		//fmt.Println(linha)
	case "C390":
		//fmt.Println(linha)
	case "C400":
		ln := strings.Split(linha, "|")
		regC400sped := BlocoC.RegC400Sped{Ln: ln, Reg0000: r.Reg0000}
		regC400, err := BlocoC.CreateRegC400(regC400sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C400: %v", err)
		}
		db.NewRecord(regC400)
		db.Create(&regC400)
	case "C405":
		ln := strings.Split(linha, "|")
		regC405sped := BlocoC.RegC405Sped{Ln: ln, Reg0000: r.Reg0000}
		regC405, err := BlocoC.CreateRegC405(regC405sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C405: %v", err)
		}
		db.NewRecord(regC405)
		db.Create(&regC405)
	case "C410":
		//fmt.Println(linha)
	case "C420":
		ln := strings.Split(linha, "|")
		regC420sped := BlocoC.RegC420Sped{Ln: ln, Reg0000: r.Reg0000}
		regC420, err := BlocoC.CreateRegC420(regC420sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C420: %v", err)
		}
		db.NewRecord(regC420)
		db.Create(&regC420)
	case "C425":
		ln := strings.Split(linha, "|")
		regC425sped := BlocoC.RegC425Sped{
			Ln:      ln,
			Reg0000: r.Reg0000,
			Digito:  r.Digito,
		}
		regC425, err := BlocoC.CreateRegC425(regC425sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C425: %v", err)
		}
		db.NewRecord(regC425)
		db.Create(&regC425)
	case "C465":
		ln := strings.Split(linha, "|")
		regC465sped := BlocoC.RegC465Sped{Ln: ln, Reg0000: r.Reg0000}
		regC465, err := BlocoC.CreateRegC465(regC465sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C465: %v", err)
		}
		db.NewRecord(regC465)
		db.Create(&regC465)
	case "C470":
		ln := strings.Split(linha, "|")
		regC470sped := BlocoC.RegC470Sped{
			Ln:      ln,
			Reg0000: r.Reg0000,
			Digito:  r.Digito,
		}
		regC470, err := BlocoC.CreateRegC470(regC470sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C470: %v", err)
		}
		db.NewRecord(regC470)
		db.Create(&regC470)
	case "C490":
		ln := strings.Split(linha, "|")
		regC490sped := BlocoC.RegC490Sped{Ln: ln, Reg0000: r.Reg0000}
		regC490, err := BlocoC.CreateRegC490(regC490sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C490: %v", err)
		}
		db.NewRecord(regC490)
		db.Create(&regC490)
	case "C495":
		//fmt.Println(linha)
	case "C500":
		//fmt.Println(linha)
	case "C510":
		//fmt.Println(linha)
	case "C590":
		//fmt.Println(linha)
	case "C600":
		//fmt.Println(linha)
	case "C601":
		//fmt.Println(linha)
	case "C610":
		//fmt.Println(linha)
	case "C690":
		//fmt.Println(linha)
	case "C700":
		//fmt.Println(linha)
	case "C790":
		//fmt.Println(linha)
	case "C791":
		//fmt.Println(linha)
	case "C800":
		ln := strings.Split(linha, "|")
		regC800sped := BlocoC.RegC800Sped{Ln: ln, Reg0000: r.Reg0000}
		regC800, err := BlocoC.CreateRegC800(regC800sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C800: %v", err)
		}
		db.NewRecord(regC800)
		db.Create(&regC800)
	case "C850":
		//fmt.Println(linha)
	case "C860":
		ln := strings.Split(linha, "|")
		regC860sped := BlocoC.RegC860Sped{Ln: ln, Reg0000: r.Reg0000}
		regC860, err := BlocoC.CreateRegC860(regC860sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C860: %v", err)
		}
		db.NewRecord(regC860)
		db.Create(&regC860)
	case "C890":
		ln := strings.Split(linha, "|")
		regC890sped := BlocoC.RegC890Sped{Ln: ln, Reg0000: r.Reg0000}
		regC890, err := BlocoC.CreateRegC890(regC890sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C890: %v", err)
		}
		db.NewRecord(regC890)
		db.Create(&regC890)
	case "C990":
		//fmt.Println(linha)
	case "D001":
		//fmt.Println(linha)
	case "D100":
		//fmt.Println(linha)
	case "D101":
		//fmt.Println(linha)
	case "D110":
		//fmt.Println(linha)
	case "D120":
		//fmt.Println(linha)
	case "D130":
		//fmt.Println(linha)
	case "D140":
		//fmt.Println(linha)
	case "D150":
		//fmt.Println(linha)
	case "D160":
		//fmt.Println(linha)
	case "D161":
		//fmt.Println(linha)
	case "D162":
		//fmt.Println(linha)
	case "D170":
		//fmt.Println(linha)
	case "D180":
		//fmt.Println(linha)
	case "D190":
		//fmt.Println(linha)
	case "D195":
		//fmt.Println(linha)
	case "D197":
		//fmt.Println(linha)
	case "D300":
		//fmt.Println(linha)
	case "D301":
		//fmt.Println(linha)
	case "D310":
		//fmt.Println(linha)
	case "D350":
		//fmt.Println(linha)
	case "D355":
		//fmt.Println(linha)
	case "D360":
		//fmt.Println(linha)
	case "D365":
		//fmt.Println(linha)
	case "D370":
		//fmt.Println(linha)
	case "D390":
		//fmt.Println(linha)
	case "D400":
		//fmt.Println(linha)
	case "D410":
		//fmt.Println(linha)
	case "D411":
		//fmt.Println(linha)
	case "D420":
		//fmt.Println(linha)
	case "D500":
		//fmt.Println(linha)
	case "D510":
		//fmt.Println(linha)
	case "D530":
		//fmt.Println(linha)
	case "D590":
		//fmt.Println(linha)
	case "D600":
		//fmt.Println(linha)
	case "D610":
		//fmt.Println(linha)
	case "D690":
		//fmt.Println(linha)
	case "D695":
		//fmt.Println(linha)
	case "D697":
		//fmt.Println(linha)
	case "D990":
		//fmt.Println(linha)
	case "E001":
		//fmt.Println(linha)
	case "E100":
		//fmt.Println(linha)
	case "E110":
		//fmt.Println(linha)
	case "E111":
		//fmt.Println(linha)
	case "E112":
		//fmt.Println(linha)
	case "E113":
		//fmt.Println(linha)
	case "E115":
		//fmt.Println(linha)
	case "E116":
		//fmt.Println(linha)
	case "E200":
		//fmt.Println(linha)
	case "E210":
		//fmt.Println(linha)
	case "E220":
		//fmt.Println(linha)
	case "E230":
		//fmt.Println(linha)
	case "E240":
		//fmt.Println(linha)
	case "E250":
		//fmt.Println(linha)
	case "E300":
		//fmt.Println(linha)
	case "E310":
		//fmt.Println(linha)
	case "E311":
		//fmt.Println(linha)
	case "E312":
		//fmt.Println(linha)
	case "E313":
		//fmt.Println(linha)
	case "E316":
		//fmt.Println(linha)
	case "E500":
		//fmt.Println(linha)
	case "E510":
		//fmt.Println(linha)
	case "E520":
		//fmt.Println(linha)
	case "E530":
		//fmt.Println(linha)
	case "E990":
		//fmt.Println(linha)
	case "G001":
		//fmt.Println(linha)
	case "G110":
		//fmt.Println(linha)
	case "G125":
		//fmt.Println(linha)
	case "G126":
		//fmt.Println(linha)
	case "G130":
		//fmt.Println(linha)
	case "G140":
		//fmt.Println(linha)
	case "G990":
		//fmt.Println(linha)
	case "H001":
		//fmt.Println(linha)
	case "H005":
		ln := strings.Split(linha, "|")
		regH005Sped := BlocoH.RegH005Sped{Ln: ln, Reg0000: r.Reg0000}
		regH005 := BlocoH.CreateRegH005(regH005Sped)
		db.NewRecord(regH005)
		db.Create(&regH005)
	case "H010":
		ln := strings.Split(linha, "|")
		regH010Sped := BlocoH.RegH010Sped{Ln: ln, Reg0000: r.Reg0000, RegH005: r.RegH005, Digito: r.Digito}
		regH010 := BlocoH.CreateRegH010(regH010Sped)
		db.NewRecord(regH010)
		db.Create(&regH010)
	case "H020":
		//fmt.Println(linha)
	case "H990":
		//fmt.Println(linha)
	case "K001":
		//fmt.Println(linha)
	case "K100":
		//fmt.Println(linha)
	case "K200":
		//fmt.Println(linha)
	case "K210":
		//fmt.Println(linha)
	case "K215":
		//fmt.Println(linha)
	case "K220":
		//fmt.Println(linha)
	case "K230":
		//fmt.Println(linha)
	case "K235":
		//fmt.Println(linha)
	case "K250":
		//fmt.Println(linha)
	case "K255":
		//fmt.Println(linha)
	case "K260":
		//fmt.Println(linha)
	case "K265":
		//fmt.Println(linha)
	case "K270":
		//fmt.Println(linha)
	case "K275":
		//fmt.Println(linha)
	case "K280":
		//fmt.Println(linha)
	case "K990":
		//fmt.Println(linha)
	case "1001":
		//fmt.Println(linha)
	case "1010":
		//fmt.Println(linha)
	case "1100":
		//fmt.Println(linha)
	case "1105":
		//fmt.Println(linha)
	case "1110":
		//fmt.Println(linha)
	case "1200":
		//fmt.Println(linha)
	case "1210":
		//fmt.Println(linha)
	case "1300":
		//fmt.Println(linha)
	case "1310":
		//fmt.Println(linha)
	case "1320":
		//fmt.Println(linha)
	case "1350":
		//fmt.Println(linha)
	case "1360":
		//fmt.Println(linha)
	case "1370":
		//fmt.Println(linha)
	case "1390":
		//fmt.Println(linha)
	case "1391":
		//fmt.Println(linha)
	case "1400":
		//fmt.Println(linha)
	case "1500":
		//fmt.Println(linha)
	case "1510":
		//fmt.Println(linha)
	case "1600":
		//fmt.Println(linha)
	case "1700":
		//fmt.Println(linha)
	case "1710":
		//fmt.Println(linha)
	case "1800":
		//fmt.Println(linha)
	case "1900":
		//fmt.Println(linha)
	case "1910":
		//fmt.Println(linha)
	case "1920":
		//fmt.Println(linha)
	case "1921":
		//fmt.Println(linha)
	case "1922":
		//fmt.Println(linha)
	case "1923":
		//fmt.Println(linha)
	case "1925":
		//fmt.Println(linha)
	case "1926":
		//fmt.Println(linha)
	case "1990":
		//fmt.Println(linha)
	case "9001":
		//fmt.Println(linha)
	case "9900":
		//fmt.Println(linha)
	case "9990":
		//fmt.Println(linha)
	case "9999":
		//fmt.Println(linha)
	case "C870":
		ln := strings.Split(linha, "|")
		regC870sped := BlocoC.RegC870Sped{
			Ln:      ln,
			Reg0000: r.Reg0000,
			Digito:  r.Digito,
		}
		regC870, err := BlocoC.CreateRegC870(regC870sped)
		if err != nil {
			return fmt.Errorf("erro ao criar registro C870: %v", err)
		}
		db.NewRecord(regC870)
		db.Create(&regC870)
	default:

	}
	return nil
}
