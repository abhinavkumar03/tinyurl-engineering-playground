package errors

import (
	"net/http"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/dto"
	"github.com/gin-gonic/gin"
)

func Write(
	c *gin.Context,
	status int,
	code string,
	message string,
) {

	c.JSON(
		status,
		dto.ErrorResponse{
			Code:    code,
			Message: message,
		},
	)
}

func BadRequest(
	c *gin.Context,
	message string,
) {

	Write(
		c,
		http.StatusBadRequest,
		"BAD_REQUEST",
		message,
	)
}

func Internal(
	c *gin.Context,
	message string,
) {

	Write(
		c,
		http.StatusInternalServerError,
		"INTERNAL_SERVER_ERROR",
		message,
	)
}

func NotFound(
	c *gin.Context,
	message string,
) {

	Write(
		c,
		http.StatusNotFound,
		"NOT_FOUND",
		message,
	)
}
