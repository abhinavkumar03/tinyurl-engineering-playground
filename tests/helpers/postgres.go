package helpers

import (
	"context"

	postgrescontainer "github.com/testcontainers/testcontainers-go/modules/postgres"
)

func StartPostgres(
	ctx context.Context,
) (*postgrescontainer.PostgresContainer, error) {

	return postgrescontainer.Run(
		ctx,
		"postgres:16",
		postgrescontainer.WithDatabase(
			"tinyurl",
		),
		postgrescontainer.WithUsername(
			"postgres",
		),
		postgrescontainer.WithPassword(
			"postgres",
		),
	)
}
