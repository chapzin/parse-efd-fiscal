package exec

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/chapzin/parse-efd-fiscal/Models/NotaFiscal"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/clbanning/mxj"
	"github.com/jinzhu/gorm"
)

func InsertXmlNfeOld(db *gorm.DB, xmlPath string, digitosCodigo string) {
	digitosCodigo2 := tools.ConvInt(digitosCodigo)

	// Teste de lista produtos
	xmlFile, err := ioutil.ReadFile(xmlPath)
	reader := tools.ConvXml(xmlPath)
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
		// fmt.Printf("%#v\n",Item)
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
}

func InsertXmlNfe(db *gorm.DB, input NotaFiscal.NfeProc, digitosCodigo string) error {
	digitosCodigo2 := tools.ConvInt(digitosCodigo)

	header, err := headerNfe(input, digitosCodigo2)
	if err != nil {
		return err
	}

	errInsertHeader := db.Create(&header).Error
	if errInsertHeader != nil {
		if strings.Contains(errInsertHeader.Error(), "Error 1062") {
			return nil
		}

		return fmt.Errorf("erro ao inserir nota fiscal: %w", errInsertHeader)
	}

	return nil
}

func destinatarioNfe(input NotaFiscal.NfeProc) NotaFiscal.Destinatario {
	return NotaFiscal.Destinatario{
		CNPJ:    input.NFe.InfNFe.Dest.CNPJ,
		XNome:   input.NFe.InfNFe.Dest.XNome,
		XLgr:    input.NFe.InfNFe.Dest.EnderDest.XLgr,
		Nro:     input.NFe.InfNFe.Dest.EnderDest.Nro,
		XBairro: input.NFe.InfNFe.Dest.EnderDest.XBairro,
		CMun:    input.NFe.InfNFe.Dest.EnderDest.CMun,
		XMun:    input.NFe.InfNFe.Dest.EnderDest.XMun,
		Uf:      input.NFe.InfNFe.Dest.EnderDest.UF,
		Cep:     input.NFe.InfNFe.Dest.EnderDest.CEP,
		CPais:   input.NFe.InfNFe.Dest.EnderDest.CPais,
		XPais:   input.NFe.InfNFe.Dest.EnderDest.XPais,
		Fone:    input.NFe.InfNFe.Dest.EnderDest.Fone,
		Ie:      input.NFe.InfNFe.Dest.IE,
	}
}

func emitenteNfe(input NotaFiscal.NfeProc) NotaFiscal.Emitente {
	return NotaFiscal.Emitente{
		CNPJ:    input.NFe.InfNFe.Emit.CNPJ,
		XNome:   input.NFe.InfNFe.Emit.XNome,
		XLgr:    input.NFe.InfNFe.Emit.EnderEmit.XLgr,
		Nro:     input.NFe.InfNFe.Emit.EnderEmit.Nro,
		XCpl:    input.NFe.InfNFe.Emit.EnderEmit.XCpl,
		XBairro: input.NFe.InfNFe.Emit.EnderEmit.XBairro,
		CMun:    input.NFe.InfNFe.Emit.EnderEmit.CMun,
		XMun:    input.NFe.InfNFe.Emit.EnderEmit.XMun,
		Uf:      input.NFe.InfNFe.Emit.EnderEmit.UF,
		Cep:     input.NFe.InfNFe.Emit.EnderEmit.CEP,
		CPais:   input.NFe.InfNFe.Emit.EnderEmit.CPais,
		XPais:   input.NFe.InfNFe.Emit.EnderEmit.XPais,
		Fone:    input.NFe.InfNFe.Emit.EnderEmit.Fone,
		Ie:      input.NFe.InfNFe.Emit.IE,
	}
}

func itensNfe(input NotaFiscal.NfeProc, digitosCodigo int) ([]NotaFiscal.Item, error) {
	itens := make([]NotaFiscal.Item, 0, len(input.NFe.InfNFe.Det))

	for _, det := range input.NFe.InfNFe.Det {
		codigo := tools.AdicionaDigitosCodigo(det.Prod.CProd, digitosCodigo)

		qtd, err := stringToFloat(det.Prod.QCom)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter quantidade: %w", err)
		}

		vUnit, err := stringToFloat(det.Prod.VUnCom)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter valor unitário: %w", err)
		}

		item := NotaFiscal.Item{
			Codigo:    codigo,
			Ean:       det.Prod.CEAN,
			Descricao: det.Prod.XProd,
			Ncm:       det.Prod.NCM,
			Cfop:      det.Prod.CFOP,
			Unid:      det.Prod.UCom,
			Qtd:       qtd,
			VUnit:     vUnit,
			VTotal:    det.Prod.VProd,
			DtEmit:    tools.ConvertDataXml(input.NFe.InfNFe.Ide.DhEmi),
		}

		itens = append(itens, item)
	}

	return itens, nil
}

func headerNfe(input NotaFiscal.NfeProc, digitosCodigo int) (NotaFiscal.NotaFiscal, error) {
	emitente := emitenteNfe(input)
	destinatario := destinatarioNfe(input)
	itens, err := itensNfe(input, digitosCodigo)
	if err != nil {
		return NotaFiscal.NotaFiscal{}, err
	}

	notaFiscal := NotaFiscal.NotaFiscal{
		NNF:          input.NFe.InfNFe.Ide.NNF,
		ChNFe:        input.ProtNFe.InfProt.ChNFe,
		NatOp:        input.NFe.InfNFe.Ide.NatOp,
		IndPag:       input.NFe.InfNFe.Pag.DetPag.IndPag,
		Mod:          input.NFe.InfNFe.Ide.Mod,
		Serie:        input.NFe.InfNFe.Ide.Serie,
		DEmi:         tools.ConvertDataXml(input.NFe.InfNFe.Ide.DhEmi),
		TpNF:         input.NFe.InfNFe.Ide.TpNF,
		TpImp:        input.NFe.InfNFe.Ide.TpImp,
		TpEmis:       input.NFe.InfNFe.Ide.TpEmis,
		CDV:          input.NFe.InfNFe.Ide.CDV,
		TpAmb:        input.NFe.InfNFe.Ide.TpAmb,
		FinNFe:       input.NFe.InfNFe.Ide.FinNFe,
		ProcEmi:      input.NFe.InfNFe.Ide.ProcEmi,
		Emitente:     emitente,
		Destinatario: destinatario,
		Itens:        itens,
	}

	return notaFiscal, nil
}
