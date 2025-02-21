package Controllers

import (
	"strconv"
	"sync"
	"time"

	"github.com/chapzin/parse-efd-fiscal/Models"
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoC"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoH"
	"github.com/chapzin/parse-efd-fiscal/Models/NotaFiscal"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"github.com/tealeg/xlsx"

	"github.com/chapzin/parse-efd-fiscal/Controllers/repository"
	"github.com/chapzin/parse-efd-fiscal/Controllers/service"
)

type InventoryController struct {
	service *service.InventoryService
}

func NewInventoryController(db *gorm.DB) *InventoryController {
	repo := repository.NewInventoryRepository(db)
	svc := service.NewInventoryService(repo)
	return &InventoryController{service: svc}
}

func (c *InventoryController) ProcessarFatorConversao(db *gorm.DB) error {
	return c.service.ProcessConversionFactors()
}

func (c *InventoryController) DeletarItensNotasCanceladas(db *gorm.DB, dtIni, dtFin string) error {
	return c.service.DeleteCancelledInvoiceItems(dtIni, dtFin)
}

func ProcessarFatorConversao(db *gorm.DB, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	color.Green("Começo Processa Fator de Conversao %s", time.Now())
	db.Exec("DELETE FROM reg_0220 WHERE fat_conv=1")
	var fator []Bloco0.Reg0220
	db.Where("feito = ?", 0).Find(&fator)
	for _, vFator := range fator {
		c170 := []BlocoC.RegC170{}
		db.Where("cod_item = ? and unid = ? and dt_ini = ? and dt_fin = ?", vFator.CodItem, vFator.UnidConv, vFator.DtIni, vFator.DtFin).Find(&c170)
		for _, vC170 := range c170 {
			nvC170 := BlocoC.RegC170{}
			nvC170.Qtd = vC170.Qtd * vFator.FatConv
			nvC170.Unid = vFator.UnidCod
			db.Table("reg_c170").Where("id = ? and cod_item = ?", vC170.ID, vC170.CodItem).Update(&nvC170)
			nvFator := Bloco0.Reg0220{}
			nvFator.Feito = "1"
			db.Table("reg_0220").Where("id = ?", vFator.ID).Update(&nvFator)
		}
	}
	color.Green("Fim Processa Fator de Conversao %s", time.Now())
	wg.Done()
}

func DeletarItensNotasCanceladas(db *gorm.DB, dtIni string, dtFin string, wg *sync.WaitGroup) {
	color.Green("Começo Deleta itens notas canceladas %s", time.Now())
	var c100 []BlocoC.RegC100
	var nota []NotaFiscal.NotaFiscal
	db.Where("cod_sit <> ? and dt_ini >= ? and dt_ini <= ? ", "00", dtIni, dtFin).Find(&c100)
	db.Find(&nota)
	for _, v := range c100 {
		//fmt.Println(v.NumDoc)
		for _, vNota := range nota {
			func() {
				if v.ChvNfe == vNota.ChNFe {
					db.Exec("Delete from items where nota_fiscal_id=?", vNota.ID)
					//db.Where("nota_fiscal_id =?", vNota.ID).Delete(NotaFiscal.Item{})
				}
			}()

		}
	}
	db.Exec("DELETE FROM items WHERE deleted_at is not null")
	color.Green("Fim deleta itens notas canceladas %s", time.Now())
	wg.Done()
}

func PopularReg0200(db *gorm.DB, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	color.Green("Comeco popula reg0200 %s", time.Now())
	var reg0200 []Bloco0.Reg0200
	db.Where("tipo_item=00").Select("distinct cod_item,descr_item,tipo_item,unid_inv").Find(&reg0200)
	for _, v := range reg0200 {
		inv2 := Models.Inventario{
			Codigo:    v.CodItem,
			Descricao: v.DescrItem,
			Tipo:      v.TipoItem,
			UnidInv:   v.UnidInv,
			Ncm:       v.CodNcm,
		}
		db.NewRecord(inv2)
		db.Create(&inv2)

	}
	wg.Done()
	color.Green("Fim popula reg0200 %s", time.Now())
}

