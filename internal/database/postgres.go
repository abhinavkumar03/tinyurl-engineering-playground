package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(conn string) (*pgxpool.Pool, error) {

	pool, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	slog.Info("postgres_connected")

	return pool, nil
}
