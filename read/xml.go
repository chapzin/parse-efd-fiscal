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

// Funções auxiliares
func isXML(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))

	return ext == ".xml"
}
