package service

import (
	"fmt"
	"log"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Controllers/repository"
	"github.com/chapzin/parse-efd-fiscal/Models"
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoC"
)

type InventoryService struct {
	repo *repository.InventoryRepository
}

func NewInventoryService(repo *repository.InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) ProcessConversionFactors() error {
	log.Printf("Iniciando processamento de fatores de conversão em %v", time.Now())
	defer log.Printf("Finalizado processamento de fatores de conversão em %v", time.Now())

	if err := s.repo.DeleteUnitaryConversionFactors(); err != nil {
		return fmt.Errorf("erro ao deletar fatores de conversão unitários: %v", err)
	}

	factors, err := s.repo.GetUnprocessedConversionFactors()
	if err != nil {
		return fmt.Errorf("erro ao obter fatores de conversão: %v", err)
	}

	for _, factor := range factors {
		if err := s.processConversionFactor(factor); err != nil {
			log.Printf("Erro ao processar fator de conversão %s: %v", factor.CodItem, err)
			continue
		}
	}

	return nil
}

func (s *InventoryService) processConversionFactor(factor Bloco0.Reg0220) error {
	items, err := s.repo.GetC170ByItemAndUnit(factor.CodItem, factor.UnidConv, factor.DtIni, factor.DtFin)
	if err != nil {
		return fmt.Errorf("erro ao obter itens para conversão: %v", err)
	}

	for _, item := range items {
		newItem := BlocoC.RegC170{
			Qtd:  item.Qtd * factor.FatConv,
			Unid: factor.UnidCod,
		}

		if err := s.repo.UpdateC170(item.ID, item.CodItem, newItem); err != nil {
			return fmt.Errorf("erro ao atualizar item %s: %v", item.CodItem, err)
		}

		if err := s.repo.MarkFactorAsProcessed(factor.ID); err != nil {
			return fmt.Errorf("erro ao marcar fator como processado: %v", err)
		}
	}

	return nil
}

func (s *InventoryService) DeleteCancelledInvoiceItems(dtIni, dtFin string) error {
	log.Printf("Iniciando exclusão de itens de notas canceladas em %v", time.Now())
	defer log.Printf("Finalizada exclusão de itens de notas canceladas em %v", time.Now())

	invoices, err := s.repo.GetCancelledInvoices(dtIni, dtFin)
	if err != nil {
		return fmt.Errorf("erro ao obter notas canceladas: %v", err)
	}

	allInvoices, err := s.repo.GetAllInvoices()
	if err != nil {
		return fmt.Errorf("erro ao obter todas as notas: %v", err)
	}

	// Processa notas canceladas
	for _, invoice := range invoices {
		for _, fullInvoice := range allInvoices {
			if invoice.ChvNfe == fullInvoice.ChNFe {
				if err := s.repo.DeleteInvoiceItems(fullInvoice.ID); err != nil {
					log.Printf("Erro ao deletar itens da nota %s: %v", fullInvoice.ChNFe, err)
					continue
				}
			}
		}
	}

	// Limpa itens marcados como deletados
	if err := s.repo.DeleteSoftDeletedItems(); err != nil {
		return fmt.Errorf("erro ao limpar itens deletados: %v", err)
	}

	return nil
}

func (s *InventoryService) ProcessInventory(anoInicial, anoFinal int) error {
	log.Printf("Iniciando processamento do inventário para período %d-%d", anoInicial, anoFinal)
	defer log.Printf("Finalizado processamento do inventário")

	// Obtém produtos distintos
	products, err := s.repo.GetDistinctProducts()
	if err != nil {
		return fmt.Errorf("erro ao obter produtos: %v", err)
	}

	// Cria itens de inventário iniciais
	for _, product := range products {
		item := Models.Inventario{
			Codigo:    product.CodItem,
			Descricao: product.DescrItem,
			Tipo:      product.TipoItem,
			UnidInv:   product.UnidInv,
			Ncm:       product.CodNcm,
		}

		if err := s.repo.CreateInventoryItem(item); err != nil {
			log.Printf("Erro ao criar item de inventário %s: %v", product.CodItem, err)
			continue
		}
	}

	// Processa anos
	for ano := anoInicial; ano <= anoFinal; ano++ {
		if err := s.processInventoryYear(ano); err != nil {
			log.Printf("Erro ao processar ano %d: %v", ano, err)
			continue
		}
	}

	// Limpa itens não relevantes
	if err := s.repo.DeleteNonSalesItems(); err != nil {
		return fmt.Errorf("erro ao limpar itens não relevantes: %v", err)
	}

	return nil
}

