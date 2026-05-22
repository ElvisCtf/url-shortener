package util

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Addr      string
	BaseURL   string
	GinMode   string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
	TZ        string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load() // Loads .env file if present
	return &Config{
		Addr:      getEnv("ADDR", ":8082"),
		BaseURL:   getEnv("BASE_URL", "http://localhost:8082"),
		GinMode:   getEnv("GIN_MODE", "debug"),
		DBHost:    getEnv("DB_HOST", "postgres16"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "postgres16"),
		DBPass:    getEnv("DB_PASSWORD", "postgres16"),
		DBName:    getEnv("DB_NAME", "urlshortener"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
		TZ:        getEnv("TZ", "UTC"),
	}, nil
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
