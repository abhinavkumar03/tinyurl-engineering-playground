package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL     string
	ServerPort  string
	FrontendURL string

	PostgresHost string
	PostgresPort string
	PostgresDB   string
	PostgresUser string
	PostgresPass string

	RedisURL string
}

func Load() *Config {

	_ = godotenv.Load()

	return &Config{
		BaseURL:     os.Getenv("BASE_URL"),
		ServerPort:  os.Getenv("SERVER_PORT"),
		FrontendURL: os.Getenv("FRONTEND_URL"),

		PostgresHost: os.Getenv("POSTGRES_HOST"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
		PostgresDB:   os.Getenv("POSTGRES_DB"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPass: os.Getenv("POSTGRES_PASSWORD"),

		RedisURL: os.Getenv("REDIS_URL"),
	}
}
