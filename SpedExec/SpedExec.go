package SpedExec

import (
	"fmt"
	"strings"
	"github.com/chapzin/parse-efd-fiscal/model/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/model/BlocoC"
	"github.com/chapzin/parse-efd-fiscal/model/BlocoH"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
	"github.com/chapzin/parse-efd-fiscal/SpedClean"
)

type Regs struct {
	RegC100 BlocoC.RegC100
	Reg0000 Bloco0.Reg0000
	Reg0200 Bloco0.Reg0200
}

func TrataLinha(ln1 string, linha string,r *Regs, db gorm.DB) {
	switch ln1 {
	case "0000":
		ln := strings.Split(linha, "|")
		reg0000Sped := Bloco0.Reg0000Sped{ln}
		r.Reg0000 = Bloco0.CreateReg0000(reg0000Sped)
		// Caso já exista informacoes da movimentacao dos produtos referente ao sped que está sendo importado os dados são deletados
		SpedClean.CleanSpedItems(r.Reg0000.Cnpj,r.Reg0000.DtIni,r.Reg0000.DtFin,db)
		db.NewRecord(r.Reg0000)
		db.Create(&r.Reg0000)
	case "0001":
		fmt.Println(linha)
	case "0005":
		fmt.Println(linha)
	case "0015":
		fmt.Println(linha)
	case "0100":
		fmt.Println(linha)
	case "0150":
		ln := strings.Split(linha, "|")
		reg0150sped := Bloco0.Reg0150Sped{ln,r.Reg0000}
		reg0150 := Bloco0.CreateReg0150(reg0150sped)
		db.NewRecord(reg0150)
		db.Create(&reg0150)
	case "0190":
		ln := strings.Split(linha, "|")
		reg0190sped := Bloco0.Reg0190Sped{ln,r.Reg0000}
		reg0190 := Bloco0.CreateReg0190(reg0190sped)
		db.NewRecord(reg0190)
		db.Create(&reg0190)
	case "0200":
		ln := strings.Split(linha, "|")
		reg0200Sped := Bloco0.Reg0200Sped{ln,r.Reg0000}
		r.Reg0200 = Bloco0.CreateReg0200(reg0200Sped)
		db.NewRecord(r.Reg0200)
		db.Create(&r.Reg0200)
	case "0205":
		fmt.Println(linha)
	case "0206":
		fmt.Println(linha)
	case "0210":
		fmt.Println(linha)
	case "0220":
		ln := strings.Split(linha,"|")
		reg0220sped := Bloco0.Reg0220Sped{ln,r.Reg0000,r.Reg0200}
		reg0220 := Bloco0.CreateReg0220(reg0220sped)
		db.NewRecord(reg0220)
		db.Create(&reg0220)
	case "0300":
		fmt.Println(linha)
	case "0305":
		fmt.Println(linha)
	case "0400":
		fmt.Println(linha)
	case "0450":
		fmt.Println(linha)
	case "0460":
		fmt.Println(linha)
	case "0500":
		fmt.Println(linha)
	case "0600":
		fmt.Println(linha)
	case "0990":
		fmt.Println(linha)
	case "C001":
		fmt.Println(linha)
	case "C100":
		ln := strings.Split(linha,"|")
		regC100sped := BlocoC.RegC100Sped{ln,r.Reg0000}
		r.RegC100 = BlocoC.CreateRegC100(regC100sped)
		db.NewRecord(r.RegC100)
		db.Create(&r.RegC100)
	case "C101":
		fmt.Println(linha)
	case "C105":
		fmt.Println(linha)
	case "C110":
		fmt.Println(linha)
	case "C111":
		fmt.Println(linha)
	case "C112":
		fmt.Println(linha)
	case "C113":
		fmt.Println(linha)
	case "C114":
		fmt.Println(linha)
	case "C115":
		fmt.Println(linha)
	case "C116":
		fmt.Println(linha)
	case "C120":
		fmt.Println(linha)
	case "C130":
		fmt.Println(linha)
	case "C140":
		fmt.Println(linha)
	case "C141":
		fmt.Println(linha)
	case "C160":
		fmt.Println(linha)
	case "C165":
		fmt.Println(linha)
	case "C170":
		ln := strings.Split(linha, "|")
		regC100sped := BlocoC.RegC170Sped{ln,r.Reg0000,r.RegC100}
		regC170 :=BlocoC.CreateRegC170(regC100sped)
		db.NewRecord(regC170)
		db.Create(&regC170)
	case "C171":
		fmt.Println(linha)
	case "C172":
		fmt.Println(linha)
	case "C173":
		fmt.Println(linha)
	case "C174":
		fmt.Println(linha)
	case "C175":
		fmt.Println(linha)
	case "C176":
		fmt.Println(linha)
	case "C177":
		fmt.Println(linha)
	case "C178":
		fmt.Println(linha)
	case "C179":
		fmt.Println(linha)
	case "C190":
		fmt.Println(linha)
	case "C195":
		fmt.Println(linha)
	case "C197":
		fmt.Println(linha)
	case "C300":
		fmt.Println(linha)
	case "C310":
		fmt.Println(linha)
	case "C320":
		fmt.Println(linha)
	case "C321":
		fmt.Println(linha)
	case "C350":
		fmt.Println(linha)
	case "C370":
		fmt.Println(linha)
	case "C390":
		fmt.Println(linha)
	case "C400":
		ln := strings.Split(linha, "|")
		regC400sped := BlocoC.RegC400Sped{ln,r.Reg0000}
		regC400 := BlocoC.CreateRegC400(regC400sped)
		db.NewRecord(regC400)
		db.Create(&regC400)
	case "C405":
		ln := strings.Split(linha, "|")
		regC405sped :=BlocoC.RegC405Sped{ln,r.Reg0000}
		regC405 := BlocoC.CreateRegC405(regC405sped)
		db.NewRecord(regC405)
		db.Create(&regC405)
	case "C410":
		fmt.Println(linha)
	case "C420":
		ln := strings.Split(linha, "|")
		regC420sped := BlocoC.RegC420Sped{ln,r.Reg0000}
		regC420 := BlocoC.CreateRegC420(regC420sped)
		db.NewRecord(regC420)
		db.Create(&regC420)
	case "C425":
		ln := strings.Split(linha, "|")
		regC425sped := BlocoC.RegC425Sped{ln,r.Reg0000}
		regC425 := BlocoC.CreateRegC425(regC425sped)
		db.NewRecord(regC425)
		db.Create(&regC425)
	case "C460":
		fmt.Println(linha)
	case "C465":
		fmt.Println(linha)
	case "C470":
		fmt.Println(linha)
	case "C490":
		fmt.Println(linha)
	case "C495":
		fmt.Println(linha)
	case "C500":
		fmt.Println(linha)
	case "C510":
		fmt.Println(linha)
	case "C590":
		fmt.Println(linha)
	case "C600":
		fmt.Println(linha)
	case "C601":
		fmt.Println(linha)
	case "C610":
		fmt.Println(linha)
	case "C690":
		fmt.Println(linha)
	case "C700":
		fmt.Println(linha)
	case "C790":
		fmt.Println(linha)
	case "C791":
		fmt.Println(linha)
	case "C800":
		fmt.Println(linha)
	case "C850":
		fmt.Println(linha)
	case "C860":
		fmt.Println(linha)
	case "C890":
		fmt.Println(linha)
	case "C990":
		fmt.Println(linha)
	case "D001":
		fmt.Println(linha)
	case "D100":
		fmt.Println(linha)
	case "D101":
		fmt.Println(linha)
	case "D110":
		fmt.Println(linha)
	case "D120":
		fmt.Println(linha)
	case "D130":
		fmt.Println(linha)
	case "D140":
		fmt.Println(linha)
	case "D150":
		fmt.Println(linha)
	case "D160":
		fmt.Println(linha)
	case "D161":
		fmt.Println(linha)
	case "D162":
		fmt.Println(linha)
	case "D170":
		fmt.Println(linha)
	case "D180":
		fmt.Println(linha)
	case "D190":
		fmt.Println(linha)
	case "D195":
		fmt.Println(linha)
	case "D197":
		fmt.Println(linha)
	case "D300":
		fmt.Println(linha)
	case "D301":
		fmt.Println(linha)
	case "D310":
		fmt.Println(linha)
	case "D350":
		fmt.Println(linha)
	case "D355":
		fmt.Println(linha)
	case "D360":
		fmt.Println(linha)
	case "D365":
		fmt.Println(linha)
	case "D370":
		fmt.Println(linha)
	case "D390":
		fmt.Println(linha)
	case "D400":
		fmt.Println(linha)
	case "D410":
		fmt.Println(linha)
	case "D411":
		fmt.Println(linha)
	case "D420":
		fmt.Println(linha)
	case "D500":
		fmt.Println(linha)
	case "D510":
		fmt.Println(linha)
	case "D530":
		fmt.Println(linha)
	case "D590":
		fmt.Println(linha)
	case "D600":
		fmt.Println(linha)
	case "D610":
		fmt.Println(linha)
	case "D690":
		fmt.Println(linha)
	case "D695":
		fmt.Println(linha)
	case "D697":
		fmt.Println(linha)
	case "D990":
		fmt.Println(linha)
	case "E001":
		fmt.Println(linha)
	case "E100":
		fmt.Println(linha)
	case "E110":
		fmt.Println(linha)
	case "E111":
		fmt.Println(linha)
	case "E112":
		fmt.Println(linha)
	case "E113":
		fmt.Println(linha)
	case "E115":
		fmt.Println(linha)
	case "E116":
		fmt.Println(linha)
	case "E200":
		fmt.Println(linha)
	case "E210":
		fmt.Println(linha)
	case "E220":
		fmt.Println(linha)
	case "E230":
		fmt.Println(linha)
	case "E240":
		fmt.Println(linha)
	case "E250":
		fmt.Println(linha)
	case "E300":
		fmt.Println(linha)
	case "E310":
		fmt.Println(linha)
	case "E311":
		fmt.Println(linha)
	case "E312":
		fmt.Println(linha)
	case "E313":
		fmt.Println(linha)
	case "E316":
		fmt.Println(linha)
	case "E500":
		fmt.Println(linha)
	case "E510":
		fmt.Println(linha)
	case "E520":
		fmt.Println(linha)
	case "E530":
		fmt.Println(linha)
	case "E990":
		fmt.Println(linha)
	case "G001":
		fmt.Println(linha)
	case "G110":
		fmt.Println(linha)
	case "G125":
		fmt.Println(linha)
	case "G126":
		fmt.Println(linha)
	case "G130":
		fmt.Println(linha)
	case "G140":
		fmt.Println(linha)
	case "G990":
		fmt.Println(linha)
	case "H001":
		fmt.Println(linha)
	case "H005":
		ln := strings.Split(linha, "|")
		regH005 := BlocoH.RegH005{
			Reg: 		ln[1],
			DtInv: 		SpedConvert.ConvertData(ln[2]),
			VlInv:		SpedConvert.ConvFloat(ln[3]),
			MotInv:		ln[4],
			DtIni:		r.Reg0000.DtIni,
			DtFin:		r.Reg0000.DtFin,
			Cnpj:		r.Reg0000.Cnpj,
		}
		db.NewRecord(regH005)
		db.Create(&regH005)
	case "H010":
		ln := strings.Split(linha, "|")
		regH010 := BlocoH.RegH010{
			Reg: 		ln[1],
			CodItem: 	ln[2],
			Unid: 		ln[3],
			Qtd: 		SpedConvert.ConvFloat(ln[4]),
			VlUnit: 	SpedConvert.ConvFloat(ln[5]),
			VlItem: 	SpedConvert.ConvFloat(ln[6]),
			IndProp: 	ln[7],
			CodPart: 	ln[8],
			TxtCompl: 	ln[9],
			CodCta: 	ln[10],
			VlItemIr:	SpedConvert.ConvFloat(ln[11]),
			DtIni: 		r.Reg0000.DtIni,
			DtFin:		r.Reg0000.DtFin,
			Cnpj:		r.Reg0000.Cnpj,
		}
		db.NewRecord(regH010)
		db.Create(&regH010)

	case "H020":
		fmt.Println(linha)
	case "H990":
		fmt.Println(linha)
	case "K001":
		fmt.Println(linha)
	case "K100":
		fmt.Println(linha)
	case "K200":
		fmt.Println(linha)
	case "K210":
		fmt.Println(linha)
	case "K215":
		fmt.Println(linha)
	case "K220":
		fmt.Println(linha)
	case "K230":
		fmt.Println(linha)
	case "K235":
		fmt.Println(linha)
	case "K250":
		fmt.Println(linha)
	case "K255":
		fmt.Println(linha)
	case "K260":
		fmt.Println(linha)
	case "K265":
		fmt.Println(linha)
	case "K270":
		fmt.Println(linha)
	case "K275":
		fmt.Println(linha)
	case "K280":
		fmt.Println(linha)
	case "K990":
		fmt.Println(linha)
	case "1001":
		fmt.Println(linha)
	case "1010":
		fmt.Println(linha)
	case "1100":
		fmt.Println(linha)
	case "1105":
		fmt.Println(linha)
	case "1110":
		fmt.Println(linha)
	case "1200":
		fmt.Println(linha)
	case "1210":
		fmt.Println(linha)
	case "1300":
		fmt.Println(linha)
	case "1310":
		fmt.Println(linha)
	case "1320":
		fmt.Println(linha)
	case "1350":
		fmt.Println(linha)
	case "1360":
		fmt.Println(linha)
	case "1370":
		fmt.Println(linha)
	case "1390":
		fmt.Println(linha)
	case "1391":
		fmt.Println(linha)
	case "1400":
		fmt.Println(linha)
	case "1500":
		fmt.Println(linha)
	case "1510":
		fmt.Println(linha)
	case "1600":
		fmt.Println(linha)
	case "1700":
		fmt.Println(linha)
	case "1710":
		fmt.Println(linha)
	case "1800":
		fmt.Println(linha)
	case "1900":
		fmt.Println(linha)
	case "1910":
		fmt.Println(linha)
	case "1920":
		fmt.Println(linha)
	case "1921":
		fmt.Println(linha)
	case "1922":
		fmt.Println(linha)
	case "1923":
		fmt.Println(linha)
	case "1925":
		fmt.Println(linha)
	case "1926":
		fmt.Println(linha)
	case "1990":
		fmt.Println(linha)
	case "9001":
		fmt.Println(linha)
	case "9900":
		fmt.Println(linha)
	case "9990":
		fmt.Println(linha)
	case "9999":
		fmt.Println(linha)
	default:

	}
}

