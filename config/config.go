package config

import (
	"fmt"
	"time"

	"github.com/chapzin/parse-efd-fiscal/pkg/env"
)

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	DBName   string
	Charset  string
	Host     string
	Port     string
}

type WorkerConfig struct {
	MaxWorkers  int
	TaskTimeout time.Duration
}

type Config struct {
	DB        DBConfig
	Worker    WorkerConfig
	DigitCode string
	SpedsPath string
}

// Valores padrão
const (
	DefaultMaxWorkers  = 3
	DefaultTaskTimeout = 30 * time.Minute
)

func LoadConfig() (*Config, error) {
	if err := env.LoadEnv(); err != nil {
		return nil, err
	}

	return &Config{
		DB: DBConfig{
			Dialect:  env.GetEnvOrDefault("DB_DIALECT", "mysql"),
			Username: env.GetEnvOrDefault("DB_USERNAME", "root"),
			Password: env.GetEnvOrDefault("DB_PASSWORD", ""),
			DBName:   env.GetEnvOrDefault("DB_NAME", "auditoria2"),
			Charset:  env.GetEnvOrDefault("DB_CHARSET", "utf8"),
			Host:     env.GetEnvOrDefault("DB_HOST", "localhost"),
			Port:     env.GetEnvOrDefault("DB_PORT", "3306"),
		},
		Worker: WorkerConfig{
			MaxWorkers:  env.GetEnvIntOrDefault("WORKER_MAX_WORKERS", DefaultMaxWorkers),
			TaskTimeout: env.GetEnvDurationOrDefault("WORKER_TASK_TIMEOUT", DefaultTaskTimeout),
		},
		DigitCode: env.GetEnvOrDefault("DIGIT_CODE", "10"),
		SpedsPath: env.GetEnvOrDefault("SPEDS_PATH", "./speds"),
	}, nil
}

func (c *Config) GetMySQLConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		c.DB.Username,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
		c.DB.DBName,
		c.DB.Charset,
	)
}
