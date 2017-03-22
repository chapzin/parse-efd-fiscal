package tools

import "github.com/sirupsen/logrus"

func CheckErr(err error) {
	if err != nil {
		logrus.Warn(err)
		// Habilita abaixo para pausar no erro apresentado
		//bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
