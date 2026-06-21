package handler

import (
	"net/http"
	"strconv"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/dto"
	apperrors "github.com/abhinavkumar03/tinyurl-engineering-playground/internal/errors"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/service"
	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	service service.AnalyticsService
}

func NewAnalyticsHandler(
	service service.AnalyticsService,
) *AnalyticsHandler {

	return &AnalyticsHandler{
		service: service,
	}
}

func (h *AnalyticsHandler) URLAnalytics(
	c *gin.Context,
) {

	shortCode := c.Param(
		"shortCode",
	)

	totalClicks, err := h.service.GetTotalClicks(
		c.Request.Context(),
		shortCode,
	)
	if err != nil {

		apperrors.Internal(
			c,
			err.Error(),
		)

		return
	}

	uniqueVisitors, err := h.service.GetUniqueVisitors(
		c.Request.Context(),
		shortCode,
	)
	if err != nil {

		apperrors.Internal(
			c,
			err.Error(),
		)

		return
	}

	c.JSON(
		http.StatusOK,
		dto.URLAnalyticsResponse{
			ShortCode:      shortCode,
			TotalClicks:    totalClicks,
			UniqueVisitors: uniqueVisitors,
		},
	)
}

func (h *AnalyticsHandler) TopURLs(
	c *gin.Context,
) {

	limit := 10

	if value := c.Query("limit"); value != "" {

		parsed, err := strconv.Atoi(
			value,
		)
		if err == nil && parsed > 0 {
			limit = parsed
		}
	}

	results, err := h.service.GetTopURLs(
		c.Request.Context(),
		limit,
	)
	if err != nil {

		apperrors.Internal(
			c,
			err.Error(),
		)

		return
	}

	response := make(
		[]dto.TopURLResponse,
		0,
		len(results),
	)

	for _, item := range results {

		response = append(
			response,
			dto.TopURLResponse{
				ShortCode:   item.ShortCode,
				OriginalURL: item.OriginalURL,
				Clicks:      item.Clicks,
			},
		)
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}

func (h *AnalyticsHandler) DailyClicks(
	c *gin.Context,
) {

	shortCode := c.Param(
		"shortCode",
	)

	results, err := h.service.GetClicksByDay(
		c.Request.Context(),
		shortCode,
	)
	if err != nil {

		apperrors.Internal(
			c,
			err.Error(),
		)

		return
	}

	response := make(
		[]dto.DailyClicksResponse,
		0,
		len(results),
	)

	for _, item := range results {

		response = append(
			response,
			dto.DailyClicksResponse{
				Date:   item.Date.Format("2006-01-02"),
				Clicks: item.Clicks,
			},
		)
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}

func (h *AnalyticsHandler) Dashboard(
	c *gin.Context,
) {

	topURLs, err := h.service.GetTopURLs(
		c.Request.Context(),
		10,
	)
	if err != nil {

		apperrors.Internal(
			c,
			err.Error(),
		)

		return
	}

	response := dto.DashboardResponse{
		TopURLs: make(
			[]dto.TopURLResponse,
			0,
			len(topURLs),
		),
	}

	for _, item := range topURLs {

		response.TopURLs = append(
			response.TopURLs,
			dto.TopURLResponse{
				ShortCode:   item.ShortCode,
				OriginalURL: item.OriginalURL,
				Clicks:      item.Clicks,
			},
		)
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}
