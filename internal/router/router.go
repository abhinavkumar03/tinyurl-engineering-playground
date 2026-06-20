package router

import (
	"net/http"
	"time"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/dto"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(
	urlHandler *handler.URLHandler,
	frontendURL string,
) *gin.Engine {

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			frontendURL,
		},
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
		MaxAge:           12 * time.Hour,
	}))

	r.GET(
		"/health",
		func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				dto.HealthResponse{
					Status:   "healthy",
					Postgres: "up",
					Redis:    "up",
				},
			)
		},
	)

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
