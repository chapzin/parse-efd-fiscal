package read

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	id    int
	maxid = 98
	mu    sync.RWMutex // Mutex para proteger variáveis globais
)

// Estrutura para controlar o processamento concorrente
type processControl struct {
	wg       sync.WaitGroup
	errChan  chan error
	doneChan chan struct{}
	db       *gorm.DB
	digito   string
}

// newProcessControl cria uma nova instância de controle de processamento
func newProcessControl(db *gorm.DB, digito string) *processControl {
	return &processControl{
		errChan:  make(chan error, 100),
		doneChan: make(chan struct{}),
		db:       db,
		digito:   digito,
	}
}

// wait implementa um mecanismo de espera com timeout
func wait() error {
	timeout := time.After(5 * time.Minute)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if id < maxid {
				return nil
			}
		case <-timeout:
			return fmt.Errorf("timeout esperando processamento")
		}
	}
}
