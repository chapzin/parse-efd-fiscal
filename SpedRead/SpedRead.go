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
	"github.com/chapzin/parse-efd-fiscal/SpedConvert"
	"io/ioutil"
	"github.com/clbanning/mxj"
	"sync"

	"github.com/chapzin/parse-efd-fiscal/model/NotaFiscal"
)

var id int
var maxid = 100
// Ler todos os arquivos de uma determinada pasta
func RecursiveSpeds(path string, wg *sync.WaitGroup) {
	filepath.Walk(path, func(file string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			ext := filepath.Ext(file)
			if ext == ".txt" {
				SpedError.CheckErr(err)
				// Possivelmente uma goroutines começando aqui
				r := SpedExec.Regs{}
				id++
				go InsertSped(file, &r)
				// Goroutines finalizando aqui
			}

			if ext == ".xml" {
				id++
				go InsertXml(file)
			}
			wait()
		}
		return nil
	})
	wg.Done()
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

func InsertXml(xml string) {
	db, err := gorm.Open("mysql", "root@/auditoria2?charset=utf8")
	// Teste de lista produtos
	xmlFile, err := ioutil.ReadFile(xml)
	reader := SpedConvert.ConvXml(xml)
	SpedError.CheckErr(err)
	nfe, errOpenXml := mxj.NewMapXml(xmlFile)
	SpedError.CheckErr(errOpenXml)
	// Preenchendo o header da nfe
	nNf := reader("ide", "nNF")
	chnfe := reader("infProt", "chNFe")
	natOp := reader("ide", "natOp")
	indPag := reader("ide", "indPag")
	mod := reader("ide", "mod")
	serie := reader("ide", "serie")
	dEmit := reader("ide", "dEmi")
	if dEmit == "" {
		dhEmit := reader("ide", "dhEmi")
		dEmit = dhEmit
	}
	tpNf := reader("ide", "tpNF")
	tpImp := reader("ide", "tpImp")
	tpEmis := reader("ide", "tpEmis")
	cdv := reader("ide", "cDV")
	tpAmb := reader("ide", "tpAmb")
	finNFe := reader("ide", "finNFe")
	procEmi := reader("ide", "procEmi")

	// Preenchendo itens
	codigo, err := nfe.ValuesForKey("cProd")
	ean, err := nfe.ValuesForKey("cEAN")
	descricao, err := nfe.ValuesForKey("xProd")
	ncm, err := nfe.ValuesForKey("NCM")
	cfop, err := nfe.ValuesForKey("CFOP")
	unid, err := nfe.ValuesForKey("uCom")
	qtd, err := nfe.ValuesForKey("qCom")
	vUnit, err := nfe.ValuesForKey("vUnCom")
	vTotal, err := nfe.ValuesForKey("vProd")
	// Preenchendo Destinatario
	cnpj := reader("dest", "CNPJ")
	xNome := reader("dest", "xNome")
	xLgr := reader("enderDest", "xLgr")
	nro := reader("enderDest", "nro")
	xCpl := reader("enderDest", "xCpl")
	xBairro := reader("enderDest", "xBairro")
	cMun := reader("enderDest", "cMun")
	xMun := reader("enderDest", "xMun")
	uf := reader("enderDest", "UF")
	cep := reader("enderDest", "CEP")
	cPais := reader("enderDest", "cPais")
	xPais := reader("enderDest", "xPais")
	fone := reader("enderDest", "fone")
	ie := reader("dest", "IE")
	// Preenchendo Emitente
	cnpje := reader("emit", "CNPJ")
	xNomee := reader("emit", "xNome")
	xLgre := reader("enderEmit", "xLgr")
	nroe := reader("enderEmit", "nro")
	xCple := reader("enderEmit", "xCpl")
	xBairroe := reader("enderEmit", "xBairro")
	cMune := reader("enderEmit", "cMun")
	xMune := reader("enderEmit", "xMun")
	ufe := reader("enderEmit", "UF")
	cepe := reader("enderEmit", "CEP")
	cPaise := reader("enderEmit", "cPais")
	xPaise := reader("enderEmit", "xPais")
	fonee := reader("enderEmit", "fone")
	iee := reader("emit", "IE")

	destinatario := NotaFiscal.Destinatario{
		CNPJ:    cnpj,
		XNome:   xNome,
		XLgr:    xLgr,
		Nro:     nro,
		XCpl:    xCpl,
		XBairro: xBairro,
		CMun:    cMun,
		XMun:    xMun,
		Uf:      uf,
		Cep:     cep,
		CPais:   cPais,
		XPais:   xPais,
		Fone:    fone,
		Ie:      ie,
	}

	emitentede := NotaFiscal.Emitente{
		CNPJ:    cnpje,
		XNome:   xNomee,
		XLgr:    xLgre,
		Nro:     nroe,
		XCpl:    xCple,
		XBairro: xBairroe,
		CMun:    cMune,
		XMun:    xMune,
		Uf:      ufe,
		Cep:     cepe,
		CPais:   cPaise,
		XPais:   xPaise,
		Fone:    fonee,
		Ie:      iee,
	}

	var itens []NotaFiscal.Item

	for i, _ := range codigo {
		i2 := i+1
		codigoi := codigo[i].(string)
		eani := ean[i].(string)
		descricaoi := descricao[i].(string)
		ncmi := ncm[i].(string)
		cfopi := cfop[i].(string)
		unidi := unid[i].(string)
		qtdi := qtd[i].(string)
		vuniti := vUnit[i].(string)
		vtotali := vTotal[i2].(string)

		Item := NotaFiscal.Item{
			Codigo:    codigoi,
			Ean:       eani,
			Descricao: descricaoi,
			Ncm:       ncmi,
			Cfop:      cfopi,
			Unid:      unidi,
			Qtd:       SpedConvert.ConvFloat(qtdi),
			VUnit:     SpedConvert.ConvFloat(vuniti),
			VTotal:    SpedConvert.ConvFloat(vtotali),
			DtEmit:    SpedConvert.ConvertDataXml(dEmit),
		}
		itens = append(itens, Item)
		//fmt.Printf("%#v\n",Item)
	}

	notafiscal := NotaFiscal.NotaFiscal{
		NNF:          nNf,
		ChNFe:        chnfe,
		NatOp:        natOp,
		IndPag:       indPag,
		Mod:          mod,
		Serie:        serie,
		DEmi:         SpedConvert.ConvertDataXml(dEmit),
		TpNF:         tpNf,
		TpImp:        tpImp,
		TpEmis:       tpEmis,
		CDV:          cdv,
		TpAmb:        tpAmb,
		FinNFe:       finNFe,
		ProcEmi:      procEmi,
		Emitente:     emitentede,
		Destinatario: destinatario,
		Itens:        itens,
	}
	db.NewRecord(notafiscal)
	db.Create(&notafiscal)
	db.Close()
	id--
}

func InsertSped(sped string, r *SpedExec.Regs) {
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
