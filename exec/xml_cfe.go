package exec

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	cupomfiscal "github.com/chapzin/parse-efd-fiscal/Models/CupomFiscal"
	"github.com/jinzhu/gorm"
)

var ErrCSTICMSNotFound = errors.New("CST ICMS do item não encontrado")

func InsertXmlCfe(db *gorm.DB, input cupomfiscal.CFeXML) error {
	header, err := headerCfe(input)
	if err != nil {
		return err
	}

	errInsertHeader := db.Create(&header).Error
	if errInsertHeader != nil {
		if strings.Contains(errInsertHeader.Error(), "Error 1062") {
			return nil
		}

		return fmt.Errorf("erro ao inserir header: %w", errInsertHeader)
	}

	items, err := itensCfe(input)
	if err != nil {
		return err
	}

	for i := range items {
		items[i].IDHeader = header.ID
		errInsertItens := db.Create(&items[i]).Error
		if errInsertItens != nil {
			return fmt.Errorf("erro ao inserir itens: %w", errInsertItens)
		}
	}

	return nil
}

func headerCfe(input cupomfiscal.CFeXML) (cupomfiscal.CfeHeader, error) {
	output := new(cupomfiscal.CfeHeader)

	numero, err := strconv.Atoi(input.InfCFe.Ide.CNF)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter numero da nota fiscal: %w", err)
	}

	output.Numero = int64(numero)
	output.Chave = input.InfCFe.ID[3:]
	output.CUf = input.InfCFe.Ide.CUF
	output.Mod = input.InfCFe.Ide.Mod
	output.NserieSAT = input.InfCFe.Ide.NserieSAT
	output.NCFe = input.InfCFe.Ide.NCFe

	dtEmissao := input.InfCFe.Ide.DEmi + " " + input.InfCFe.Ide.HEmi

	output.DEmi, err = time.Parse("20060102 150405", dtEmissao)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter data de emissão: %w", err)
	}

	output.CDV = input.InfCFe.Ide.CDV
	output.TpAmb = input.InfCFe.Ide.TpAmb
	output.CNPJ = input.InfCFe.Emit.CNPJ
	output.NumeroCaixa = input.InfCFe.Ide.NumeroCaixa
	output.CnpjEmitente = input.InfCFe.Emit.CNPJ
	output.Emitente = input.InfCFe.Emit.XNome
	output.XFant = input.InfCFe.Emit.XFant
	// Precisa ser implementado
	// output.CnpjCpfDestinatario = input.InfCFe.Dest
	output.Destinatario = input.InfCFe.Dest.XNome

	output.VlICMS, err = stringToFloat(input.InfCFe.Total.ICMSTot.VICMS)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do ICMS: %w", err)
	}

	output.VlProd, err = stringToFloat(input.InfCFe.Total.ICMSTot.VProd)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor dos produtos: %w", err)
	}

	output.VlDesc, err = stringToFloat(input.InfCFe.Total.ICMSTot.VDesc)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do desconto: %w", err)
	}

	output.VlPIS, err = stringToFloat(input.InfCFe.Total.ICMSTot.VPIS)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do PIS: %w", err)
	}

	output.VlCOFINS, err = stringToFloat(input.InfCFe.Total.ICMSTot.VCOFINS)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do COFINS: %w", err)
	}

	output.VlPISST, err = stringToFloat(input.InfCFe.Total.ICMSTot.VPISST)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do PISST: %w", err)
	}

	output.VlCOFINSST, err = stringToFloat(input.InfCFe.Total.ICMSTot.VCOFINSST)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do COFINSST: %w", err)
	}

	output.VlOutro, err = stringToFloat(input.InfCFe.Total.ICMSTot.VOutro)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do Outro: %w", err)
	}

	output.VlCFe, err = stringToFloat(input.InfCFe.Total.VCFe)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do CFe: %w", err)
	}

	output.VlCFeLei12741, err = stringToFloat(input.InfCFe.Total.VCFeLei12741)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do CFe Lei 12741: %w", err)
	}

	output.CMP = input.InfCFe.Pgto.MP.CMP

	output.VMP, err = stringToFloat(input.InfCFe.Pgto.MP.VMP)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do MP: %w", err)
	}

	output.VTroco, err = stringToFloat(input.InfCFe.Pgto.VTroco)
	if err != nil {
		return *output, fmt.Errorf("erro ao converter valor do troco: %w", err)
	}

	return *output, nil
}

