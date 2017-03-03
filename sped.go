package main

import (
	"bufio"
	"fmt"
	"model/Bloco0"
	"os"
	"strings"
	"model/BlocoC"
	"model/BlocoH"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"Util"
	"github.com/jinzhu/gorm"
)
var db, err = gorm.Open("mysql","root@/auditoria2?charset=utf8")

var count0000, count0190, count0200, countC100, countC170 int
var literalLines []string
var reg0220slice []Bloco0.Reg0220
var regC100slice []BlocoC.RegC100
var regC170slice []BlocoC.RegC170
var regH010slice []BlocoH.RegH010
var regC100 BlocoC.RegC100
var reg0000 Bloco0.Reg0000

func main() {
	file, err := os.Open("./sped.txt")
	checkErr(err)
	defer file.Close()
	defer db.Close()
	scanner := bufio.NewScanner(file)
	// guarda cada linha em indice diferente do slice
	for scanner.Scan() {
		literalLines = append(literalLines, scanner.Text())
	}

// teste


	// busca linha
	for _, line := range literalLines {
		ln := strings.Split(line, "|")
		// quando importado mesmo arquivo ele remove os dados
		if ln[1]== "0000" {
			util.LimpaSped("0000",ln[4],ln[5],ln[7])
			util.LimpaSped("0190",ln[4],ln[5],ln[7])
			util.LimpaSped("0200",ln[4],ln[5],ln[7])
		}
		trataLinha(ln[1], line)

	}

	fmt.Println("Total de registro 0000: ",count0000)
	fmt.Println("Total de registro 0190: ",count0190)
	fmt.Println("Total de registro 0200: ",count0200)
	fmt.Println("Total de registro de entrada: ", countC100)
	fmt.Println("Total de registro de itens: ", countC170)

}

