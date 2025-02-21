package SpedRead

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models/NotaFiscal"
	"github.com/chapzin/parse-efd-fiscal/SpedExec"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/clbanning/mxj"
	"github.com/jinzhu/gorm"
)

var (
	id    int
	maxid = 98
	mu    sync.RWMutex // Mutex para proteger variáveis globais
)

// Estrutura para controlar o processamento concorrente
type processControl struct {
	wg       sync.WaitGroup
	errChan  chan error
	doneChan chan struct{}
	db       *gorm.DB
	digito   string
}

// newProcessControl cria uma nova instância de controle de processamento
func newProcessControl(db *gorm.DB, digito string) *processControl {
	return &processControl{
		errChan:  make(chan error, 100),
		doneChan: make(chan struct{}),
		db:       db,
		digito:   digito,
	}
}

// RecursiveSpeds processa recursivamente arquivos SPED em um diretório
func RecursiveSpeds(db *gorm.DB, path string, digito string) error {
	files, err := os.ReadDir(path) // Usando ReadDir ao invés de ioutil.ReadDir (depreciado)
	if err != nil {
		return fmt.Errorf("erro ao ler diretório: %v", err)
	}

	pc := newProcessControl(db, digito)
	defer close(pc.errChan)
	defer close(pc.doneChan)

	const maxWorkers = 4
	sem := make(chan struct{}, maxWorkers)
	defer close(sem)

	for _, f := range files {
		if !f.IsDir() && isTXT(f.Name()) {
			pc.wg.Add(1)
			sem <- struct{}{}
			go func(file os.DirEntry) {
				defer pc.wg.Done()
				defer func() { <-sem }()

				if err := processFile(pc, filepath.Join(path, file.Name())); err != nil {
					pc.errChan <- err
				}
			}(f)
		}
	}

	// Espera processamento terminar ou erro ocorrer
	done := make(chan struct{})
	go func() {
		pc.wg.Wait()
		close(done)
	}()

	select {
	case err := <-pc.errChan:
		return err
	case <-done:
		return nil
	}
}

// processFile processa um único arquivo SPED
func processFile(pc *processControl, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regs := &SpedExec.Regs{
		Digito: pc.digito,
		DB:     pc.db,
	}

	mu.Lock()
	id++
	currentID := id
	mu.Unlock()

	if currentID > maxid {
		return nil
	}

	for scanner.Scan() {
		linha := scanner.Text()
		campos := strings.Split(linha, "|")
		if len(campos) > 1 {
			if err := SpedExec.TrataLinha(campos[1], linha, regs, pc.db); err != nil {
				return fmt.Errorf("erro ao processar linha: %v", err)
			}
		}
	}

	return scanner.Err()
}

// RecursiveXmls processa recursivamente arquivos XML em um diretório
func RecursiveXmls(db *gorm.DB, path string, digito string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("erro ao ler diretório: %v", err)
	}

	pc := newProcessControl(db, digito)
	defer close(pc.errChan)
	defer close(pc.doneChan)

	const maxWorkers = 4
	sem := make(chan struct{}, maxWorkers)
	defer close(sem)

	for _, f := range files {
		if !f.IsDir() && isXML(f.Name()) {
			pc.wg.Add(1)
			sem <- struct{}{}
			go func(file os.DirEntry) {
				defer pc.wg.Done()
				defer func() { <-sem }()

				if err := processXMLFile(pc, filepath.Join(path, file.Name())); err != nil {
					pc.errChan <- err
				}
			}(f)
		}
	}

	done := make(chan struct{})
	go func() {
		pc.wg.Wait()
		close(done)
	}()

	select {
	case err := <-pc.errChan:
		return err
	case <-done:
		return nil
	}
}

// processXMLFile processa um único arquivo XML
func processXMLFile(pc *processControl, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("erro ao abrir XML: %v", err)
	}
	defer file.Close()

	// Processa o XML usando a função existente do projeto
	if err := processXML(pc.db, file, pc.digito); err != nil {
		return fmt.Errorf("erro ao processar XML: %v", err)
	}

	return nil
}

// processXML processa o conteúdo do arquivo XML
func processXML(db *gorm.DB, reader io.Reader, digito string) error {
	// Implemente aqui a lógica de processamento do XML
	// usando as estruturas existentes do projeto
	return nil
}

// Funções auxiliares
func isXML(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".xml"
}

func isTXT(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".txt"
}

// wait implementa um mecanismo de espera com timeout
func wait() error {
	timeout := time.After(5 * time.Minute)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if id < maxid {
				return nil
			}
		case <-timeout:
			return fmt.Errorf("timeout esperando processamento")
		}
	}
}

func InsertXml(xml string, dialect string, conexao string, digitosCodigo string) {
	digitosCodigo2 := tools.ConvInt(digitosCodigo)
	db, err := gorm.Open(dialect, conexao)
	tools.CheckErr(err)
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
	tools.CheckErr(err)
	ean, err := nfe.ValuesForKey("cEAN")
	tools.CheckErr(err)
	descricao, err := nfe.ValuesForKey("xProd")
	tools.CheckErr(err)
	ncm, err := nfe.ValuesForKey("NCM")
	tools.CheckErr(err)
	cfop, err := nfe.ValuesForKey("CFOP")
	tools.CheckErr(err)
	unid, err := nfe.ValuesForKey("uCom")
	tools.CheckErr(err)
	qtd, err := nfe.ValuesForKey("qCom")
	tools.CheckErr(err)
	vUnit, err := nfe.ValuesForKey("vUnCom")
	tools.CheckErr(err)
	vTotal, err := nfe.ValuesForKey("vProd")
	tools.CheckErr(err)
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

	for i := range codigo {
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
	tools.CheckErr(err)
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
