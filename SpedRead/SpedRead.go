package SpedRead

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/NotaFiscal"
	"github.com/chapzin/parse-efd-fiscal/SpedExec"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/clbanning/mxj"
	"github.com/jinzhu/gorm"
)

var id int
var maxid = 98

// Ler todos os arquivos de uma determinada pasta
func RecursiveSpeds(path string, dialect string, conexao string, digitosCodigo string) {
	filepath.Walk(path, func(file string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			ext := filepath.Ext(file)
			if ext == ".txt" || ext == ".TXT" {
				tools.CheckErr(err)
				r := SpedExec.Regs{}
				r.Digito = digitosCodigo
				id++
				go InsertSped(file, &r, dialect, conexao)

			}

			if ext == ".xml" || ext == ".XML" {
				id++
				go InsertXml(file, dialect, conexao, digitosCodigo)

			}
			wait()
		}
		return nil
	})
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

func InsertXml(xml string, dialect string, conexao string, digitosCodigo string) {
	digitosCodigo2 := tools.ConvInt(digitosCodigo)
	db, err := gorm.Open(dialect, conexao)
	// Teste de lista produtos
	xmlFile, err := ioutil.ReadFile(xml)
	reader := tools.ConvXml(xml)
	tools.CheckErr(err)
	nfe, errOpenXml := mxj.NewMapXml(xmlFile)
	tools.CheckErr(errOpenXml)
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
		i2 := i + 1
		codigoi := tools.AdicionaDigitosCodigo(codigo[i].(string), digitosCodigo2)
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
			Qtd:       tools.ConvFloat(qtdi),
			VUnit:     tools.ConvFloat(vuniti),
			VTotal:    tools.ConvFloat(vtotali),
			DtEmit:    tools.ConvertDataXml(dEmit),
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
		DEmi:         tools.ConvertDataXml(dEmit),
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

func InsertSped(sped string, r *SpedExec.Regs, dialect string, conexao string) {
	db, err := gorm.Open(dialect, conexao)
	file, err := os.Open(sped)
	tools.CheckErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// guarda cada linha em indice diferente do slice
	for scanner.Scan() {
		ProcessRows(scanner.Text(), r, db)
	}
	id--
}

func ProcessRows(line string, r *SpedExec.Regs, db *gorm.DB) {
	if line == "" {
		return
	}
	if line[:1] == "|" {
		line = strings.Replace(line, ",", ".", -1)
		ln := strings.Split(line, "|")
		SpedExec.TrataLinha(ln[1], line, r, db)
	}
}