func PopularItensXmls(db *gorm.DB, wg *sync.WaitGroup) {
	color.Green("Comeco popula Itens Xmls %s", time.Now())
	var items []NotaFiscal.Item
	db.Select("distinct codigo,descricao").Find(&items)
	for _, v := range items {
		var inventario Models.Inventario
		db.Where("codigo=?", v.Codigo).First(&inventario)
		if inventario.Codigo == "" {
			inv2 := Models.Inventario{
				Codigo:    v.Codigo,
				Descricao: v.Descricao,
			}
			db.NewRecord(inv2)
			db.Create(&inv2)
		}
	}
	wg.Done()
	color.Green("Fim popula xmls %s", time.Now())

}

func PopularInventarios(AnoInicial int, AnoFinal int, wg *sync.WaitGroup, db *gorm.DB) {
	time.Sleep(1 * time.Second)
	color.Green("Começo popula Inventario %s", time.Now())
	qtdAnos := AnoFinal - AnoInicial
	ano1 := AnoInicial
	qtdAnos = qtdAnos + 1
	ano := 0
	for qtdAnos >= 0 {
		qtdAnos = qtdAnos - 1
		ano = ano + 1
		var regH010 []BlocoH.RegH010
		var inv []Models.Inventario
		AnoInicialString := strconv.Itoa(ano1)
		db.Where("dt_ini= ?", AnoInicialString+"-02-01").Find(&regH010)
		db.Find(&inv)
		for _, vInv := range inv {
			for _, vH010 := range regH010 {
				if vH010.CodItem == vInv.Codigo {
					inv3 := Models.Inventario{}
					switch ano {
					case 1:
						inv3.InvFinalAno1 = vH010.Qtd
						inv3.VlInvAno1 = vH010.VlUnit
					case 2:
						inv3.InvFinalAno2 = vH010.Qtd
						inv3.VlInvAno2 = vH010.VlUnit
					case 3:
						inv3.InvFinalAno3 = vH010.Qtd
						inv3.VlInvAno3 = vH010.VlUnit
					case 4:
						inv3.InvFinalAno4 = vH010.Qtd
						inv3.VlInvAno4 = vH010.VlUnit
					case 5:
						inv3.InvFinalAno5 = vH010.Qtd
						inv3.VlInvAno5 = vH010.VlUnit
					case 6:
						inv3.InvFinalAno6 = vH010.Qtd
						inv3.VlInvAno6 = vH010.VlUnit

					}
					db.Table("inventarios").Where("codigo = ?", vH010.CodItem).Update(&inv3)
				}
			}
		}
		ano1 = ano1 + 1
	}
	color.Green("Fim popula inventario %s", time.Now())
	wg.Done()
}

func PopularEntradas(AnoInicial int, AnoFinal int, wg *sync.WaitGroup, db *gorm.DB) {
	time.Sleep(1 * time.Second)
	color.Green("Começo popula entradas %s", time.Now())
	qtdAnos := AnoFinal - AnoInicial
	ano1 := AnoInicial
	ano := 0
	for qtdAnos >= 0 {
		qtdAnos = qtdAnos - 1
		ano = ano + 1
		AnoInicialString := strconv.Itoa(ano1)

		dtIni := AnoInicialString + "-01-01"
		dtFin := AnoInicialString + "-12-31"

		var inv []Models.Inventario
		var c170 []BlocoC.RegC170
		var itens []NotaFiscal.Item

		db.Find(&inv)
		db.Where("entrada_saida = ? and dt_ini >= ? and dt_fin <= ? ", "0", dtIni, dtFin).Find(&c170)
		db.Where("cfop < 3999 and dt_emit >= ? and dt_emit <= ?", dtIni, dtFin).Find(&itens)
		for _, vInv := range inv {
			var qtd_tot = 0.0
			var vl_tot = 0.0
			for _, vc170 := range c170 {
				if vc170.CodItem == vInv.Codigo {
					qtd_tot = qtd_tot + vc170.Qtd
					vl_tot = vl_tot + vc170.VlItem
				}
			}

			for _, vitens := range itens {
				if vitens.Codigo == vInv.Codigo {
					qtd_tot = qtd_tot + vitens.Qtd
					vl_tot = vl_tot + vitens.VTotal
				}
			}
			inv2 := Models.Inventario{}
			switch ano {
			case 1:
				inv2.EntradasAno2 = qtd_tot
				inv2.VlTotalEntradasAno2 = vl_tot
			case 2:
				inv2.EntradasAno3 = qtd_tot
				inv2.VlTotalEntradasAno3 = vl_tot
			case 3:
				inv2.EntradasAno4 = qtd_tot
				inv2.VlTotalEntradasAno4 = vl_tot
			case 4:
				inv2.EntradasAno5 = qtd_tot
				inv2.VlTotalEntradasAno5 = vl_tot
			case 5:
				inv2.EntradasAno6 = qtd_tot
				inv2.VlTotalEntradasAno6 = vl_tot

			}
			db.Table("inventarios").Where("codigo = ?", vInv.Codigo).Update(&inv2)
		}

		ano1 = ano1 + 1
	}
	color.Green("Fim popula entradas %s", time.Now())
	wg.Done()
}

