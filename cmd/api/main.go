package main

import (
	"fmt"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/config"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/database"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/handler"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/repository"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/router"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/service"
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

	urlRepository := repository.NewURLRepository(pg)

	urlService := service.NewURLService(
		urlRepository,
		redisClient,
		cfg.BaseURL,
	)

	urlHandler := handler.NewURLHandler(
		urlService,
		cfg.BaseURL,
	)

	r := router.Setup(urlHandler)

	if err := r.Run(
		":" + cfg.ServerPort,
	); err != nil {
		panic(err)
	}
}
