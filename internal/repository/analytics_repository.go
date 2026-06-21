package repository

import (
	"context"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresAnalyticsRepository struct {
	db *pgxpool.Pool
}

func NewAnalyticsRepository(
	db *pgxpool.Pool,
) *PostgresAnalyticsRepository {

	return &PostgresAnalyticsRepository{
		db: db,
	}
}

func (r *PostgresAnalyticsRepository) CreateEvent(
	ctx context.Context,
	event *model.URLEvent,
) error {

	query := `
	INSERT INTO url_events (
		url_id,
		ip_address,
		user_agent,
		referrer
	)
	VALUES (
		$1,
		$2,
		$3,
		$4
	)
	RETURNING
		id,
		created_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		event.URLID,
		event.IPAddress,
		event.UserAgent,
		event.Referrer,
	).Scan(
		&event.ID,
		&event.CreatedAt,
	)
}

func (r *PostgresAnalyticsRepository) GetTotalClicks(
	ctx context.Context,
	shortCode string,
) (int64, error) {

	query := `
	SELECT
		COUNT(*)
	FROM url_events ue
	INNER JOIN urls u
		ON ue.url_id = u.id
	WHERE u.short_code = $1
	`

	var total int64

	err := r.db.QueryRow(
		ctx,
		query,
		shortCode,
	).Scan(
		&total,
	)

	return total, err
}

func (r *PostgresAnalyticsRepository) GetUniqueVisitors(
	ctx context.Context,
	shortCode string,
) (int64, error) {

	query := `
	SELECT
		COUNT(
			DISTINCT ip_address
		)
	FROM url_events ue
	INNER JOIN urls u
		ON ue.url_id = u.id
	WHERE u.short_code = $1
	`

	var total int64

	err := r.db.QueryRow(
		ctx,
		query,
		shortCode,
	).Scan(
		&total,
	)

	return total, err
}

func (r *PostgresAnalyticsRepository) GetTopURLs(
	ctx context.Context,
	limit int,
) ([]TopURL, error) {

	query := `
	SELECT
		u.short_code,
		u.original_url,
		COUNT(ue.id) AS clicks
	FROM urls u
	LEFT JOIN url_events ue
		ON ue.url_id = u.id
	GROUP BY
		u.id,
		u.short_code,
		u.original_url
	ORDER BY clicks DESC
	LIMIT $1
	`

	rows, err := r.db.Query(
		ctx,
		query,
		limit,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []TopURL

	for rows.Next() {

		var item TopURL

		err := rows.Scan(
			&item.ShortCode,
			&item.OriginalURL,
			&item.Clicks,
		)
		if err != nil {
			return nil, err
		}

		results = append(
			results,
			item,
		)
	}

	return results, rows.Err()
}

func (r *PostgresAnalyticsRepository) GetClicksByDay(
	ctx context.Context,
	shortCode string,
) ([]DailyClicks, error) {

	query := `
	SELECT
		DATE(ue.created_at),
		COUNT(*)
	FROM url_events ue
	INNER JOIN urls u
		ON ue.url_id = u.id
	WHERE u.short_code = $1
	GROUP BY DATE(ue.created_at)
	ORDER BY DATE(ue.created_at)
	`

	rows, err := r.db.Query(
		ctx,
		query,
		shortCode,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []DailyClicks

	for rows.Next() {

		var item DailyClicks

		err := rows.Scan(
			&item.Date,
			&item.Clicks,
		)
		if err != nil {
			return nil, err
		}

		results = append(
			results,
			item,
		)
	}

	return results, rows.Err()
}