func (s *InventoryService) processInventoryYear(ano int) error {
	// Datas do período
	startDate := time.Date(ano, 1, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(ano, 12, 31, 23, 59, 59, 0, time.Local)

	// Obtém dados do inventário
	items, err := s.repo.GetAllInventoryItems()
	if err != nil {
		return fmt.Errorf("erro ao obter itens do inventário: %v", err)
	}

	// Processa cada item
	for _, item := range items {
		// Calcula entradas
		entradas, err := s.repo.GetEntradasPeriodo(item.Codigo, startDate, endDate)
		if err != nil {
			log.Printf("Erro ao obter entradas do item %s: %v", item.Codigo, err)
			continue
		}

		// Calcula saídas
		saidas, err := s.repo.GetSaidasPeriodo(item.Codigo, startDate, endDate)
		if err != nil {
			log.Printf("Erro ao obter saídas do item %s: %v", item.Codigo, err)
			continue
		}

		// Atualiza item com os valores calculados
		updatedItem := s.calculateInventoryValues(item, entradas, saidas, ano)
		if err := s.repo.UpdateInventoryItem(item.Codigo, updatedItem); err != nil {
			log.Printf("Erro ao atualizar item %s: %v", item.Codigo, err)
			continue
		}
	}

	return nil
}

func (s *InventoryService) calculateInventoryValues(item Models.Inventario, entradas, saidas float64, ano int) Models.Inventario {
	// Copia o item para não modificar o original
	updatedItem := item

	// Calcula valores baseado no ano absoluto
	switch ano {
	case 2012: // Primeiro ano
		updatedItem.EntradasAno2 = entradas
		updatedItem.SaidasAno2 = saidas
		if entradas > 0 {
			updatedItem.VlUnitEntAno2 = updatedItem.VlTotalEntradasAno2 / entradas
		}
		if saidas > 0 {
			updatedItem.VlUnitSaiAno2 = updatedItem.VlTotalSaidasAno2 / saidas
		}
		updatedItem.DiferencasAno2 = (updatedItem.InvFinalAno1 + entradas) - (saidas + updatedItem.InvFinalAno2)

	case 2013:
		updatedItem.EntradasAno3 = entradas
		updatedItem.SaidasAno3 = saidas
		if entradas > 0 {
			updatedItem.VlUnitEntAno3 = updatedItem.VlTotalEntradasAno3 / entradas
		}
		if saidas > 0 {
			updatedItem.VlUnitSaiAno3 = updatedItem.VlTotalSaidasAno3 / saidas
		}
		updatedItem.DiferencasAno3 = (updatedItem.InvFinalAno2 + entradas) - (saidas + updatedItem.InvFinalAno3)

	case 2014:
		updatedItem.EntradasAno4 = entradas
		updatedItem.SaidasAno4 = saidas
		if entradas > 0 {
			updatedItem.VlUnitEntAno4 = updatedItem.VlTotalEntradasAno4 / entradas
		}
		if saidas > 0 {
			updatedItem.VlUnitSaiAno4 = updatedItem.VlTotalSaidasAno4 / saidas
		}
		updatedItem.DiferencasAno4 = (updatedItem.InvFinalAno3 + entradas) - (saidas + updatedItem.InvFinalAno4)

	case 2015:
		updatedItem.EntradasAno5 = entradas
		updatedItem.SaidasAno5 = saidas
		if entradas > 0 {
			updatedItem.VlUnitEntAno5 = updatedItem.VlTotalEntradasAno5 / entradas
		}
		if saidas > 0 {
			updatedItem.VlUnitSaiAno5 = updatedItem.VlTotalSaidasAno5 / saidas
		}
		updatedItem.DiferencasAno5 = (updatedItem.InvFinalAno4 + entradas) - (saidas + updatedItem.InvFinalAno5)

	case 2016:
		updatedItem.EntradasAno6 = entradas
		updatedItem.SaidasAno6 = saidas
		if entradas > 0 {
			updatedItem.VlUnitEntAno6 = updatedItem.VlTotalEntradasAno6 / entradas
		}
		if saidas > 0 {
			updatedItem.VlUnitSaiAno6 = updatedItem.VlTotalSaidasAno6 / saidas
		}
		updatedItem.DiferencasAno6 = (updatedItem.InvFinalAno5 + entradas) - (saidas + updatedItem.InvFinalAno6)

	default:
		log.Printf("Ano %d fora do intervalo suportado (2012-2016)", ano)
	}

	return updatedItem
}

// ... outros métodos do serviço
