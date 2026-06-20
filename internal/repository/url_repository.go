package repository

import (
	"context"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLRepository struct {
	db *pgxpool.Pool
}

func NewURLRepository(db *pgxpool.Pool) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

func (r *URLRepository) Create(ctx context.Context, url *model.URL) error {

	query := `
	INSERT INTO urls (
		original_url
	)
	VALUES ($1)
	RETURNING
		id,
		created_at,
		updated_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		url.OriginalURL,
	).Scan(
		&url.ID,
		&url.CreatedAt,
		&url.UpdatedAt,
	)
}

func (r *URLRepository) UpdateShortCode(
	ctx context.Context,
	id int64,
	shortCode string,
) error {

	query := `
	UPDATE urls
	SET short_code = $1
	WHERE id = $2
	`

	_, err := r.db.Exec(
		ctx,
		query,
		shortCode,
		id,
	)

	return err
}

func (r *URLRepository) GetByShortCode(
	ctx context.Context,
	shortCode string,
) (*model.URL, error) {

	var url model.URL

	query := `
	SELECT
		id,
		original_url,
		short_code,
		click_count,
		created_at,
		updated_at
	FROM urls
	WHERE short_code = $1
	`

	err := r.db.QueryRow(
		ctx,
		query,
		shortCode,
	).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortCode,
		&url.ClickCount,
		&url.CreatedAt,
		&url.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *URLRepository) IncrementClickCount(
	ctx context.Context,
	shortCode string,
) error {

	query := `
	UPDATE urls
	SET click_count = click_count + 1
	WHERE short_code = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		shortCode,
	)

	return err
}
