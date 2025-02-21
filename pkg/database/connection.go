package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Config struct {
	Dialect  string
	ConnStr  string
	LogMode  bool
}

func NewConnection(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(cfg.Dialect, cfg.ConnStr)
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir conexão com banco de dados: %v", err)
	}
	
	db.LogMode(cfg.LogMode)
	
	return db, nil
} 