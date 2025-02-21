package cupomfiscal

import (
	"encoding/xml"
	"fmt"
)

// CFe2 representa a estrutura principal do CFe.
type CFeXML struct {
	ID        int64     `bun:"id,pk,autoincrement" xml:"-"`
	XMLName   xml.Name  `xml:"CFe"`
	Text      string    `xml:",chardata"`
	InfCFe    InfCFe    `xml:"infCFe"`
	Signature Signature `xml:"Signature"`
}

// InfCFe contém as informações detalhadas do CFe.
type InfCFe struct {
	Text           string     `xml:",chardata"`
	VersaoSB       string     `xml:"versaoSB,attr"`
	VersaoDadosEnt string     `xml:"versaoDadosEnt,attr"`
	Versao         string     `xml:"versao,attr"`
	ID             string     `xml:"Id,attr"`
	Ide            IdeCfe     `xml:"ide"`
	Emit           EmitCfe    `xml:"emit"`
	Dest           DestCfe    `xml:"dest"`
	Det            []DetCfe   `xml:"det"`
	Total          TotalCfe   `xml:"total"`
	Pgto           PgtoCfe    `xml:"pgto"`
	InfAdic        InfAdicCfe `xml:"infAdic"`
}

// Ide contém as informações de identificação.
type IdeCfe struct {
	Text             string `xml:",chardata"`
	CUF              string `xml:"cUF"`
	CNF              string `xml:"cNF"`
	Mod              string `xml:"mod"`
	NserieSAT        string `xml:"nserieSAT"`
	NCFe             string `xml:"nCFe"`
	DEmi             string `xml:"dEmi"`
	HEmi             string `xml:"hEmi"`
	CDV              string `xml:"cDV"`
	TpAmb            string `xml:"tpAmb"`
	CNPJ             string `xml:"CNPJ"`
	SignAC           string `xml:"signAC"`
	AssinaturaQRCODE string `xml:"assinaturaQRCODE"`
	NumeroCaixa      string `xml:"numeroCaixa"`
}

// Emit contém as informações do emissor.
type EmitCfe struct {
	Text        string       `xml:",chardata"`
	CNPJ        string       `xml:"CNPJ"`
	XNome       string       `xml:"xNome"`
	XFant       string       `xml:"xFant"`
	EnderEmit   EnderEmitCfe `xml:"enderEmit"`
	IE          string       `xml:"IE"`
	CRegTrib    string       `xml:"cRegTrib"`
	IndRatISSQN string       `xml:"indRatISSQN"`
}

// EnderEmit contém o endereço do emissor.
type EnderEmitCfe struct {
	Text    string `xml:",chardata"`
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XCpl    string `xml:"xCpl"`
	XBairro string `xml:"xBairro"`
	XMun    string `xml:"xMun"`
	CEP     string `xml:"CEP"`
}

// Dest contém as informações do destinatário.
type DestCfe struct {
	Text  string `xml:",chardata"`
	XNome string `xml:"xNome"`
}

// Det representa os detalhes dos produtos e impostos.
type DetCfe struct {
	Text    string     `xml:",chardata"`
	NItem   string     `xml:"nItem,attr"`
	Prod    ProdCfe    `xml:"prod"`
	Imposto ImpostoCfe `xml:"imposto"`
}

// Prod contém informações do produto.
type ProdCfe struct {
	Text     string `xml:",chardata"`
	CProd    string `xml:"cProd"`
	CEAN     string `xml:"cEAN"`
	XProd    string `xml:"xProd"`
	NCM      string `xml:"NCM"`
	CFOP     string `xml:"CFOP"`
	UCom     string `xml:"uCom"`
	QCom     string `xml:"qCom"`
	VUnCom   string `xml:"vUnCom"`
	VProd    string `xml:"vProd"`
	IndRegra string `xml:"indRegra"`
	VItem    string `xml:"vItem"`
	VDesc    string `xml:"vDesc"`
	VOutro   string `xml:"vOutro"`
}

// Imposto contém informações sobre impostos.
type ImpostoCfe struct {
	Text       string    `xml:",chardata"`
	VItem12741 string    `xml:"vItem12741"`
	ICMS       ICMSCfe   `xml:"ICMS"`
	PIS        PISCfe    `xml:"PIS"`
	COFINS     COFINSCfe `xml:"COFINS"`
}

// ICMS contém informações sobre ICMS.
type ICMSCfe struct {
	Text      string `xml:",chardata"`
	ICMS00    *ICMSC `xml:"ICMS00"`
	ICMS40    *ICMSC `xml:"ICMS40"`
	ICMSSN102 *ICMSN `xml:"ICMSSN102"`
	ICMSSN500 *ICMSN `xml:"ICMSSN500"`
	ICMSSN900 *ICMSN `xml:"ICMSSN900"`
}

type ICMSC struct {
	Text  string `xml:",chardata"`
	Orig  string `xml:"Orig"`
	CST   string `xml:"CST"`
	PIcms string `xml:"pICMS"`
	VIcms string `xml:"vICMS"`
}

