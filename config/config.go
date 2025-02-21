package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chapzin/parse-efd-fiscal/pkg/env"
	conf "github.com/robfig/config"
)

var (
	c            *conf.Config
	Propriedades ConfigInterface
)

type ConfigInterface interface {
	ObterTexto(chave string) (string, error)
	GetConfiguracoes()
	ObterWorkerConfig() WorkerConfig
}

type Configurador struct {
	config ConfigInterface
}

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

// ObterTexto é um método responsavel por obter valor de texto do arquivo de propriedades do sistema
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

// ObterWorkerConfig retorna a configuração do worker com valores padrão se não encontrados
func (cfg Configurador) ObterWorkerConfig() WorkerConfig {
	maxWorkersStr, err := cfg.ObterTexto("worker.max_workers")
	maxWorkers := DefaultMaxWorkers
	if err == nil {
		if n, err := strconv.Atoi(maxWorkersStr); err == nil && n > 0 {
			maxWorkers = n
		}
	}

	timeoutStr, err := cfg.ObterTexto("worker.task_timeout")
	timeout := DefaultTaskTimeout
	if err == nil {
		if d, err := time.ParseDuration(timeoutStr); err == nil && d > 0 {
			timeout = d
		}
	}

	return WorkerConfig{
		MaxWorkers:  maxWorkers,
		TaskTimeout: timeout,
	}
}

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
			MaxWorkers:  env.GetEnvIntOrDefault("WORKER_MAX_WORKERS", 3),
			TaskTimeout: env.GetEnvDurationOrDefault("WORKER_TASK_TIMEOUT", 30*time.Minute),
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
