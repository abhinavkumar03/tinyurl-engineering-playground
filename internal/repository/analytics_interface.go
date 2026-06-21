package repository

import (
	"context"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
)

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
}
