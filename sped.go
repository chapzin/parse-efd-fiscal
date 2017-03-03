package main

import (
	"bufio"
	"fmt"
	"./model/Bloco0"
	"os"
	"strings"
	"./model/BlocoC"
	"./model/BlocoH"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"time"
	"./limpaSped"

	"strconv"
)
var db, err = gorm.Open("mysql","root@/auditoria2?charset=utf8")

var count0000, count0190, count0200, countC100, countC170 int
var literalLines []string
var regC100 BlocoC.RegC100
var reg0000 Bloco0.Reg0000




func main() {
	// TODO -- Criar leitura de uma pasta todos arquivos txt e processar os speds
	file, err := os.Open("./sped.txt")
	checkErr(err)
	defer file.Close()
	defer db.Close()
	scanner := bufio.NewScanner(file)
	// guarda cada linha em indice diferente do slice
	for scanner.Scan() {
		literalLines = append(literalLines, scanner.Text())
	}

	// busca linha
	for _, line := range literalLines {
		line = strings.Replace(line,",",".",-1)
		ln := strings.Split(line, "|")
		// quando importado mesmo arquivo ele remove os dados
		if ln[1]== "0000" {

			limpaSped.LimpaSped("0000",ln[4],ln[5],ln[7])
			limpaSped.LimpaSped("0190",ln[4],ln[5],ln[7])
			limpaSped.LimpaSped("0200",ln[4],ln[5],ln[7])
			//CodFin :=
			fmt.Println(ln[3])

		}
		trataLinha(ln[1], line)

	}

	fmt.Println("Total de registro 0000: ",count0000)
	fmt.Println("Total de registro 0150: ",count0000)
	fmt.Println("Total de registro 0190: ",count0190)
	fmt.Println("Total de registro 0200: ",count0200)
	fmt.Println("Total de registro de entrada: ", countC100)
	fmt.Println("Total de registro de itens: ", countC170)

}

func convInt (string string) int {
	inteiro, err := strconv.Atoi(string)
	if err != nil {
		return 0
	}
	return inteiro
}

func convFloat (string string) float64 {
	float, err := strconv.ParseFloat(string,64)
	if err != nil {
		return 0
	}
	return float
}

