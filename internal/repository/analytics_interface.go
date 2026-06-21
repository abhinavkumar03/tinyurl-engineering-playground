package repository

import (
	"context"
	"time"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
)

type TopURL struct {
	ShortCode   string
	OriginalURL string
	Clicks      int64
}

type DailyClicks struct {
	Date   time.Time
	Clicks int64
}

type AnalyticsRepository interface {
	CreateEvent(
		ctx context.Context,
		event *model.URLEvent,
	) error

	GetTotalClicks(
		ctx context.Context,
		shortCode string,
	) (int64, error)

	GetUniqueVisitors(
		ctx context.Context,
		shortCode string,
	) (int64, error)

	GetTopURLs(
		ctx context.Context,
		limit int,
	) ([]TopURL, error)

	GetClicksByDay(
		ctx context.Context,
		shortCode string,
	) ([]DailyClicks, error)
}
