package service

import (
	"context"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/repository"
)

type AnalyticsService interface {
	Track(
		ctx context.Context,
		event model.RedirectEvent,
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
	) ([]repository.TopURL, error)

	GetClicksByDay(
		ctx context.Context,
		shortCode string,
	) ([]repository.DailyClicks, error)
}
