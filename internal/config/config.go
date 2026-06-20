package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL    string
	ServerPort string

	PostgresHost string
	PostgresPort string
	PostgresDB   string
	PostgresUser string
	PostgresPass string

	RedisHost string
	RedisPort string
}

func Load() *Config {

	_ = godotenv.Load()

	return &Config{
		BaseURL:    os.Getenv("BASE_URL"),
		ServerPort: os.Getenv("SERVER_PORT"),

		PostgresHost: os.Getenv("POSTGRES_HOST"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
		PostgresDB:   os.Getenv("POSTGRES_DB"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPass: os.Getenv("POSTGRES_PASSWORD"),

		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
	}
}