func checkErr(err error) {
	if err != nil {
	logrus.Warn(err)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func dataSpedMysql(dtsped string) string  {
	dia := dtsped[0:2]
	mes := dtsped[2:4]
	ano := dtsped[4:8]
	dtmysql := ano+ "-"+ mes + "-" + dia
	return dtmysql
}

func trataLinha(ln1 string, linha string) {
	switch ln1 {
	case "0000":
		ln := strings.Split(linha, "|")

		reg0000 = Bloco0.Reg0000{
			Reg:		ln[1],
			CodVer:		ln[2],
			CodFin:		ln[3],
			DtIni:		dataSpedMysql(ln[4]),
			DtFin:		dataSpedMysql(ln[5]),
			Nome:		ln[6],
			Cnpj:		ln[7],
			Cpf:		ln[8],
			Uf:		ln[9],
			Ie:		ln[10],
			CodMun:		ln[11],
			Im:		ln[12],
			Suframa:	ln[13],
			IndPerfil:	ln[14],
			IndAtiv:	ln[15],
		}
		//_, err = db.Exec("INSERT INTO reg_0000 (REG,COD_VER,COD_FIN,DT_INI,DT_FIN,NOME,CNPJ,CPF,UF,IE,COD_MUN,IM,SUFRAMA,IND_PERFIL,IND_ATIV) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",reg0000.Reg,reg0000.CodVer,reg0000.CodFin,reg0000.DtIni,reg0000.DtFin,reg0000.Nome,reg0000.Cnpj,reg0000.Cpf,reg0000.Uf,reg0000.Ie,reg0000.CodMun,reg0000.Im,reg0000.Suframa,reg0000.IndPerfil,reg0000.IndAtiv)
		checkErr(err)
		count0000++

	case "0001":
		fmt.Println(linha)
	case "0005":
		fmt.Println(linha)
	case "0015":
		fmt.Println(linha)
	case "0100":
		fmt.Println(linha)
	case "0150":
		fmt.Println(linha)
	case "0190":
		ln := strings.Split(linha, "|")

		reg0190 := Bloco0.Reg0190{
			Reg:	ln[1],
			Unid:	ln[2],
			Descr:	ln[3],
			DtIni:	reg0000.DtIni,
			DtFin:	reg0000.DtFin,
			Cnpj:	reg0000.Cnpj,
		}
		//_, err = db.Exec("INSERT INTO reg_0190 (REG,UNID,DESCR,DT_INI,DT_FIN,CNPJ) VALUES (?,?,?,?,?,?)", reg0190.Reg, reg0190.Unid, reg0190.Descr, reg0190.DtIni, reg0190.DtFin, reg0190.Cnpj)
		checkErr(err)
		count0190++

		// Listando em json
		// reg190Json, _ := json.Marshal(reg0190)

	case "0200":
		ln := strings.Split(linha, "|")
		reg0200 := Bloco0.Reg0200{
			Reg:        ln[1],
			CodItem:    ln[2],
			DescrItem:  ln[3],
			CodBarra:   ln[4],
			CodAntItem: ln[5],
			UnidInv:    ln[6],
			TipoItem:   ln[7],
			CodNcm:     ln[8],
			ExIpi:      ln[9],
			CodGen:     ln[10],
			CodLst:     ln[11],
			AliqIcms:   ln[12],
			DtIni: reg0000.DtIni,
			DtFin: reg0000.DtFin,
			Cnpj: reg0000.Cnpj,
		}
		//_, err := db.Exec("INSERT INTO reg_0200 (REG,COD_ITEM,DESCR_ITEM,COD_BARRA,COD_ANT_ITEM,UNID_INV,TIPO_ITEM,COD_NCM,EX_IPI,COD_GEN,COD_LST,ALIQ_ICMS,DT_INI,DT_FIN,CNPJ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",reg0200.Reg,reg0200.CodItem,reg0200.DescrItem,reg0200.CodBarra,reg0200.CodAntItem,reg0200.UnidInv,reg0200.TipoItem,reg0200.CodNcm,reg0200.ExIpi,reg0200.CodGen,reg0200.CodLst,reg0200.AliqIcms,reg0200.DtIni,reg0200.DtFin,reg0200.Cnpj)
		checkErr(err)
		count0200++
	case "0205":
		fmt.Println(linha)
	case "0206":
		fmt.Println(linha)
	case "0210":
		fmt.Println(linha)
	case "0220":
		ln := strings.Split(linha,"|")
		reg0220 := Bloco0.Reg0220{
			Reg: ln[1],
			UnidConv: ln[2],
			FatConv: ln[3],
		}
		reg0220slice = append(reg0220slice, reg0220)
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
		regC100 =BlocoC.RegC100{
			Reg : ln[1],
			IndOper : ln[2],
			IndEmit : ln[3],
			CodPart : ln[4],
			CodMod : ln[5],
			CodSit : ln[6],
			Ser : ln[7],
			NumDoc : ln[8],
			ChvNfe : ln[9],
			DtDoc : ln[10],
			DtES : ln[11],
			VlDoc : ln[12],
			IndPgto : ln[13],
			VlDesc : ln[14],
			VlAbatNt : ln[15],
			VlMerc : ln[16],
			IndFrt : ln[17],
			VlFrt : ln[18],
			VlSeg : ln[19],
			VlOutDa : ln[20],
			VlBcIcms : ln[21],
			VlIcms : ln[22],
			VlBcIcmsSt : ln[23],
			VlIcmsSt : ln[24],
			VlIpi : ln[25],
			VlPis : ln[26],
			VlCofins : ln[27],
			VlPisSt : ln[28],
			VlCofinsSt : ln[29],
		}
		regC100slice = append(regC100slice, regC100)
		countC100++
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
		regC170 := BlocoC.RegC170{
			Reg: ln[1],
			NumItem: ln[2],
			CodItem : ln[3],
			DescrCompl : ln[4],
			Qtd : ln[5],
			Unid : ln[6],
			VlItem : ln[7],
			VlDesc : ln[8],
			IndMov : ln[9],
			CstIcms : ln[10],
			Cfop : ln[11],
			CodNat : ln[12],
			VlBcIcms : ln[13],
			AliqIcms : ln[14],
			VlIcms : ln[15],
			VlBcIcmsSt : ln[16],
			AliqSt : ln[17],
			VlIcmsSt : ln[18],
			IndApur : ln[19],
			CstIpi : ln[20],
			CodEnq : ln[21],
			VlBcIpi : ln[22],
			AliqIpi : ln[23],
			VlIpi : ln[24],
			CstPis : ln[25],
			VlBcPis : ln[26],
			AliqPis01 : ln[27],
			QuantBcPis : ln[28],
			AliqPis02 : ln[29],
			VlPis : ln[30],
			CstCofins : ln[31],
			VlBcCofins : ln[32],
			AliqCofins01 : ln[33],
			QuantBcCofins : ln[34],
			AliqCofins02 : ln[35],
			VlCofins : ln[36],
			CodCta : ln[37],
			EntradaSaida: regC100.IndOper,
			NumDoc: regC100.NumDoc,

		}
		regC170slice = append(regC170slice, regC170)
		fmt.Println(linha)
		countC170++
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
		fmt.Println(linha)
	case "C405":
		fmt.Println(linha)
	case "C410":
		fmt.Println(linha)
	case "C420":
		fmt.Println(linha)
	case "C425":
		fmt.Println(linha)
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
		fmt.Println(linha)
	case "H010":
		ln := strings.Split(linha, "|")
		regH010 := BlocoH.RegH010{
			Reg : ln[1],
			CodItem : ln[2],
			Unid : ln[3],
			Qtd : ln[4],
			VlUnit : ln[5],
			VlItem : ln[6],
			IndProp : ln[7],
			CodPart : ln[8],
			TxtCompl : ln[9],
			CodCta : ln[10],
			VlItemIr : ln[11],
		}
		regH010slice = append(regH010slice, regH010)
		fmt.Println(linha)
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