func PopularSaidas(AnoInicial int, AnoFinal int, wg *sync.WaitGroup, db *gorm.DB) {
	time.Sleep(2 * time.Second)
	color.Green("Comeco popula saidas %s", time.Now())
	qtdAnos := AnoFinal - AnoInicial
	ano1 := AnoInicial
	ano := 0
	for qtdAnos >= 0 {
		qtdAnos = qtdAnos - 1
		ano = ano + 1
		AnoInicialString := strconv.Itoa(ano1)

		dtIni := AnoInicialString + "-01-01"
		dtFin := AnoInicialString + "-12-31"

		var inv []Models.Inventario
		var itens []NotaFiscal.Item
		var c425 []BlocoC.RegC425
		var c470 []BlocoC.RegC470

		db.Find(&inv)
		db.Where("cfop > 3999 and cfop <> 5929 and cfop <> 6929 and dt_emit >= ? and dt_emit <= ?", dtIni, dtFin).Find(&itens)
		db.Where("dt_ini >= ? and dt_ini <= ?", dtIni, dtFin).Find(&c425)
		db.Where("dt_ini >= ? and dt_ini <= ?", dtIni, dtFin).Find(&c470)

		for _, vInv := range inv {
			var qtd_saida = 0.0
			var vl_tot_saida = 0.0
			for _, vItens := range itens {
				if vItens.Codigo == vInv.Codigo {
					qtd_saida = qtd_saida + vItens.Qtd
					vl_tot_saida = vl_tot_saida + vItens.VTotal
				}
			}
			for _, vc425 := range c425 {
				if vc425.CodItem == vInv.Codigo {
					qtd_saida = qtd_saida + vc425.Qtd
					vl_tot_saida = vl_tot_saida + vc425.VlItem
				}
			}
			for _, vc470 := range c470 {
				if vc470.CodItem == vInv.Codigo {
					qtd_saida = qtd_saida + vc470.Qtd
					vl_tot_saida = vl_tot_saida + vc470.VlItem
				}
			}
			inv3 := Models.Inventario{}
			switch ano {
			case 1:
				inv3.SaidasAno2 = qtd_saida
				inv3.VlTotalSaidasAno2 = vl_tot_saida
			case 2:
				inv3.SaidasAno3 = qtd_saida
				inv3.VlTotalSaidasAno3 = vl_tot_saida
			case 3:
				inv3.SaidasAno4 = qtd_saida
				inv3.VlTotalSaidasAno4 = vl_tot_saida
			case 4:
				inv3.SaidasAno5 = qtd_saida
				inv3.VlTotalSaidasAno5 = vl_tot_saida
			case 5:
				inv3.SaidasAno6 = qtd_saida
				inv3.VlTotalSaidasAno6 = vl_tot_saida

			}
			db.Table("inventarios").Where("codigo = ?", vInv.Codigo).Update(&inv3)

		}
		ano1 = ano1 + 1
	}
	color.Green("Fim popula saidas %s", time.Now())
	wg.Done()
}

