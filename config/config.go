package config

import (
	"fmt"

	conf "github.com/robfig/config"
)

var (
	c            *conf.Config
	Propriedades ConfigInterface
)

type ConfigInterface interface {
	ObterTexto(chave string) (string, error)
	GetConfiguracoes()
}

type Configurador struct {
	config ConfigInterface
}

//ObterTexto é um método responsavel por obter valor de texto do arquivo de propriedades do sistema
func (cfg Configurador) ObterTexto(chave string) (string, error) {
	valor, err := c.String("DEFAULT", chave)
	if err != nil {
		fmt.Println("Falha ao obter o valor de texto", chave)

		return "", err
	}

	return valor, nil
}

func (cfg Configurador) GetConfiguracoes() {

	config, err := conf.ReadDefault("config/config.cfg")
	if err != nil {
		panic("Arquivo não encontrado")
	}

	c = config

}

func InicializaConfiguracoes(prop ConfigInterface) {
	Propriedades = prop
	Propriedades.GetConfiguracoes()

}
