package database

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

func NewRedis(redisURL string) (*redis.Client, error) {

	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	slog.Info("redis_connected")

	return client, nil
}
