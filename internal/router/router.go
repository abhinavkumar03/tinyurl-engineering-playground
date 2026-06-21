package router

import (
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/handler"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(
	urlHandler *handler.URLHandler,
	healthHandler *handler.HealthHandler,
	frontendURLs []string,
) *gin.Engine {

	r := gin.New()

	r.Use(
		middleware.RequestID(),
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.SecurityHeaders(),
	)

	r.Use(cors.New(cors.Config{
		AllowOrigins: frontendURLs,
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
	}))

	r.GET("/health", healthHandler.Health)
	r.GET("/ready", healthHandler.Readiness)
	r.GET("/live", healthHandler.Liveness)

	v1 := r.Group("/api/v1")

	{
		v1.POST(
			"/urls",
			urlHandler.Create,
		)

		v1.GET(
			"/urls/:shortCode",
			urlHandler.Get,
		)
	}

	r.GET(
		"/:shortCode",
		urlHandler.Redirect,
	)

	return r
}
