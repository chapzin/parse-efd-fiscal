package tools

import "github.com/sirupsen/logrus"

// Funcao criada para fazer o analise de erro debugando
func CheckErr(err error) {
	if err != nil {
		logrus.Warn(err)
		// Habilita abaixo para pausar no erro apresentado
		//bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
