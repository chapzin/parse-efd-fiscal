package NotaFiscal

import "encoding/xml"

type NfeProc struct {
	XMLName xml.Name `xml:"nfeProc"`
	Text    string   `xml:",chardata"`
	Versao  string   `xml:"versao,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	NFe     struct {
		Text   string `xml:",chardata"`
		Xmlns  string `xml:"xmlns,attr"`
		InfNFe struct {
			Text   string `xml:",chardata"`
			Versao string `xml:"versao,attr"`
			ID     string `xml:"Id,attr"`
			Ide    struct {
				Text     string `xml:",chardata"`
				CUF      string `xml:"cUF"`
				CNF      string `xml:"cNF"`
				NatOp    string `xml:"natOp"`
				Mod      string `xml:"mod"`
				Serie    string `xml:"serie"`
				NNF      string `xml:"nNF"`
				DhEmi    string `xml:"dhEmi"`
				DhSaiEnt string `xml:"dhSaiEnt"`
				TpNF     string `xml:"tpNF"`
				IdDest   string `xml:"idDest"`
				CMunFG   string `xml:"cMunFG"`
				TpImp    string `xml:"tpImp"`
				TpEmis   string `xml:"tpEmis"`
				CDV      string `xml:"cDV"`
				TpAmb    string `xml:"tpAmb"`
				FinNFe   string `xml:"finNFe"`
				IndFinal string `xml:"indFinal"`
				IndPres  string `xml:"indPres"`
				ProcEmi  string `xml:"procEmi"`
				VerProc  string `xml:"verProc"`
			} `xml:"ide"`
			Emit struct {
				Text      string `xml:",chardata"`
				CNPJ      string `xml:"CNPJ"`
				XNome     string `xml:"xNome"`
				XFant     string `xml:"xFant"`
				EnderEmit struct {
					Text    string `xml:",chardata"`
					XLgr    string `xml:"xLgr"`
					Nro     string `xml:"nro"`
					XCpl    string `xml:"xCpl"`
					XBairro string `xml:"xBairro"`
					CMun    string `xml:"cMun"`
					XMun    string `xml:"xMun"`
					UF      string `xml:"UF"`
					CEP     string `xml:"CEP"`
					CPais   string `xml:"cPais"`
					XPais   string `xml:"xPais"`
					Fone    string `xml:"fone"`
				} `xml:"enderEmit"`
				IE  string `xml:"IE"`
				CRT string `xml:"CRT"`
			} `xml:"emit"`
			Dest struct {
				Text      string `xml:",chardata"`
				CNPJ      string `xml:"CNPJ"`
				CPF       string `xml:"CPF"`
				XNome     string `xml:"xNome"`
				EnderDest struct {
					Text    string `xml:",chardata"`
					XLgr    string `xml:"xLgr"`
					Nro     string `xml:"nro"`
					XBairro string `xml:"xBairro"`
					CMun    string `xml:"cMun"`
					XMun    string `xml:"xMun"`
					UF      string `xml:"UF"`
					CEP     string `xml:"CEP"`
					CPais   string `xml:"cPais"`
					XPais   string `xml:"xPais"`
					Fone    string `xml:"fone"`
				} `xml:"enderDest"`
				IndIEDest string `xml:"indIEDest"`
				IE        string `xml:"IE"`
				Email     string `xml:"email"`
			} `xml:"dest"`
			Det []struct {
				Text  string `xml:",chardata"`
				NItem string `xml:"nItem,attr"`
				Prod  struct {
					Text      string  `xml:",chardata"`
					CProd     string  `xml:"cProd"`
					CEAN      string  `xml:"cEAN"`
					XProd     string  `xml:"xProd"`
					NCM       string  `xml:"NCM"`
					CEST      string  `xml:"CEST"`
					IndEscala string  `xml:"indEscala"`
					EXTIPI    string  `xml:"EXTIPI"`
					CFOP      string  `xml:"CFOP"`
					UCom      string  `xml:"uCom"`
					QCom      string  `xml:"qCom"`
					VUnCom    string  `xml:"vUnCom"`
					VProd     float64 `xml:"vProd"`
					CEANTrib  string  `xml:"cEANTrib"`
					UTrib     string  `xml:"uTrib"`
					QTrib     string  `xml:"qTrib"`
					VUnTrib   string  `xml:"vUnTrib"`
					VOutro    string  `xml:"vOutro"`
					VDesc     float64 `xml:"vDesc"`
					VFrete    float64 `xml:"vFrete"`
					IndTot    string  `xml:"indTot"`
					DI        struct {
						Text         string `xml:",chardata"`
						NDI          string `xml:"nDI"`
						DDI          string `xml:"dDI"`
						XLocDesemb   string `xml:"xLocDesemb"`
						UFDesemb     string `xml:"UFDesemb"`
						DDesemb      string `xml:"dDesemb"`
						TpViaTransp  string `xml:"tpViaTransp"`
						VAFRMM       string `xml:"vAFRMM"`
						TpIntermedio string `xml:"tpIntermedio"`
						CExportador  string `xml:"cExportador"`
						Adi          struct {
							Text        string `xml:",chardata"`
							NAdicao     string `xml:"nAdicao"`
							NSeqAdic    string `xml:"nSeqAdic"`
							CFabricante string `xml:"cFabricante"`
						} `xml:"adi"`
					} `xml:"DI"`
				} `xml:"prod"`
				Imposto struct {
					Text     string `xml:",chardata"`
					VTotTrib string `xml:"vTotTrib"`
					ICMS     struct {
						Text   string `xml:",chardata"`
						ICMS60 struct {
							Text            string `xml:",chardata"`
							Orig            string `xml:"orig"`
							CST             string `xml:"CST"`
							VBCSTRet        string `xml:"vBCSTRet"`
							PST             string `xml:"pST"`
							VICMSSubstituto string `xml:"vICMSSubstituto"`
							VICMSSTRet      string `xml:"vICMSSTRet"`
							PRedBCEfet      string `xml:"pRedBCEfet"`
							VBCEfet         string `xml:"vBCEfet"`
							PICMSEfet       string `xml:"pICMSEfet"`
							VICMSEfet       string `xml:"vICMSEfet"`
						} `xml:"ICMS60"`
						ICMS20 struct {
							Text   string `xml:",chardata"`
							Orig   string `xml:"orig"`
							CST    string `xml:"CST"`
							ModBC  string `xml:"modBC"`
							PRedBC string `xml:"pRedBC"`
							VBC    string `xml:"vBC"`
							PICMS  string `xml:"pICMS"`
							VICMS  string `xml:"vICMS"`
						} `xml:"ICMS20"`
					} `xml:"ICMS"`
					IPI struct {
						Text  string `xml:",chardata"`
						CEnq  string `xml:"cEnq"`
						IPINT struct {
							Text string `xml:",chardata"`
							CST  string `xml:"CST"`
						} `xml:"IPINT"`
						IPITrib struct {
							Text string `xml:",chardata"`
							CST  string `xml:"CST"`
							VBC  string `xml:"vBC"`
							PIPI string `xml:"pIPI"`
							VIPI string `xml:"vIPI"`
						} `xml:"IPITrib"`
					} `xml:"IPI"`
					PIS struct {
						Text    string `xml:",chardata"`
						PISAliq struct {
							Text string  `xml:",chardata"`
							CST  string  `xml:"CST"`
							VBC  float64 `xml:"vBC"`
							PPIS string  `xml:"pPIS"`
							VPIS string  `xml:"vPIS"`
						} `xml:"PISAliq"`
						PISNT struct {
							Text string `xml:",chardata"`
							CST  string `xml:"CST"`
						} `xml:"PISNT"`
						PISOutr struct {
							Text string `xml:",chardata"`
							CST  string `xml:"CST"`
							VBC  string `xml:"vBC"`
							PPIS string `xml:"pPIS"`
							VPIS string `xml:"vPIS"`
						} `xml:"PISOutr"`
					} `xml:"PIS"`
					COFINS struct {
						Text       string `xml:",chardata"`
						COFINSAliq struct {
							Text    string  `xml:",chardata"`
							CST     string  `xml:"CST"`
							VBC     float64 `xml:"vBC"`
							PCOFINS string  `xml:"pCOFINS"`
							VCOFINS string  `xml:"vCOFINS"`
						} `xml:"COFINSAliq"`
						COFINSNT struct {
							Text string `xml:",chardata"`
							CST  string `xml:"CST"`
						} `xml:"COFINSNT"`
						COFINSOutr struct {
							Text    string `xml:",chardata"`
							CST     string `xml:"CST"`
							VBC     string `xml:"vBC"`
							PCOFINS string `xml:"pCOFINS"`
							VCOFINS string `xml:"vCOFINS"`
						} `xml:"COFINSOutr"`
					} `xml:"COFINS"`
					II struct {
						Text     string `xml:",chardata"`
						VBC      string `xml:"vBC"`
						VDespAdu string `xml:"vDespAdu"`
						VII      string `xml:"vII"`
						VIOF     string `xml:"vIOF"`
					} `xml:"II"`
				} `xml:"imposto"`
			} `xml:"det"`
			Total struct {
				Text    string `xml:",chardata"`
				ICMSTot struct {
					Text       string `xml:",chardata"`
					VBC        string `xml:"vBC"`
					VICMS      string `xml:"vICMS"`
					VICMSDeson string `xml:"vICMSDeson"`
					VFCP       string `xml:"vFCP"`
					VBCST      string `xml:"vBCST"`
					VST        string `xml:"vST"`
					VFCPST     string `xml:"vFCPST"`
					VFCPSTRet  string `xml:"vFCPSTRet"`
					VProd      string `xml:"vProd"`
					VFrete     string `xml:"vFrete"`
					VSeg       string `xml:"vSeg"`
					VDesc      string `xml:"vDesc"`
					VII        string `xml:"vII"`
					VIPI       string `xml:"vIPI"`
					VIPIDevol  string `xml:"vIPIDevol"`
					VPIS       string `xml:"vPIS"`
					VCOFINS    string `xml:"vCOFINS"`
					VOutro     string `xml:"vOutro"`
					VNF        string `xml:"vNF"`
					VTotTrib   string `xml:"vTotTrib"`
				} `xml:"ICMSTot"`
			} `xml:"total"`
			Transp struct {
				Text       string `xml:",chardata"`
				ModFrete   string `xml:"modFrete"`
				Transporta struct {
					Text   string `xml:",chardata"`
					CNPJ   string `xml:"CNPJ"`
					XNome  string `xml:"xNome"`
					IE     string `xml:"IE"`
					XEnder string `xml:"xEnder"`
					XMun   string `xml:"xMun"`
					UF     string `xml:"UF"`
				} `xml:"transporta"`
				Vol struct {
					Text  string `xml:",chardata"`
					QVol  string `xml:"qVol"`
					Esp   string `xml:"esp"`
					Marca string `xml:"marca"`
					PesoL string `xml:"pesoL"`
					PesoB string `xml:"pesoB"`
				} `xml:"vol"`
			} `xml:"transp"`
			Cobr struct {
				Text string `xml:",chardata"`
				Fat  struct {
					Text  string `xml:",chardata"`
					NFat  string `xml:"nFat"`
					VOrig string `xml:"vOrig"`
					VDesc string `xml:"vDesc"`
					VLiq  string `xml:"vLiq"`
				} `xml:"fat"`
				Dup struct {
					Text  string `xml:",chardata"`
					NDup  string `xml:"nDup"`
					DVenc string `xml:"dVenc"`
					VDup  string `xml:"vDup"`
				} `xml:"dup"`
			} `xml:"cobr"`
			Pag struct {
				Text   string `xml:",chardata"`
				DetPag struct {
					Text   string `xml:",chardata"`
					IndPag string `xml:"indPag"`
					TPag   string `xml:"tPag"`
					VPag   string `xml:"vPag"`
				} `xml:"detPag"`
			} `xml:"pag"`
			InfAdic struct {
				Text   string `xml:",chardata"`
				InfCpl string `xml:"infCpl"`
			} `xml:"infAdic"`
			InfRespTec struct {
				Text     string `xml:",chardata"`
				CNPJ     string `xml:"CNPJ"`
				XContato string `xml:"xContato"`
				Email    string `xml:"email"`
				Fone     string `xml:"fone"`
			} `xml:"infRespTec"`
		} `xml:"infNFe"`
		Signature struct {
			Text       string `xml:",chardata"`
			Xmlns      string `xml:"xmlns,attr"`
			SignedInfo struct {
				Text                   string `xml:",chardata"`
				CanonicalizationMethod struct {
					Text      string `xml:",chardata"`
					Algorithm string `xml:"Algorithm,attr"`
				} `xml:"CanonicalizationMethod"`
				SignatureMethod struct {
					Text      string `xml:",chardata"`
					Algorithm string `xml:"Algorithm,attr"`
				} `xml:"SignatureMethod"`
				Reference struct {
					Text       string `xml:",chardata"`
					URI        string `xml:"URI,attr"`
					Transforms struct {
						Text      string `xml:",chardata"`
						Transform []struct {
							Text      string `xml:",chardata"`
							Algorithm string `xml:"Algorithm,attr"`
						} `xml:"Transform"`
					} `xml:"Transforms"`
					DigestMethod struct {
						Text      string `xml:",chardata"`
						Algorithm string `xml:"Algorithm,attr"`
					} `xml:"DigestMethod"`
					DigestValue string `xml:"DigestValue"`
				} `xml:"Reference"`
			} `xml:"SignedInfo"`
			SignatureValue string `xml:"SignatureValue"`
			KeyInfo        struct {
				Text     string `xml:",chardata"`
				X509Data struct {
					Text            string `xml:",chardata"`
					X509Certificate string `xml:"X509Certificate"`
				} `xml:"X509Data"`
			} `xml:"KeyInfo"`
		} `xml:"Signature"`
	} `xml:"NFe"`
	ProtNFe struct {
		Text    string `xml:",chardata"`
		Versao  string `xml:"versao,attr"`
		InfProt struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"Id,attr"`
			TpAmb    string `xml:"tpAmb"`
			VerAplic string `xml:"verAplic"`
			ChNFe    string `xml:"chNFe"`
			DhRecbto string `xml:"dhRecbto"`
			NProt    string `xml:"nProt"`
			DigVal   string `xml:"digVal"`
			CStat    string `xml:"cStat"`
			XMotivo  string `xml:"xMotivo"`
		} `xml:"infProt"`
	} `xml:"protNFe"`
}

func (n *NfeProc) Popula(byteValue []byte) error {
	err := xml.Unmarshal(byteValue, &n)
	if err != nil {
		return err
	}
	return nil
}
