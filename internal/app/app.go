package app

import (
	"net/http"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/container"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/router"
)

func Run() error {

	c, err := container.Build()
	if err != nil {
		return err
	}

	defer c.Postgres.Close()
	defer c.Redis.Close()

	r := router.Setup(c.URLHandler, c.Config.FrontendURLs)

	server := &http.Server{
		Addr:    ":" + c.Config.ServerPort,
		Handler: r,
	}

	return server.ListenAndServe()
}
