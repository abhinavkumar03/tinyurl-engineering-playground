package database

import (
	"context"
	"fmt"

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

	fmt.Println("postgres connected")

	return pool, nil
}