package main

import (

	"fmt"
	"encoding/xml"
	"github.com/clbanning/mxj"
	"io/ioutil"
)

type NotasFiscal struct {
	XMLName xml.Name	`xml:"nfeProc"`
	cUF string		`xml:"NFe>infNFe>ide>cUF"`

}


func main() {

	xmlFile, err := ioutil.ReadFile("23130141334079000760550010000060781002141849-procNFe.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	nfe, _ := mxj.NewMapXml(xmlFile)
	fmt.Println(nfe["nfeProc"])

}
