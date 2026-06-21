package handler

import (
	"net/http"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/dto"
	apperrors "github.com/abhinavkumar03/tinyurl-engineering-playground/internal/errors"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/service"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service          service.URLService
	analyticsService service.AnalyticsService
	baseURL          string
}

func NewURLHandler(
	urlService service.URLService,
	analyticsService service.AnalyticsService,
	baseURL string,
) *URLHandler {

	return &URLHandler{
		service:          urlService,
		analyticsService: analyticsService,
		baseURL:          baseURL,
	}
}

func (h *URLHandler) Create(c *gin.Context) {

	var request dto.CreateURLRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		apperrors.BadRequest(
			c,
			err.Error(),
		)

		return
	}

	url, err := h.service.Create(
		c.Request.Context(),
		request.URL,
	)

	if err != nil {

		apperrors.Internal(
			c,
			"failed to create short url",
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

func (h *URLHandler) Redirect(
	c *gin.Context,
) {

	shortCode := c.Param(
		"shortCode",
	)

	resolvedURL, err := h.service.Resolve(
		c.Request.Context(),
		shortCode,
	)

	if err != nil {

		apperrors.NotFound(
			c,
			"url not found",
		)

		return
	}

	_ = h.analyticsService.Track(
		c.Request.Context(),
		model.RedirectEvent{
			URLID: resolvedURL.ID,

			ShortCode: resolvedURL.ShortCode,

			IPAddress: c.ClientIP(),

			UserAgent: c.Request.UserAgent(),

			Referrer: c.Request.Referer(),
		},
	)

	c.Redirect(
		http.StatusMovedPermanently,
		resolvedURL.OriginalURL,
	)
}

func (h *URLHandler) Get(c *gin.Context) {

	shortCode := c.Param("shortCode")

	url, err := h.service.Get(
		c.Request.Context(),
		shortCode,
	)

	if err != nil {

		apperrors.NotFound(
			c,
			"url not found",
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