func itensCfe(input cupomfiscal.CFeXML) ([]cupomfiscal.CfeItem, error) {
	itens := make([]cupomfiscal.CfeItem, 0, len(input.InfCFe.Det))

	chave := input.InfCFe.ID[3:]

	for _, item := range input.InfCFe.Det {
		cfeItem := new(cupomfiscal.CfeItem)

		numeroItem, err := strconv.Atoi(item.NItem)
		if err != nil {
			return itens, fmt.Errorf("erro ao converter numero do item: %w", err)
		}

		cfeItem.NItem = numeroItem

		cfeItem.CProd = item.Prod.CProd
		cfeItem.CEAN = item.Prod.CEAN
		cfeItem.XProd = item.Prod.XProd
		cfeItem.NCM = item.Prod.NCM
		cfeItem.CFOP = item.Prod.CFOP
		cfeItem.UCom = item.Prod.UCom

		cfeItem.QCom, err = stringToFloat(item.Prod.QCom)
		if err != nil {
			return itens, fmt.Errorf("erro ao converter quantidade do item: %w", err)
		}

		cfeItem.VUnCom, err = stringToFloat(item.Prod.VUnCom)
		if err != nil {
			return itens, fmt.Errorf("erro ao converter valor unitário do item: %w", err)
		}

		cfeItem.VProd, err = stringToFloat(item.Prod.VProd)
		if err != nil {
			return itens, fmt.Errorf("erro ao converter valor do produto: %w", err)
		}

		cfeItem.IndRegra = item.Prod.IndRegra

		cfeItem.VItem, err = stringToFloat(item.Prod.VProd)
		if err != nil {
			return itens, fmt.Errorf("erro ao converter valor do item: %w", err)
		}

		cfeItem.VDesc, err = stringToFloat(item.Prod.VDesc)
		if err != nil {
			return itens, fmt.Errorf("erro ao converter valor do desconto: %w", err)
		}

		cfeItem.VOutro, err = stringToFloat(item.Prod.VOutro)
		if err != nil {
			return itens, fmt.Errorf("erro ao converter valor de outros: %w", err)
		}

		switch {
		case item.Imposto.ICMS.ICMSSN102 != nil:
			cfeItem.CSTICMS = item.Imposto.ICMS.ICMSSN102.Orig + item.Imposto.ICMS.ICMSSN102.CSOSN

		case item.Imposto.ICMS.ICMSSN500 != nil:
			cfeItem.CSTICMS = item.Imposto.ICMS.ICMSSN500.Orig + item.Imposto.ICMS.ICMSSN500.CSOSN

		case item.Imposto.ICMS.ICMSSN900 != nil:
			cfeItem.CSTICMS = item.Imposto.ICMS.ICMSSN900.Orig + item.Imposto.ICMS.ICMSSN900.CSOSN

		case item.Imposto.ICMS.ICMS40 != nil:
			cfeItem.CSTICMS = item.Imposto.ICMS.ICMS40.Orig + item.Imposto.ICMS.ICMS40.CST

			cfeItem.PICMS, err = stringToFloat(item.Imposto.ICMS.ICMS40.PIcms)
			if err != nil {
				return itens, fmt.Errorf("erro ao converter valor do ICMS: %w", err)
			}

			cfeItem.VICMS, err = stringToFloat(item.Imposto.ICMS.ICMS40.VIcms)
			if err != nil {
				return itens, fmt.Errorf("erro ao converter valor do ICMS: %w", err)
			}

		case item.Imposto.ICMS.ICMS00 != nil:
			cfeItem.CSTICMS = item.Imposto.ICMS.ICMS00.Orig + item.Imposto.ICMS.ICMS00.CST

			cfeItem.PICMS, err = stringToFloat(item.Imposto.ICMS.ICMS00.PIcms)
			if err != nil {
				return itens, fmt.Errorf("erro ao converter valor do ICMS: %w", err)
			}

			cfeItem.VICMS, err = stringToFloat(item.Imposto.ICMS.ICMS00.VIcms)
			if err != nil {
				return itens, fmt.Errorf("erro ao converter valor do ICMS: %w", err)
			}

		default:
			return itens, fmt.Errorf("%s- %w ", cfeItem.CProd, ErrCSTICMSNotFound)
		}

		cfeItem.UniqueKey = fmt.Sprintf("%s-%d", chave, cfeItem.NItem)

		itens = append(itens, *cfeItem)
	}

	return itens, nil
}

func stringToFloat(str string) (float64, error) {
	if str == "" {
		return 0, nil
	}

	if strings.Contains(str, ",") {
		str = strings.ReplaceAll(str, ",", ".")
	}

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("erro ao converter %s para float: %w", str, err)
	}

	return f, nil
}