// fazer uma refactory completo recursive
func ProcessarDiferencas(db *gorm.DB) {
	db.Exec("Delete from inventarios where inv_inicial=0 and entradas=0 and vl_total_entradas=0 and saidas=0 and vl_total_saidas=0 and inv_final=0")
	var inv []Models.Inventario
	var reg0200 []Bloco0.Reg0200
	db.Select("distinct cod_item,descr_item,tipo_item,unid_inv").Find(&reg0200)
	db.Find(&inv)
	for _, vInv := range inv {
		inv3 := Models.Inventario{}
		// Calculando as diferencas
		diferencasAno2 := (vInv.InvFinalAno1 + vInv.EntradasAno2) - (vInv.SaidasAno2 + vInv.InvFinalAno2)
		diferencasAno3 := (vInv.InvFinalAno2 + vInv.EntradasAno3) - (vInv.SaidasAno3 + vInv.InvFinalAno3)
		diferencasAno4 := (vInv.InvFinalAno3 + vInv.EntradasAno4) - (vInv.SaidasAno4 + vInv.InvFinalAno4)
		diferencasAno5 := (vInv.InvFinalAno4 + vInv.EntradasAno5) - (vInv.SaidasAno5 + vInv.InvFinalAno5)
		diferencasAno6 := (vInv.InvFinalAno5 + vInv.EntradasAno6) - (vInv.SaidasAno6 + vInv.InvFinalAno6)

		// Calculando o valor unitário de entrada ano 2
		if vInv.VlTotalEntradasAno2 > 0 && vInv.EntradasAno2 > 0 {
			inv3.VlUnitEntAno2 = vInv.VlTotalEntradasAno2 / vInv.EntradasAno2
		} else if vInv.VlTotalEntradasAno2 == 0 && vInv.EntradasAno2 == 0 && vInv.VlInvAno1 > 0 {
			inv3.VlUnitEntAno2 = vInv.VlInvAno1
		} else if vInv.VlTotalEntradasAno2 == 0 && vInv.EntradasAno2 == 0 && vInv.VlInvAno1 == 0 && vInv.VlInvAno2 > 0 {
			inv3.VlUnitEntAno2 = vInv.VlInvAno2
		} else {
			inv3.VlUnitEntAno2 = 1
		}

		// Calculando o valor unitário de entrada ano 3
		if vInv.VlTotalEntradasAno3 > 0 && vInv.EntradasAno3 > 0 {
			inv3.VlUnitEntAno3 = vInv.VlTotalEntradasAno3 / vInv.EntradasAno3
		} else if vInv.VlTotalEntradasAno3 == 0 && vInv.EntradasAno3 == 0 && vInv.VlInvAno2 > 0 {
			inv3.VlUnitEntAno3 = vInv.VlInvAno2
		} else if vInv.VlTotalEntradasAno3 == 0 && vInv.EntradasAno3 == 0 && vInv.VlInvAno2 == 0 && vInv.VlInvAno3 > 0 {
			inv3.VlUnitEntAno3 = vInv.VlInvAno3
		} else {
			inv3.VlUnitEntAno3 = 1
		}

		// Calculando o valor unitário de entrada ano 4
		if vInv.VlTotalEntradasAno4 > 0 && vInv.EntradasAno4 > 0 {
			inv3.VlUnitEntAno4 = vInv.VlTotalEntradasAno4 / vInv.EntradasAno4
		} else if vInv.VlTotalEntradasAno4 == 0 && vInv.EntradasAno4 == 0 && vInv.VlInvAno3 > 0 {
			inv3.VlUnitEntAno4 = vInv.VlInvAno3
		} else if vInv.VlTotalEntradasAno4 == 0 && vInv.EntradasAno4 == 0 && vInv.VlInvAno3 == 0 && vInv.VlInvAno4 > 0 {
			inv3.VlUnitEntAno4 = vInv.VlInvAno4
		} else {
			inv3.VlUnitEntAno4 = 1
		}

		// Calculando o valor unitário de entrada ano 5
		if vInv.VlTotalEntradasAno5 > 0 && vInv.EntradasAno5 > 0 {
			inv3.VlUnitEntAno5 = vInv.VlTotalEntradasAno5 / vInv.EntradasAno5
		} else if vInv.VlTotalEntradasAno5 == 0 && vInv.EntradasAno5 == 0 && vInv.VlInvAno4 > 0 {
			inv3.VlUnitEntAno5 = vInv.VlInvAno5
		} else if vInv.VlTotalEntradasAno5 == 0 && vInv.EntradasAno5 == 0 && vInv.VlInvAno4 == 0 && vInv.VlInvAno5 > 0 {
			inv3.VlUnitEntAno5 = vInv.VlInvAno5
		} else {
			inv3.VlUnitEntAno5 = 1
		}

		// Calculando o valor unitário de entrada ano 6
		if vInv.VlTotalEntradasAno6 > 0 && vInv.EntradasAno6 > 0 {
			inv3.VlUnitEntAno6 = vInv.VlTotalEntradasAno6 / vInv.EntradasAno6
		} else if vInv.VlTotalEntradasAno6 == 0 && vInv.EntradasAno6 == 0 && vInv.VlInvAno5 > 0 {
			inv3.VlUnitEntAno6 = vInv.VlInvAno6
		} else if vInv.VlTotalEntradasAno6 == 0 && vInv.EntradasAno6 == 0 && vInv.VlInvAno5 == 0 && vInv.VlInvAno6 > 0 {
			inv3.VlUnitEntAno6 = vInv.VlInvAno6
		} else {
			inv3.VlUnitEntAno6 = 1
		}

		// Calculando o valor unitário de saida Ano 2
		if vInv.VlTotalSaidasAno2 > 0 && vInv.SaidasAno2 > 0 {
			inv3.VlUnitSaiAno2 = vInv.VlTotalSaidasAno2 / vInv.SaidasAno2
		} else if vInv.VlTotalSaidasAno2 == 0 && vInv.SaidasAno2 == 0 && vInv.VlInvAno1 > 0 {
			inv3.VlUnitSaiAno2 = vInv.VlInvAno1
		} else {
			inv3.VlUnitSaiAno2 = 0
		}

		// Calculando o valor unitário de saida Ano 3
		if vInv.VlTotalSaidasAno3 > 0 && vInv.SaidasAno3 > 0 {
			inv3.VlUnitSaiAno3 = vInv.VlTotalSaidasAno3 / vInv.SaidasAno3
		} else if vInv.VlTotalSaidasAno3 == 0 && vInv.SaidasAno3 == 0 && vInv.VlInvAno2 > 0 {
			inv3.VlUnitSaiAno3 = vInv.VlInvAno2
		} else {
			inv3.VlUnitSaiAno3 = 0
		}

		// Calculando o valor unitário de saida Ano 4
		if vInv.VlTotalSaidasAno4 > 0 && vInv.SaidasAno4 > 0 {
			inv3.VlUnitSaiAno4 = vInv.VlTotalSaidasAno4 / vInv.SaidasAno4
		} else if vInv.VlTotalSaidasAno4 == 0 && vInv.SaidasAno4 == 0 && vInv.VlInvAno3 > 0 {
			inv3.VlUnitSaiAno4 = vInv.VlInvAno3
		} else {
			inv3.VlUnitSaiAno4 = 0
		}

		// Calculando o valor unitário de saida Ano 5
		if vInv.VlTotalSaidasAno5 > 0 && vInv.SaidasAno5 > 0 {
			inv3.VlUnitSaiAno5 = vInv.VlTotalSaidasAno5 / vInv.SaidasAno5
		} else if vInv.VlTotalSaidasAno5 == 0 && vInv.SaidasAno5 == 0 && vInv.VlInvAno4 > 0 {
			inv3.VlUnitSaiAno5 = vInv.VlInvAno4
		} else {
			inv3.VlUnitSaiAno5 = 0
		}

		// Calculando o valor unitário de saida Ano 6
		if vInv.VlTotalSaidasAno6 > 0 && vInv.SaidasAno6 > 0 {
			inv3.VlUnitSaiAno6 = vInv.VlTotalSaidasAno6 / vInv.SaidasAno6
		} else if vInv.VlTotalSaidasAno6 == 0 && vInv.SaidasAno6 == 0 && vInv.VlInvAno5 > 0 {
			inv3.VlUnitSaiAno6 = vInv.VlInvAno5
		} else {
			inv3.VlUnitSaiAno6 = 0
		}

		// Adicionando Tipo e unidade de medida no inventario
		for _, v0200 := range reg0200 {
			if v0200.CodItem == vInv.Codigo {
				inv3.Tipo = v0200.TipoItem
				inv3.UnidInv = v0200.UnidInv
				inv3.Ncm = v0200.CodNcm
			}
		}
		inv3.DiferencasAno2 = diferencasAno2
		inv3.DiferencasAno3 = diferencasAno3
		inv3.DiferencasAno4 = diferencasAno4
		inv3.DiferencasAno5 = diferencasAno5
		inv3.DiferencasAno6 = diferencasAno6
		db.Table("inventarios").Where("codigo = ?", vInv.Codigo).Update(&inv3)
	}
	// Deleta tudo tipo de inventario que nao seja material de revenda
	db.Exec("Delete from inventarios where tipo <> '00'")
}

