package service

import (
	"context"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/repository"
)

type AnalyticsServiceImpl struct {
	analyticsRepository repository.AnalyticsRepository
	urlRepository       repository.URLRepository
}

func NewAnalyticsService(
	analyticsRepository repository.AnalyticsRepository,
	urlRepository repository.URLRepository,
) *AnalyticsServiceImpl {

	return &AnalyticsServiceImpl{
		analyticsRepository: analyticsRepository,
		urlRepository:       urlRepository,
	}
}

func (s *AnalyticsServiceImpl) Track(
	ctx context.Context,
	event model.RedirectEvent,
) error {

	url, err := s.urlRepository.GetByID(
		ctx,
		event.URLID,
	)
	if err != nil {
		return err
	}

	urlEvent := &model.URLEvent{
		URLID:     url.ID,
		IPAddress: event.IPAddress,
		UserAgent: event.UserAgent,
		Referrer:  event.Referrer,
	}

	if err := s.analyticsRepository.CreateEvent(
		ctx,
		urlEvent,
	); err != nil {
		return err
	}

	if err := s.urlRepository.IncrementClickCount(
		ctx,
		url.ShortCode,
	); err != nil {
		return err
	}

	return nil
}

func (s *AnalyticsServiceImpl) GetTotalClicks(
	ctx context.Context,
	shortCode string,
) (int64, error) {

	return s.analyticsRepository.GetTotalClicks(
		ctx,
		shortCode,
	)
}

func (s *AnalyticsServiceImpl) GetUniqueVisitors(
	ctx context.Context,
	shortCode string,
) (int64, error) {

	return s.analyticsRepository.GetUniqueVisitors(
		ctx,
		shortCode,
	)
}

func (s *AnalyticsServiceImpl) GetTopURLs(
	ctx context.Context,
	limit int,
) ([]repository.TopURL, error) {

	return s.analyticsRepository.GetTopURLs(
		ctx,
		limit,
	)
}

func (s *AnalyticsServiceImpl) GetClicksByDay(
	ctx context.Context,
	shortCode string,
) ([]repository.DailyClicks, error) {

	return s.analyticsRepository.GetClicksByDay(
		ctx,
		shortCode,
	)
}