type ICMSN struct {
	Text  string `xml:",chardata"`
	Orig  string `xml:"Orig"`
	CSOSN string `xml:"CSOSN"`
}

// PIS contém informações sobre PIS.
type PISCfe struct {
	Text    string  `xml:",chardata"`
	PISAliq PISAliq `xml:"PISAliq"`
}

// PISAliq contém detalhes do PIS.
type PISAliq struct {
	Text string `xml:",chardata"`
	CST  string `xml:"CST"`
	VBC  string `xml:"vBC"`
	PPIS string `xml:"pPIS"`
	VPIS string `xml:"vPIS"`
}

// COFINS contém informações sobre COFINS.
type COFINSCfe struct {
	Text       string     `xml:",chardata"`
	COFINSAliq COFINSAliq `xml:"COFINSAliq"`
}

// COFINSAliq contém detalhes do COFINS.
type COFINSAliq struct {
	Text    string `xml:",chardata"`
	CST     string `xml:"CST"`
	VBC     string `xml:"vBC"`
	PCOFINS string `xml:"pCOFINS"`
	VCOFINS string `xml:"vCOFINS"`
}

// Total contém informações totais da nota fiscal.
type TotalCfe struct {
	Text         string  `xml:",chardata"`
	ICMSTot      ICMSTot `xml:"ICMSTot"`
	VCFe         string  `xml:"vCFe"`
	VCFeLei12741 string  `xml:"vCFeLei12741"`
}

// ICMSTot contém totais de ICMS.
type ICMSTot struct {
	Text      string `xml:",chardata"`
	VICMS     string `xml:"vICMS"`
	VProd     string `xml:"vProd"`
	VDesc     string `xml:"vDesc"`
	VPIS      string `xml:"vPIS"`
	VCOFINS   string `xml:"vCOFINS"`
	VPISST    string `xml:"vPISST"`
	VCOFINSST string `xml:"vCOFINSST"`
	VOutro    string `xml:"vOutro"`
}

// Pgto contém informações sobre o pagamento.
type PgtoCfe struct {
	Text   string `xml:",chardata"`
	MP     MPCfe  `xml:"MP"`
	VTroco string `xml:"vTroco"`
}

// MP contém informações sobre meios de pagamento.
type MPCfe struct {
	Text string `xml:",chardata"`
	CMP  string `xml:"cMP"`
	VMP  string `xml:"vMP"`
}

// InfAdic contém informações adicionais.
type InfAdicCfe struct {
	Text   string `xml:",chardata"`
	InfCpl string `xml:"infCpl"`
}

// Signature contém informações sobre a assinatura.
type Signature struct {
	Text           string     `xml:",chardata"`
	Xmlns          string     `xml:"xmlns,attr"`
	SignedInfo     SignedInfo `xml:"SignedInfo"`
	SignatureValue string     `xml:"SignatureValue"`
	KeyInfo        KeyInfo    `xml:"KeyInfo"`
}

// SignedInfo contém informações sobre a assinatura.
type SignedInfo struct {
	Text                   string                 `xml:",chardata"`
	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod"`
	SignatureMethod        SignatureMethod        `xml:"SignatureMethod"`
	Reference              Reference              `xml:"Reference"`
}

// CanonicalizationMethod define o método de canonicalização.
type CanonicalizationMethod struct {
	Text      string `xml:",chardata"`
	Algorithm string `xml:"Algorithm,attr"`
}

// SignatureMethod define o método de assinatura.
type SignatureMethod struct {
	Text      string `xml:",chardata"`
	Algorithm string `xml:"Algorithm,attr"`
}

// Reference contém informações de referência na assinatura.
type Reference struct {
	Text         string       `xml:",chardata"`
	URI          string       `xml:"URI,attr"`
	Transforms   Transforms   `xml:"Transforms"`
	DigestMethod DigestMethod `xml:"DigestMethod"`
	DigestValue  string       `xml:"DigestValue"`
}

// Transforms contém transformações aplicadas.
type Transforms struct {
	Text      string      `xml:",chardata"`
	Transform []Transform `xml:"Transform"`
}

// Transform define uma transformação.
type Transform struct {
	Text      string `xml:",chardata"`
	Algorithm string `xml:"Algorithm,attr"`
}

// DigestMethod define o método de digestão.
type DigestMethod struct {
	Text      string `xml:",chardata"`
	Algorithm string `xml:"Algorithm,attr"`
}

// KeyInfo contém informações da chave.
type KeyInfo struct {
	Text     string   `xml:",chardata"`
	X509Data X509Data `xml:"X509Data"`
}

// X509Data contém informações do certificado X509.
type X509Data struct {
	Text            string `xml:",chardata"`
	X509Certificate string `xml:"X509Certificate"`
}

func (cfe *CFeXML) Popula(byteValue []byte) error {
	err := xml.Unmarshal(byteValue, &cfe)
	if err != nil {
		return fmt.Errorf("erro ao fazer unmarshal do xml: %w", err)
	}

	return nil
}