func checkErr(err error) {
	if err != nil {
	logrus.Warn(err)
	//bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func dataSpedMysql(dtsped string) string  {
	if dtsped != "" {
		dia := dtsped[0:2]
		mes := dtsped[2:4]
		ano := dtsped[4:8]
		dtmysql := ano+ "-"+ mes + "-" + dia
		return dtmysql
	}
	return ""
}

func convertData(string string) time.Time {
	const longForm = "2006-01-02"
	DtIni, err := time.Parse(longForm, dataSpedMysql(string))
	if err !=nil {
		DtIni, _ := time.Parse(longForm, dataSpedMysql("1960-01-01"))
		return DtIni
	}
	return DtIni

}

func trataLinha(ln1 string, linha string) {
	switch ln1 {
	case "0000":
		ln := strings.Split(linha, "|")
		reg0000 = Bloco0.Reg0000{
			Reg:		ln[1],
			CodVer:		ln[2],
			CodFin:		convInt(ln[3]),
			DtIni:		convertData(ln[4]),
			DtFin:		convertData(ln[5]),
			Nome:		ln[6],
			Cnpj:		ln[7],
			Cpf:		ln[8],
			Uf:		ln[9],
			Ie:		ln[10],
			CodMun:		ln[11],
			Im:		ln[12],
			Suframa:	ln[13],
			IndPerfil:	ln[14],
			IndAtiv:	convInt(ln[15]),
		}
		db.NewRecord(reg0000)
		db.Create(&reg0000)
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
		ln := strings.Split(linha, "|")
		reg0150 := Bloco0.Reg0150{
			Reg : ln[1],
			CodPart : ln[2],
			Nome : ln[3],
			CodPais : ln[4],
			Cnpj : ln[5],
			Cpf : ln[6],
			Ie : ln[7],
			CodMun : ln[8],
			Suframa : ln[9],
			Endereco : ln[10],
			Num : ln[11],
			Compl : ln[12],
			Bairro : ln[13],
		}
		db.NewRecord(reg0150)
		db.Create(&reg0150)

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
		db.NewRecord(reg0190)
		db.Create(&reg0190)
		count0190++
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
			AliqIcms:   convFloat(ln[12]),
			DtIni: reg0000.DtIni,
			DtFin: reg0000.DtFin,
			Cnpj: reg0000.Cnpj,
		}
		db.NewRecord(reg0200)
		db.Create(&reg0200)
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
			FatConv: convFloat(ln[3]),
			DtIni: reg0000.DtIni,
			DtFin: reg0000.DtFin,
			Cnpj: reg0000.Cnpj,
		}
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
			DtDoc : convertData(ln[10]),
			DtES : convertData(ln[11]),
			VlDoc : convFloat(ln[12]),
			IndPgto : ln[13],
			VlDesc : convFloat(ln[14]),
			VlAbatNt : convFloat(ln[15]),
			VlMerc : convFloat(ln[16]),
			IndFrt : ln[17],
			VlFrt : convFloat(ln[18]),
			VlSeg : convFloat(ln[19]),
			VlOutDa : convFloat(ln[20]),
			VlBcIcms : convFloat(ln[21]),
			VlIcms : convFloat(ln[22]),
			VlBcIcmsSt : convFloat(ln[23]),
			VlIcmsSt : convFloat(ln[24]),
			VlIpi : convFloat(ln[25]),
			VlPis : convFloat(ln[26]),
			VlCofins : convFloat(ln[27]),
			VlPisSt : convFloat(ln[28]),
			VlCofinsSt : convFloat(ln[29]),
			DtIni: reg0000.DtIni,
			DtFin: reg0000.DtFin,
			Cnpj: reg0000.Cnpj,
		}
		countC100++
		db.NewRecord(regC100)
		db.Create(&regC100)

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
		fmt.Println("Quantidade registros C170:",len(ln))
		regC170 := BlocoC.RegC170{
			Reg: ln[1],
			NumItem: ln[2],
			CodItem : ln[3],
			DescrCompl : ln[4],
			Qtd : convFloat(ln[5]),
			Unid : ln[6],
			VlItem : convFloat(ln[7]),
			VlDesc : convFloat(ln[8]),
			IndMov : ln[9],
			CstIcms : ln[10],
			Cfop : ln[11],
			CodNat : ln[12],
			VlBcIcms : convFloat(ln[13]),
			AliqIcms : convFloat(ln[14]),
			VlIcms : convFloat(ln[15]),
			VlBcIcmsSt : convFloat(ln[16]),
			AliqSt : convFloat(ln[17]),
			VlIcmsSt : convFloat(ln[18]),
			IndApur : ln[19],
			CstIpi : ln[20],
			CodEnq : ln[21],
			VlBcIpi : convFloat(ln[22]),
			AliqIpi : convFloat(ln[23]),
			VlIpi : convFloat(ln[24]),
			CstPis : ln[25],
			VlBcPis : convFloat(ln[26]),
			AliqPis01 : convFloat(ln[27]),
			QuantBcPis : convFloat(ln[28]),
			AliqPis02 : convFloat(ln[29]),
			VlPis : convFloat(ln[30]),
			CstCofins : ln[31],
			VlBcCofins : convFloat(ln[32]),
			AliqCofins01 : convFloat(ln[33]),
			QuantBcCofins : convFloat(ln[34]),
			AliqCofins02 : convFloat(ln[35]),
			VlCofins : convFloat(ln[36]),
			CodCta : ln[37],
			EntradaSaida: regC100.IndOper,
			NumDoc: regC100.NumDoc,
			DtIni: reg0000.DtIni,
			DtFin: reg0000.DtFin,
			Cnpj: reg0000.Cnpj,

		}
		db.NewRecord(regC170)
		db.Create(&regC170)
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
			Qtd : convFloat(ln[4]),
			VlUnit : convFloat(ln[5]),
			VlItem : convFloat(ln[6]),
			IndProp : ln[7],
			CodPart : ln[8],
			TxtCompl : ln[9],
			CodCta : ln[10],
			VlItemIr : convFloat(ln[11]),
			DtIni: reg0000.DtIni,
			DtFin: reg0000.DtFin,
			Cnpj: reg0000.Cnpj,
		}
		db.NewRecord(regH010)
		db.Create(regH010)
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
