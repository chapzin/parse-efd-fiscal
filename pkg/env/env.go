package env

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// LoadEnv carrega as variáveis de ambiente do arquivo .env
func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("erro ao carregar arquivo .env: %v", err)
	}
	return nil
}

// GetEnvOrDefault retorna o valor da variável de ambiente ou o valor padrão
func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetEnvIntOrDefault retorna o valor da variável de ambiente como inteiro ou o valor padrão
func GetEnvIntOrDefault(key string, defaultValue int) int {
	strValue := os.Getenv(key)
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return defaultValue
}

// GetEnvDurationOrDefault retorna o valor da variável de ambiente como Duration ou o valor padrão
func GetEnvDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	strValue := os.Getenv(key)
	if value, err := time.ParseDuration(strValue); err == nil {
		return value
	}
	return defaultValue
} 