func ExcelAdd(db *gorm.DB, sheet *xlsx.Sheet) {
	var inv []Models.Inventario
	db.Find(&inv)
	for _, vInv := range inv {
		ExcelItens(sheet, vInv)
	}
}

func ColunaAdd(linha *xlsx.Row, string string) {
	cell := linha.AddCell()
	cell.Value = string
}

func ColunaAddFloat(linha *xlsx.Row, valor float64) {
	cell := linha.AddCell()
	cell.SetFloat(valor)
}
func ColunaAddFloatDif(linha *xlsx.Row, valor float64) {
	cell := linha.AddCell()

	var style = xlsx.NewStyle()
	if valor < 0 {
		style.Fill = *xlsx.NewFill("solid", "00FA8072", "00FA8072")
	} else if valor > 0 {
		style.Fill = *xlsx.NewFill("solid", "0087CEFA", "0087CEFA")
	} else {
		style.Fill = *xlsx.NewFill("solid", "009ACD32", "009ACD32")
	}
	cell.SetStyle(style)

	cell.SetFloat(valor)
}

func ExcelItens(sheet *xlsx.Sheet, inv Models.Inventario) {
	menu := sheet.AddRow()
	// Produto
	ColunaAdd(menu, inv.Codigo)
	ColunaAdd(menu, inv.Descricao)
	ColunaAdd(menu, inv.Tipo)
	ColunaAdd(menu, inv.UnidInv)
	ColunaAdd(menu, inv.Ncm)
	// Ano 1
	ColunaAddFloat(menu, inv.InvFinalAno1)
	ColunaAddFloat(menu, inv.VlInvAno1)
	// Ano 2
	ColunaAddFloat(menu, inv.EntradasAno2)
	ColunaAddFloat(menu, inv.VlTotalEntradasAno2)
	ColunaAddFloat(menu, inv.VlUnitEntAno2)
	ColunaAddFloat(menu, inv.SaidasAno2)
	ColunaAddFloat(menu, inv.VlTotalSaidasAno2)
	ColunaAddFloat(menu, inv.VlUnitSaiAno2)
	ColunaAddFloat(menu, inv.MargemAno2)
	ColunaAddFloat(menu, inv.InvFinalAno2)
	ColunaAddFloat(menu, inv.VlInvAno2)
	ColunaAddFloatDif(menu, inv.DiferencasAno2)
	// Ano 3
	ColunaAddFloat(menu, inv.EntradasAno3)
	ColunaAddFloat(menu, inv.VlTotalEntradasAno3)
	ColunaAddFloat(menu, inv.VlUnitEntAno3)
	ColunaAddFloat(menu, inv.SaidasAno3)
	ColunaAddFloat(menu, inv.VlTotalSaidasAno3)
	ColunaAddFloat(menu, inv.VlUnitSaiAno3)
	ColunaAddFloat(menu, inv.MargemAno3)
	ColunaAddFloat(menu, inv.InvFinalAno3)
	ColunaAddFloat(menu, inv.VlInvAno3)
	ColunaAddFloatDif(menu, inv.DiferencasAno3)
	// Ano 4
	ColunaAddFloat(menu, inv.EntradasAno4)
	ColunaAddFloat(menu, inv.VlTotalEntradasAno4)
	ColunaAddFloat(menu, inv.VlUnitEntAno4)
	ColunaAddFloat(menu, inv.SaidasAno4)
	ColunaAddFloat(menu, inv.VlTotalSaidasAno4)
	ColunaAddFloat(menu, inv.VlUnitSaiAno4)
	ColunaAddFloat(menu, inv.MargemAno4)
	ColunaAddFloat(menu, inv.InvFinalAno4)
	ColunaAddFloat(menu, inv.VlInvAno4)
	ColunaAddFloatDif(menu, inv.DiferencasAno4)
	// Ano 5
	ColunaAddFloat(menu, inv.EntradasAno5)
	ColunaAddFloat(menu, inv.VlTotalEntradasAno5)
	ColunaAddFloat(menu, inv.VlUnitEntAno5)
	ColunaAddFloat(menu, inv.SaidasAno5)
	ColunaAddFloat(menu, inv.VlTotalSaidasAno5)
	ColunaAddFloat(menu, inv.VlUnitSaiAno5)
	ColunaAddFloat(menu, inv.MargemAno5)
	ColunaAddFloat(menu, inv.InvFinalAno5)
	ColunaAddFloat(menu, inv.VlInvAno5)
	ColunaAddFloatDif(menu, inv.DiferencasAno5)
	// Ano 6
	ColunaAddFloat(menu, inv.EntradasAno6)
	ColunaAddFloat(menu, inv.VlTotalEntradasAno6)
	ColunaAddFloat(menu, inv.VlUnitEntAno6)
	ColunaAddFloat(menu, inv.SaidasAno6)
	ColunaAddFloat(menu, inv.VlTotalSaidasAno6)
	ColunaAddFloat(menu, inv.VlUnitSaiAno6)
	ColunaAddFloat(menu, inv.MargemAno6)
	ColunaAddFloat(menu, inv.InvFinalAno6)
	ColunaAddFloat(menu, inv.VlInvAno6)
	ColunaAddFloatDif(menu, inv.DiferencasAno6)
}

