package model

import "time"

type URL struct {
	ID          int64     `db:"id"`
	OriginalURL string    `db:"original_url"`
	ShortCode   string    `db:"short_code"`
	ClickCount  int64     `db:"click_count"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
