package shared

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr               string
	Domain             string
	GinMode            string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPass             string
	DBName             string
	DBSSLMode          string
	TZ                 string
	AccessTokenExpire  int
	RefreshTokenExpire int
	JWTSecret          string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load() // Loads .env file if present
	return &Config{
		Addr:               getEnv("ADDR", ":8082"),
		Domain:             getEnv("DOMAIN", "localhost"),
		GinMode:            getEnv("GIN_MODE", "debug"),
		DBHost:             getEnv("DB_HOST", "postgres16"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres16"),
		DBPass:             getEnv("DB_PASSWORD", "postgres16"),
		DBName:             getEnv("DB_NAME", "urlshortener"),
		DBSSLMode:          getEnv("DB_SSLMODE", "disable"),
		TZ:                 getEnv("TZ", "UTC"),
		AccessTokenExpire:  getEnvInt("ACCESS_TOKEN_EXPIRE", 900),
		RefreshTokenExpire: getEnvInt("REFRESH_TOKEN_EXPIRE", 604800),
		JWTSecret:          getEnv("JWT_SECRET", "changeme"),
	}, nil
}

func getEnvInt(key string, defaultVal int) int {
	if v := os.Getenv(key); v != "" {
		var i int
		_, err := fmt.Sscanf(v, "%d", &i)
		if err == nil {
			return i
		}
	}
	return defaultVal
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

func IsSecure() bool {
	ginMode := os.Getenv("GIN_MODE")
	return ginMode == "release"
}