func ExcelMenu(sheet *xlsx.Sheet) {
	menu := sheet.AddRow()
	// Produtos
	ColunaAdd(menu, "Codigo")
	ColunaAdd(menu, "Descricao")
	ColunaAdd(menu, "Tipo")
	ColunaAdd(menu, "Unid_inv")
	ColunaAdd(menu, "NCM")
	// Ano 1
	ColunaAdd(menu, "InvFinalAno1")
	ColunaAdd(menu, "VlInvAno1")
	// Ano 2
	ColunaAdd(menu, "EntradasAno2")
	ColunaAdd(menu, "VlTotalEntradasAno2")
	ColunaAdd(menu, "VlUnitEntAno2")
	ColunaAdd(menu, "SaidasAno2")
	ColunaAdd(menu, "VlTotalSaidasAno2")
	ColunaAdd(menu, "VlUnitSaidaAno2")
	ColunaAdd(menu, "MargemAno2")
	ColunaAdd(menu, "InvFinalAno2")
	ColunaAdd(menu, "VlInvAno2")
	ColunaAdd(menu, "DiferencasAno2")
	// Ano 3
	ColunaAdd(menu, "EntradasAno3")
	ColunaAdd(menu, "VlTotalEntradasAno3")
	ColunaAdd(menu, "VlUnitEntAno3")
	ColunaAdd(menu, "SaidasAno3")
	ColunaAdd(menu, "VlTotalSaidasAno3")
	ColunaAdd(menu, "VlUnitSaidaAno3")
	ColunaAdd(menu, "MargemAno3")
	ColunaAdd(menu, "InvFinalAno3")
	ColunaAdd(menu, "VlInvAno3")
	ColunaAdd(menu, "DiferencasAno3")
	// Ano 4
	ColunaAdd(menu, "EntradasAno4")
	ColunaAdd(menu, "VlTotalEntradasAno4")
	ColunaAdd(menu, "VlUnitEntAno4")
	ColunaAdd(menu, "SaidasAno4")
	ColunaAdd(menu, "VlTotalSaidasAno4")
	ColunaAdd(menu, "VlUnitSaidaAno4")
	ColunaAdd(menu, "MargemAno4")
	ColunaAdd(menu, "InvFinalAno4")
	ColunaAdd(menu, "VlInvAno4")
	ColunaAdd(menu, "DiferencasAno4")
	// Ano 5
	ColunaAdd(menu, "EntradasAno5")
	ColunaAdd(menu, "VlTotalEntradasAno5")
	ColunaAdd(menu, "VlUnitEntAno5")
	ColunaAdd(menu, "SaidasAno5")
	ColunaAdd(menu, "VlTotalSaidasAno5")
	ColunaAdd(menu, "VlUnitSaidaAno5")
	ColunaAdd(menu, "MargemAno5")
	ColunaAdd(menu, "InvFinalAno5")
	ColunaAdd(menu, "VlInvAno5")
	ColunaAdd(menu, "DiferencasAno5")
	// Ano 6
	ColunaAdd(menu, "EntradasAno6")
	ColunaAdd(menu, "VlTotalEntradasAno6")
	ColunaAdd(menu, "VlUnitEntAno6")
	ColunaAdd(menu, "SaidasAno6")
	ColunaAdd(menu, "VlTotalSaidasAno6")
	ColunaAdd(menu, "VlUnitSaidaAno6")
	ColunaAdd(menu, "MargemAno6")
	ColunaAdd(menu, "InvFinalAno6")
	ColunaAdd(menu, "VlInvAno6")
	ColunaAdd(menu, "DiferencasAno6")
}

