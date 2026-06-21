package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		latency := time.Since(start)

		requestID, _ := c.Get(
			RequestIDKey,
		)

		slog.Info(
			"http_request",
			"request_id",
			requestID,
			"method",
			c.Request.Method,
			"path",
			c.Request.URL.Path,
			"status",
			c.Writer.Status(),
			"latency_ms",
			latency.Milliseconds(),
			"client_ip",
			c.ClientIP(),
		)
	}
}
