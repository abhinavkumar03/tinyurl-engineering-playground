package repository

import (
	"context"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
)

type URLRepository interface {
	Create(
		ctx context.Context,
		url *model.URL,
	) error

	UpdateShortCode(
		ctx context.Context,
		id int64,
		shortCode string,
	) error

	GetByShortCode(
		ctx context.Context,
		shortCode string,
	) (*model.URL, error)

	IncrementClickCount(
		ctx context.Context,
		shortCode string,
	) error
}
