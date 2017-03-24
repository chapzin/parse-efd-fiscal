package Controllers

import (
	"github.com/chapzin/GoInventario/Models"
	"github.com/chapzin/GoInventario/Tools"
	"github.com/chapzin/parse-efd-fiscal/Models/Bloco0"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoC"
	"github.com/chapzin/parse-efd-fiscal/Models/BlocoH"
	"github.com/chapzin/parse-efd-fiscal/Models/NotaFiscal"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"strconv"
	"sync"
	"time"
)

func ProcessarFatorConversao(db gorm.DB, wg *sync.WaitGroup) {
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

func DeletarItensNotasCanceladas(db gorm.DB, dtIni string, dtFin string, wg *sync.WaitGroup) {
	color.Green("Começo Deleta itens notas canceladas %s", time.Now())
	var c100 []BlocoC.RegC100
	db.Where("cod_sit <> ? and dt_ini >= ? and dt_ini <= ? ", "00", dtIni, dtFin).Find(&c100)
	for _, v := range c100 {
		//fmt.Println(v.NumDoc)
		var nota []NotaFiscal.NotaFiscal
		db.Where("ch_n_fe = ?", v.ChvNfe).Find(&nota)
		for _, v2 := range nota {
			db.Where("nota_fiscal_id =?", v2.ID).Delete(NotaFiscal.Item{})
		}
	}
	db.Exec("DELETE FROM items WHERE deleted_at is not null")
	color.Green("Fim deleta itens notas canceladas %s", time.Now())
	wg.Done()
}

func PopularReg0200(db gorm.DB, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	color.Green("Comeco popula reg0200 %s", time.Now())
	var reg0200 []Bloco0.Reg0200
	db.Where("tipo_item=00").Select("distinct cod_item,descr_item,tipo_item,unid_inv").Find(&reg0200)
	for _, v := range reg0200 {
		var inventario Models.Inventario
		db.Where("codigo=?", v.CodItem).First(&inventario)
		if inventario.Codigo == "" {
			inv2 := Models.Inventario{
				Codigo:    v.CodItem,
				Descricao: v.DescrItem,
				Tipo:      v.TipoItem,
				UnidInv:   v.UnidInv,
			}
			db.NewRecord(inv2)
			db.Create(&inv2)
		}
	}
	wg.Done()
	color.Green("Fim popula reg0200 %s", time.Now())
}

func PopularItensXmls(db gorm.DB, wg *sync.WaitGroup) {
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

func PopularInventario(InicialFinal string, ano int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	color.Green("Começo popula Inventario %s", time.Now())
	db, err := gorm.Open("mysql", "root:123@/auditoria2?charset=utf8&parseTime=true")
	Tools.CheckErr(err)
	var regH010 []BlocoH.RegH010
	if InicialFinal == "final" {
		ano = ano + 1
	}
	anostring := strconv.Itoa(ano)
	db.Where("dt_ini= ?", anostring+"-02-01").Find(&regH010)
	for _, v3 := range regH010 {
		inventario := Models.Inventario{}
		if InicialFinal == "inicial" {
			inventario.InvInicial = v3.Qtd
			inventario.VlInvIni = v3.VlUnit
			inventario.Ano = ano
			db.Table("inventarios").Where("codigo = ?", v3.CodItem).Update(&inventario)
		}
		if InicialFinal == "final" {
			inventario.InvFinal = v3.Qtd
			inventario.VlInvFin = v3.VlUnit
			inventario.Ano = ano - 1
			db.Table("inventarios").Where("codigo = ?", v3.CodItem).Update(&inventario)
		}
	}
	color.Green("Fim popula inventario %s", time.Now())
	wg.Done()

}

func PopularEntradas(ano string, wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	color.Green("Comeco popula entradas %s", time.Now())
	db, err := gorm.Open("mysql", "root:123@/auditoria2?charset=utf8&parseTime=true")
	Tools.CheckErr(err)
	dtIni := ano + "-01-01"
	dtFin := ano + "-12-31"
	// Carregando inventario na memoria
	var inv []Models.Inventario
	db.Find(&inv)
	// Carregando RegC170 na memoria do periodo
	var c170 []BlocoC.RegC170
	db.Where("entrada_saida = ? and dt_ini >= ? and dt_fin <= ? ", "0", dtIni, dtFin).Find(&c170)
	var itens []NotaFiscal.Item
	db.Where("cfop < 3999 and dt_emit >= ? and dt_emit <= ?", dtIni, dtFin).Find(&itens)

	// Listando itens que estão no inventário
	for _, vInv := range inv {
		var qtd_tot = 0.0
		var vl_tot = 0.0
		// Compara o codigo do c170 com o do inventario e adiciona a qtd
		for _, vC170 := range c170 {
			if vC170.CodItem == vInv.Codigo {
				qtd_tot = qtd_tot + vC170.Qtd
				vl_tot = vl_tot + vC170.VlItem
			}
		}
		// Listando itens das notas fiscais que o cfop seja menor que 3999 e adicionando na entrada
		for _, vItens := range itens {
			if vItens.Codigo == vInv.Codigo {
				qtd_tot = qtd_tot + vItens.Qtd
				vl_tot = vl_tot + vItens.VTotal
			}
		}
		// inserindo os valores finais das somas a entrada
		inv2 := Models.Inventario{}
		inv2.Entradas = qtd_tot
		inv2.VlTotalEntradas = vl_tot
		db.Table("inventarios").Where("codigo = ?", vInv.Codigo).Update(&inv2)
	}
	color.Green("Fim popula entradas %s", time.Now())
	wg.Done()
}

func PopularSaidas(ano string, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	color.Green("Comeco popula saidas %s", time.Now())
	db, err := gorm.Open("mysql", "root:123@/auditoria2?charset=utf8&parseTime=true")
	Tools.CheckErr(err)
	dtIni := ano + "-01-01"
	dtFin := ano + "-12-31"
	var inv []Models.Inventario
	db.Find(&inv)
	var itens []NotaFiscal.Item
	db.Where("cfop > 3999 and cfop <> 5929 and cfop <> 6929 and dt_emit >= ? and dt_emit <= ?", dtIni, dtFin).Find(&itens)
	var c425 []BlocoC.RegC425
	db.Where("dt_ini >= ? and dt_ini <= ?", dtIni, dtFin).Find(&c425)
	for _, vInv := range inv {
		var qtd_saida = 0.0
		var vl_tot_saida = 0.0
		for _, vItens := range itens {
			if vItens.Codigo == vInv.Codigo {
				qtd_saida = qtd_saida + vItens.Qtd
				vl_tot_saida = vl_tot_saida + vItens.VTotal
			}
		}
		for _, vC425 := range c425 {
			if vC425.CodItem == vInv.Codigo {
				qtd_saida = qtd_saida + vC425.Qtd
				vl_tot_saida = vl_tot_saida + vC425.VlItem
			}
		}
		inv3 := Models.Inventario{}
		inv3.Saidas = qtd_saida
		inv3.VlTotalSaidas = vl_tot_saida
		db.Table("inventarios").Where("codigo = ?", vInv.Codigo).Update(&inv3)
	}
	color.Green("Fim popula saidas %s", time.Now())
	wg.Done()

}

func ProcessarDiferencas(db gorm.DB) {
	db.Exec("Delete from inventarios where inv_inicial=0 and entradas=0 and vl_total_entradas=0 and saidas=0 and vl_total_saidas=0 and inv_final=0")
	var inv []Models.Inventario
	var reg0200 []Bloco0.Reg0200
	db.Select("distinct cod_item,descr_item,tipo_item,unid_inv").Find(&reg0200)
	db.Find(&inv)
	for _, vInv := range inv {
		inv3 := Models.Inventario{}
		// Calculando as diferencas
		diferencas := (vInv.InvInicial + vInv.Entradas) - (vInv.Saidas + vInv.InvFinal)

		// Calculando o valor unitário de entrada
		if vInv.VlTotalEntradas > 0 && vInv.Entradas > 0 {
			vlUnitEntrada := vInv.VlTotalEntradas / vInv.Entradas
			inv3.VlUnitEnt = vlUnitEntrada
		} else if vInv.VlTotalEntradas == 0 && vInv.Entradas == 0 && vInv.VlInvIni > 0 {
			vlUnitEntrada := vInv.VlInvIni
			inv3.VlUnitEnt = vlUnitEntrada
		} else {
			inv3.VlUnitEnt = 0
		}

		// Calculando o valor unitário de saida
		if vInv.VlTotalSaidas > 0 && vInv.Saidas > 0 {
			vlUnitSaida := vInv.VlTotalSaidas / vInv.Saidas
			inv3.VlUnitSai = vlUnitSaida
		} else if vInv.VlTotalSaidas == 0 && vInv.Saidas == 0 && vInv.VlInvIni > 0 {
			vlUnitSaida := vInv.VlInvIni
			inv3.VlUnitSai = vlUnitSaida
		} else {
			inv3.VlUnitSai = 0
		}

		// Criando Sugestao de novo inventário
		if diferencas >= 0 {
			// Novo inventario final somando diferencas
			nvInvFin := diferencas + vInv.InvFinal
			inv3.EstoqueFin = nvInvFin
		} else {
			inv3.EstoqueFin = vInv.InvFinal
		}
		if diferencas < 0 {
			// Caso negativo adiciona ao inventario inicial
			nvInvIni := (diferencas * -1) + vInv.InvInicial
			inv3.EstoqueIni = nvInvIni
		} else {
			// Caso nao seja negativo mantenha o inventario anterior
			inv3.EstoqueIni = vInv.InvInicial
		}
		// Adicionando Tipo e unidade de medida no inventario
		for _, v0200 := range reg0200 {
			if v0200.CodItem == vInv.Codigo {
				inv3.Tipo = v0200.TipoItem
				inv3.UnidInv = v0200.UnidInv
			}
		}
		inv3.Diferencas = diferencas
		db.Table("inventarios").Where("codigo = ?", vInv.Codigo).Update(&inv3)
	}
	// Deleta tudo tipo de inventario que nao seja material de revenda
	db.Exec("Delete from inventarios where tipo <> '00'")
}
