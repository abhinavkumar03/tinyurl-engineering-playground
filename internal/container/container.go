package container

import (
	"fmt"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/config"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/database"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/handler"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/repository"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Container struct {
	Config *config.Config

	Postgres *pgxpool.Pool
	Redis    *redis.Client

	URLHandler    *handler.URLHandler
	HealthHandler *handler.HealthHandler

	AnalyticsService service.AnalyticsService
	AnalyticsHandler *handler.AnalyticsHandler
}

func Build() (*Container, error) {

	cfg := config.Load()

	pgConn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	pg, err := database.NewPostgres(
		pgConn,
	)
	if err != nil {
		return nil, err
	}

	redisClient, err := database.NewRedis(
		cfg.RedisURL,
	)
	if err != nil {
		return nil, err
	}

	urlRepository := repository.NewURLRepository(
		pg,
	)

	analyticsRepository := repository.NewAnalyticsRepository(
		pg,
	)

	analyticsService := service.NewAnalyticsService(
		analyticsRepository,
		urlRepository,
	)

	urlService := service.NewURLService(
		urlRepository,
		redisClient,
		cfg.BaseURL,
	)

	urlHandler := handler.NewURLHandler(
		urlService,
		analyticsService,
		cfg.BaseURL,
	)

	healthService := service.NewHealthService(
		pg,
		redisClient,
		"1.0.0",
	)

	healthHandler := handler.NewHealthHandler(
		healthService,
	)

	analyticsHandler := handler.NewAnalyticsHandler(
		analyticsService,
	)

	return &Container{
		Config:           cfg,
		Postgres:         pg,
		Redis:            redisClient,
		URLHandler:       urlHandler,
		HealthHandler:    healthHandler,
		AnalyticsService: analyticsService,
		AnalyticsHandler: analyticsHandler,
	}, nil
}
