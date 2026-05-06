package read

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	cupomfiscal "github.com/chapzin/parse-efd-fiscal/Models/CupomFiscal"
	"github.com/chapzin/parse-efd-fiscal/Models/NotaFiscal"
	"github.com/chapzin/parse-efd-fiscal/exec"
	"github.com/jinzhu/gorm"
)

var ErrInvalidXML = fmt.Errorf("arquivo XML inválido")

// RecursiveXmls processa recursivamente arquivos XML em um diretório
func RecursiveXmls(db *gorm.DB, path string, digito string, expectedCNPJ string) error {
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

				if err := processXMLFile(pc, filepath.Join(path, file.Name()), expectedCNPJ); err != nil {
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
func processXMLFile(pc *processControl, path string, expectedCNPJ string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer file.Close()

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	xml := string(fileByte)

	switch {
	case strings.Contains(xml, "<nfeProc") && strings.Contains(xml, "</nfeProc>"):
		nfe := new(NotaFiscal.NfeProc)

		err := nfe.Popula(fileByte)
		if err != nil {
			return fmt.Errorf("erro ao popular struct: %w", err)
		}
		if !nfeBelongsToCNPJ(*nfe, expectedCNPJ) {
			return nil
		}

		err = exec.InsertXmlNfe(pc.db, *nfe, pc.digito)
		if err != nil {
			return fmt.Errorf("erro ao inserir xml no banco de dados: %w", err)
		}

	case strings.Contains(xml, "<CFe") && strings.Contains(xml, "</CFe>"):
		cfe := new(cupomfiscal.CFeXML)

		err := cfe.Popula(fileByte)
		if err != nil {
			return fmt.Errorf("erro ao popular struct: %w", err)
		}
		if !cfeBelongsToCNPJ(*cfe, expectedCNPJ) {
			return nil
		}

		err = exec.InsertXmlCfe(pc.db, *cfe)
		if err != nil {
			return fmt.Errorf("erro ao inserir xml no banco de dados: %w", err)
		}

		return nil

	default:
		return ErrInvalidXML
	}

	return nil
}

func nfeBelongsToCNPJ(nfe NotaFiscal.NfeProc, expectedCNPJ string) bool {
	if expectedCNPJ == "" {
		return true
	}
	return normalizeCNPJ(nfe.NFe.InfNFe.Emit.CNPJ) == expectedCNPJ || normalizeCNPJ(nfe.NFe.InfNFe.Dest.CNPJ) == expectedCNPJ
}

func cfeBelongsToCNPJ(cfe cupomfiscal.CFeXML, expectedCNPJ string) bool {
	if expectedCNPJ == "" {
		return true
	}
	return normalizeCNPJ(cfe.InfCFe.Emit.CNPJ) == expectedCNPJ
}

// Funções auxiliares
func isXML(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))

	return ext == ".xml"
}
