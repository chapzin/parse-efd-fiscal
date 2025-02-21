package repository

import (
	"fmt"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models"
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoC"
	"github.com/chapzin/parse-efd-fiscal/Models/NotaFiscal"
	"github.com/jinzhu/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

// Métodos para Fator de Conversão
func (r *InventoryRepository) DeleteUnitaryConversionFactors() error {
	return r.db.Exec("DELETE FROM reg_0220 WHERE fat_conv=1").Error
}

func (r *InventoryRepository) GetUnprocessedConversionFactors() ([]Bloco0.Reg0220, error) {
	var factors []Bloco0.Reg0220
	if err := r.db.Where("feito = ?", 0).Find(&factors).Error; err != nil {
		return nil, err
	}
	return factors, nil
}

func (r *InventoryRepository) UpdateC170(id uint, codItem string, newItem BlocoC.RegC170) error {
	return r.db.Table("reg_c170").Where("id = ? and cod_item = ?", id, codItem).Update(&newItem).Error
}

func (r *InventoryRepository) MarkFactorAsProcessed(id uint) error {
	return r.db.Table("reg_0220").Where("id = ?", id).Update("feito", "1").Error
}

// Métodos para Notas Canceladas
func (r *InventoryRepository) GetCancelledInvoices(dtIni, dtFin string) ([]BlocoC.RegC100, error) {
	var invoices []BlocoC.RegC100
	if err := r.db.Where("cod_sit <> ? and dt_ini >= ? and dt_ini <= ?", "00", dtIni, dtFin).Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *InventoryRepository) GetAllInvoices() ([]NotaFiscal.NotaFiscal, error) {
	var invoices []NotaFiscal.NotaFiscal
	if err := r.db.Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *InventoryRepository) DeleteInvoiceItems(notaID uint) error {
	return r.db.Exec("Delete from items where nota_fiscal_id=?", notaID).Error
}

func (r *InventoryRepository) DeleteSoftDeletedItems() error {
	return r.db.Exec("DELETE FROM items WHERE deleted_at is not null").Error
}

// Métodos para Inventário
func (r *InventoryRepository) GetDistinctProducts() ([]Bloco0.Reg0200, error) {
	var products []Bloco0.Reg0200
	if err := r.db.Where("tipo_item=00").Select("distinct cod_item,descr_item,tipo_item,unid_inv").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *InventoryRepository) CreateInventoryItem(item Models.Inventario) error {
	return r.db.Create(&item).Error
}

func (r *InventoryRepository) GetAllInventoryItems() ([]Models.Inventario, error) {
	var items []Models.Inventario
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *InventoryRepository) UpdateInventoryItem(codItem string, item Models.Inventario) error {
	return r.db.Table("inventarios").Where("codigo = ?", codItem).Update(&item).Error
}

func (r *InventoryRepository) DeleteNonSalesItems() error {
	return r.db.Exec("Delete from inventarios where tipo <> '00'").Error
}

func (r *InventoryRepository) GetC170ByItemAndUnit(codItem, unid string, dtIni, dtFin time.Time) ([]BlocoC.RegC170, error) {
	var items []BlocoC.RegC170
	if err := r.db.Where("cod_item = ? and unid = ? and dt_ini = ? and dt_fin = ?",
		codItem, unid, dtIni.Format("2006-01-02"), dtFin.Format("2006-01-02")).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// GetEntradasPeriodo retorna o total de entradas de um item em um período
func (r *InventoryRepository) GetEntradasPeriodo(codItem string, dtIni, dtFin time.Time) (float64, error) {
	// Entradas do SPED (C170)
	var totalC170 float64
	if err := r.db.Table("reg_c170").
		Where("cod_item = ? AND entrada_saida = '0' AND dt_ini >= ? AND dt_fin <= ?",
			codItem, dtIni.Format("2006-01-02"), dtFin.Format("2006-01-02")).
		Select("COALESCE(SUM(qtd), 0)").Row().Scan(&totalC170); err != nil {
		return 0, fmt.Errorf("erro ao calcular total de entradas C170: %v", err)
	}

	// Entradas dos XMLs
	var totalXML float64
	if err := r.db.Table("items").
		Where("codigo = ? AND cfop < 3999 AND dt_emit >= ? AND dt_emit <= ?",
			codItem, dtIni.Format("2006-01-02"), dtFin.Format("2006-01-02")).
		Select("COALESCE(SUM(qtd), 0)").Row().Scan(&totalXML); err != nil {
		return 0, fmt.Errorf("erro ao calcular total de entradas XML: %v", err)
	}

	return totalC170 + totalXML, nil
}

// GetSaidasPeriodo retorna o total de saídas de um item em um período
func (r *InventoryRepository) GetSaidasPeriodo(codItem string, dtIni, dtFin time.Time) (float64, error) {
	// Saídas dos XMLs
	var totalXML float64
	if err := r.db.Table("items").
		Where("codigo = ? AND cfop > 3999 AND cfop <> 5929 AND cfop <> 6929 AND dt_emit >= ? AND dt_emit <= ?",
			codItem, dtIni.Format("2006-01-02"), dtFin.Format("2006-01-02")).
		Select("COALESCE(SUM(qtd), 0)").Row().Scan(&totalXML); err != nil {
		return 0, fmt.Errorf("erro ao calcular total de saídas XML: %v", err)
	}

	// Saídas do SPED (C425)
	var totalC425 float64
	if err := r.db.Table("reg_c425").
		Where("cod_item = ? AND dt_ini >= ? AND dt_ini <= ?",
			codItem, dtIni.Format("2006-01-02"), dtFin.Format("2006-01-02")).
		Select("COALESCE(SUM(qtd), 0)").Row().Scan(&totalC425); err != nil {
		return 0, fmt.Errorf("erro ao calcular total de saídas C425: %v", err)
	}

	// Saídas do SPED (C470)
	var totalC470 float64
	if err := r.db.Table("reg_c470").
		Where("cod_item = ? AND dt_ini >= ? AND dt_ini <= ?",
			codItem, dtIni.Format("2006-01-02"), dtFin.Format("2006-01-02")).
		Select("COALESCE(SUM(qtd), 0)").Row().Scan(&totalC470); err != nil {
		return 0, fmt.Errorf("erro ao calcular total de saídas C470: %v", err)
	}

	return totalXML + totalC425 + totalC470, nil
}

// ... outros métodos do repositório
