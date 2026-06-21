package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type HealthService struct {
	postgres  *pgxpool.Pool
	redis     *redis.Client
	startedAt time.Time
	version   string
}

func NewHealthService(
	postgres *pgxpool.Pool,
	redisClient *redis.Client,
	version string,
) *HealthService {

	return &HealthService{
		postgres:  postgres,
		redis:     redisClient,
		startedAt: time.Now(),
		version:   version,
	}
}

func (s *HealthService) Health(
	ctx context.Context,
) (string, string, string) {

	postgresStatus := "up"
	redisStatus := "up"

	if err := s.postgres.Ping(ctx); err != nil {
		postgresStatus = "down"
	}

	if err := s.redis.Ping(ctx).Err(); err != nil {
		redisStatus = "down"
	}

	status := "healthy"

	if postgresStatus == "down" || redisStatus == "down" {
		status = "unhealthy"
	}

	return status, postgresStatus, redisStatus
}

func (s *HealthService) Version() string {
	return s.version
}

func (s *HealthService) UptimeSeconds() int64 {
	return int64(time.Since(s.startedAt).Seconds())
}
