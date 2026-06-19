package main

import (
	"fmt"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/config"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/database"
)

func main() {

	cfg := config.Load()

	pgConn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	pg, err := database.NewPostgres(pgConn)
	if err != nil {
		panic(err)
	}

	redisClient, err := database.NewRedis(
		cfg.RedisHost + ":" + cfg.RedisPort,
	)
	if err != nil {
		panic(err)
	}

	defer pg.Close()
	defer redisClient.Close()

	fmt.Println("application started")
}
