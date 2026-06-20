package handler

import (
	"net/http"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/dto"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/service"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service service.URLService
	baseURL string
}

func NewURLHandler(
	service service.URLService,
	baseURL string,
) *URLHandler {

	return &URLHandler{
		service: service,
		baseURL: baseURL,
	}
}

func (h *URLHandler) Create(c *gin.Context) {

	var request dto.CreateURLRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Error: err.Error(),
			},
		)

		return
	}

	url, err := h.service.Create(
		c.Request.Context(),
		request.URL,
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			dto.ErrorResponse{
				Error: err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		dto.CreateURLResponse{
			ID:          url.ID,
			OriginalURL: url.OriginalURL,
			ShortCode:   url.ShortCode,
			ShortURL:    h.baseURL + "/" + url.ShortCode,
		},
	)
}

func (h *URLHandler) Redirect(c *gin.Context) {

	shortCode := c.Param("shortCode")

	url, err := h.service.Resolve(
		c.Request.Context(),
		shortCode,
	)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			dto.ErrorResponse{
				Error: "url not found",
			},
		)
		return
	}

	c.Redirect(
		http.StatusMovedPermanently,
		url,
	)
}

func (h *URLHandler) Get(c *gin.Context) {

	shortCode := c.Param("shortCode")

	url, err := h.service.Get(
		c.Request.Context(),
		shortCode,
	)

	if err != nil {

		c.JSON(
			http.StatusNotFound,
			dto.ErrorResponse{
				Error: "url not found",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		dto.URLResponse{
			ID:          url.ID,
			OriginalURL: url.OriginalURL,
			ShortCode:   url.ShortCode,
			ClickCount:  url.ClickCount,
		},
	)
}
