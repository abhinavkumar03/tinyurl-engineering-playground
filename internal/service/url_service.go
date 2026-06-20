package service

import (
	"context"
	"time"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/repository"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/pkg/shortener"
	"github.com/redis/go-redis/v9"
)

type URLService struct {
	repository *repository.URLRepository
	redis      *redis.Client
	baseURL    string
}

func NewURLService(
	repository *repository.URLRepository,
	redis *redis.Client,
	baseURL string,
) *URLService {

	return &URLService{
		repository: repository,
		redis:      redis,
		baseURL:    baseURL,
	}
}

func (s *URLService) Create(
	ctx context.Context,
	originalURL string,
) (*model.URL, error) {

	url := &model.URL{
		OriginalURL: originalURL,
	}

	if err := s.repository.Create(ctx, url); err != nil {
		return nil, err
	}

	shortCode := shortener.Encode(url.ID)

	if err := s.repository.UpdateShortCode(
		ctx,
		url.ID,
		shortCode,
	); err != nil {
		return nil, err
	}

	url.ShortCode = shortCode

	return url, nil
}

func (s *URLService) Resolve(
	ctx context.Context,
	shortCode string,
) (string, error) {

	cacheKey := "url:" + shortCode

	cached, err := s.redis.Get(ctx, cacheKey).Result()

	if err == nil {
		return cached, nil
	}

	url, err := s.repository.GetByShortCode(
		ctx,
		shortCode,
	)
	if err != nil {
		return "", err
	}

	_ = s.redis.Set(
		ctx,
		cacheKey,
		url.OriginalURL,
		time.Hour,
	).Err()

	_ = s.repository.IncrementClickCount(
		ctx,
		shortCode,
	)

	return url.OriginalURL, nil
}

func (s *URLService) Get(
	ctx context.Context,
	shortCode string,
) (*model.URL, error) {

	return s.repository.GetByShortCode(
		ctx,
		shortCode,
	)
}