/*
func CriarH010InvInicial(ano int, db gorm.DB) {
	ano = ano - 1
	anoString := strconv.Itoa(ano)
	var inv []Models.Inventario
	db.Find(&inv)
	f, err := os.Create("SpedInvInicial.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, vInv := range inv {
		if vInv.SugInvInicial > 0 {
			r0200 := Bloco0.Reg0200{
				Reg:       "0200",
				CodItem:   vInv.Codigo,
				DescrItem: vInv.Descricao,
				UnidInv:   vInv.UnidInv,
				TipoItem:  vInv.Tipo,
			}
			aliqicms := tools.FloatToStringSped(r0200.AliqIcms)
			linha := "|" + r0200.Reg + "|" + r0200.CodItem + "|" + r0200.DescrItem + "|" +
				r0200.CodBarra + "|" + r0200.CodAntItem + "|" + r0200.UnidInv + "|" + r0200.TipoItem +
				"|" + r0200.CodNcm + "|" + r0200.ExIpi + "|" + r0200.CodGen + "|" + r0200.CodLst +
				"|" + aliqicms + "|\r\n"
			f.WriteString(linha)
			f.Sync()
		}
	}
	linha := "|H005|3112" + anoString + "|1726778,31|01|\r\n"
	f.WriteString(linha)
	f.Sync()

	for _, vInv2 := range inv {
		if vInv2.SugInvInicial > 0 {
			sugVlUnit := vInv2.SugVlInvInicial / vInv2.SugInvInicial
			h010 := BlocoH.RegH010{
				Reg:     "H010",
				CodItem: vInv2.Codigo,
				Unid:    vInv2.UnidInv,
				Qtd:     vInv2.SugInvInicial,
				VlUnit:  sugVlUnit,
				VlItem:  vInv2.SugVlInvInicial,
				IndProp: "0",
			}
			linha := "|" + h010.Reg + "|" + h010.CodItem + "|" + h010.Unid + "|" +
				tools.FloatToStringSped(h010.Qtd) + "|" + tools.FloatToStringSped(h010.VlUnit) +
				"|" + tools.FloatToStringSped(h010.VlItem) + "|" + h010.IndProp + "|" + h010.CodPart +
				"|" + h010.CodCta + "|" + tools.FloatToStringSped(h010.VlItemIr) + "|\r\n"
			f.WriteString(linha)
			f.Sync()
		}

	}

	w := bufio.NewWriter(f)
	w.Flush()
}

func CriarH010InvFinal(ano int, db gorm.DB) {
	anoString := strconv.Itoa(ano)
	var inv []Models.Inventario
	db.Find(&inv)
	f, err := os.Create("SpedInvFinal.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, vInv := range inv {
		if vInv.SugInvFinal > 0 {
			r0200 := Bloco0.Reg0200{
				Reg:       "0200",
				CodItem:   vInv.Codigo,
				DescrItem: vInv.Descricao,
				UnidInv:   vInv.UnidInv,
				TipoItem:  vInv.Tipo,
			}
			aliqicms := tools.FloatToStringSped(r0200.AliqIcms)
			linha := "|" + r0200.Reg + "|" + r0200.CodItem + "|" + r0200.DescrItem + "|" +
				r0200.CodBarra + "|" + r0200.CodAntItem + "|" + r0200.UnidInv + "|" + r0200.TipoItem +
				"|" + r0200.CodNcm + "|" + r0200.ExIpi + "|" + r0200.CodGen + "|" + r0200.CodLst +
				"|" + aliqicms + "|\r\n"
			f.WriteString(linha)
			f.Sync()
		}
	}
	linha := "|H005|3112" + anoString + "|1726778,31|01|\r\n"
	f.WriteString(linha)
	f.Sync()

	for _, vInv2 := range inv {
		if vInv2.SugInvFinal > 0 {
			sugVlUnit := vInv2.SugVlInvFinal / vInv2.SugInvFinal
			h010 := BlocoH.RegH010{
				Reg:     "H010",
				CodItem: vInv2.Codigo,
				Unid:    vInv2.UnidInv,
				Qtd:     vInv2.SugInvFinal,
				VlUnit:  sugVlUnit,
				VlItem:  vInv2.SugVlInvFinal,
				IndProp: "0",
			}
			linha := "|" + h010.Reg + "|" + h010.CodItem + "|" + h010.Unid + "|" +
				tools.FloatToStringSped(h010.Qtd) + "|" + tools.FloatToStringSped(h010.VlUnit) +
				"|" + tools.FloatToStringSped(h010.VlItem) + "|" + h010.IndProp + "|" + h010.CodPart +
				"|" + h010.CodCta + "|" + tools.FloatToStringSped(h010.VlItemIr) + "|\r\n"
			f.WriteString(linha)
			f.Sync()
		}

	}

	w := bufio.NewWriter(f)
	w.Flush()
}

*/
