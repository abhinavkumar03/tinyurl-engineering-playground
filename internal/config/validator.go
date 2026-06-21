package config

import (
	"log"
	"os"
)

func Validate(cfg *Config) {

	required := map[string]string{
		"POSTGRES_HOST":     cfg.PostgresHost,
		"POSTGRES_DB":       cfg.PostgresDB,
		"POSTGRES_USER":     cfg.PostgresUser,
		"POSTGRES_PASSWORD": cfg.PostgresPass,
		"REDIS_URL":         cfg.RedisURL,
	}

	for key, value := range required {
		if value == "" {
			log.Printf(
				"missing required environment variable: %s",
				key,
			)
			os.Exit(1)
		}
	}
}
