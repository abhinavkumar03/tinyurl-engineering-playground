package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv       string
	BaseURL      string
	ServerPort   string
	FrontendURLs []string

	PostgresHost string
	PostgresPort string
	PostgresDB   string
	PostgresUser string
	PostgresPass string

	RedisURL string
}

func Load() *Config {

	_ = godotenv.Load()

	cfg := &Config{
		AppEnv:       getEnv("APP_ENV", "development"),
		BaseURL:      getEnv("BASE_URL", "http://localhost:8080"),
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		FrontendURLs: getEnvSlice("FRONTEND_URLS"),

		PostgresHost: getEnv("POSTGRES_HOST", ""),
		PostgresPort: getEnv("POSTGRES_PORT", "5432"),
		PostgresDB:   getEnv("POSTGRES_DB", ""),
		PostgresUser: getEnv("POSTGRES_USER", ""),
		PostgresPass: getEnv("POSTGRES_PASSWORD", ""),

		RedisURL: getEnv("REDIS_URL", ""),
	}

	Validate(cfg)

	return cfg
}

func getEnv(key string, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func getEnvSlice(key string) []string {

	value := os.Getenv(key)

	if value == "" {
		return []string{}
	}

	parts := strings.Split(value, ",")

	result := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if part != "" {
			result = append(result, part)
		}
	}

	return result
}
