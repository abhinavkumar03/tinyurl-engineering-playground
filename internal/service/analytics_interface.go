package service

import (
	"context"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
)

type AnalyticsService interface {
	Track(
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
