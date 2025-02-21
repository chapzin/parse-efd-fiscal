package exec

import (
	"fmt"
	"strings"

	"github.com/chapzin/parse-efd-fiscal/Models/NotaFiscal"
	"github.com/chapzin/parse-efd-fiscal/tools"
	"github.com/jinzhu/gorm"
)

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
