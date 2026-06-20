package service

import (
	"context"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
)

type URLService interface {
	Create(
		ctx context.Context,
		originalURL string,
	) (*model.URL, error)

	Resolve(
		ctx context.Context,
		shortCode string,
	) (string, error)

	Get(
		ctx context.Context,
		shortCode string,
	) (*model.URL, error)
}
