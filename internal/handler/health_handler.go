package handler

import (
	"net/http"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/dto"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/service"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	service *service.HealthService
}

func NewHealthHandler(
	service *service.HealthService,
) *HealthHandler {

	return &HealthHandler{
		service: service,
	}
}

func (h *HealthHandler) Health(
	c *gin.Context,
) {

	status, postgres, redis := h.service.Health(
		c.Request.Context(),
	)

	httpStatus := http.StatusOK

	if status == "unhealthy" {
		httpStatus = http.StatusServiceUnavailable
	}

	c.JSON(
		httpStatus,
		dto.HealthResponse{
			Status:    status,
			Postgres:  postgres,
			Redis:     redis,
			Version:   h.service.Version(),
			UptimeSec: h.service.UptimeSeconds(),
		},
	)
}

func (h *HealthHandler) Readiness(
	c *gin.Context,
) {

	status, _, _ := h.service.Health(
		c.Request.Context(),
	)

	if status != "healthy" {

		c.JSON(
			http.StatusServiceUnavailable,
			gin.H{
				"ready": false,
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"ready": true,
		},
	)
}

func (h *HealthHandler) Liveness(
	c *gin.Context,
) {

	c.JSON(
		http.StatusOK,
		gin.H{
			"alive": true,
		},
	)
}
