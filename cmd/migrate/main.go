package main

import (
	"fmt"
	"os"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	cfg := config.Load()

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=require",
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	m, err := migrate.New(
		"file://internal/migrations",
		databaseURL,
	)
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		panic("missing command: up | down | force")
	}

	command := os.Args[1]

	switch command {

	case "up":

		if err := m.Up(); err != nil &&
			err.Error() != "no change" {
			panic(err)
		}

	case "down":

		if err := m.Down(); err != nil {
			panic(err)
		}

	case "force":

		if len(os.Args) < 3 {
			panic("missing version")
		}

		var version int

		fmt.Sscanf(
			os.Args[2],
			"%d",
			&version,
		)

		if err := m.Force(version); err != nil {
			panic(err)
		}

	default:

		panic("unknown command")
	}
}
