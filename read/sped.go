package read

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/chapzin/parse-efd-fiscal/exec"
	"github.com/jinzhu/gorm"
)

// RecursiveSpeds processa recursivamente arquivos SPED em um diretório
func RecursiveSpeds(db *gorm.DB, path string, digito string, expectedCNPJ string) error {
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

				if err := processFile(pc, filepath.Join(path, file.Name()), expectedCNPJ); err != nil {
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
func processFile(pc *processControl, path string, expectedCNPJ string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regs := &exec.Regs{
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
		linha := decodeSpedLine(scanner.Bytes())
		campos := strings.Split(linha, "|")
		if len(campos) > 1 {
			if campos[1] == "0000" && !belongsToCNPJ(campos, expectedCNPJ) {
				return nil
			}
			if err := exec.TrataLinha(campos[1], linha, regs, pc.db); err != nil {
				return fmt.Errorf("erro ao processar linha: %v", err)
			}
		}
	}

	return scanner.Err()
}

func decodeSpedLine(line []byte) string {
	if utf8.Valid(line) {
		return string(line)
	}

	runes := make([]rune, 0, len(line))
	for _, b := range line {
		runes = append(runes, rune(b))
	}
	return string(runes)
}

func belongsToCNPJ(campos []string, expectedCNPJ string) bool {
	if expectedCNPJ == "" {
		return true
	}
	if len(campos) <= 8 {
		return false
	}
	return normalizeCNPJ(campos[7]) == expectedCNPJ || normalizeCNPJ(campos[8]) == expectedCNPJ
}

func normalizeCNPJ(value string) string {
	digits := make([]rune, 0, len(value))
	for _, r := range value {
		if r >= '0' && r <= '9' {
			digits = append(digits, r)
		}
	}
	return string(digits)
}

func isTXT(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".txt"
}
