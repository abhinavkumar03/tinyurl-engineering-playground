package model

import "time"

type URLEvent struct {
	ID        int64     `db:"id"`
	URLID     int64     `db:"url_id"`
	IPAddress string    `db:"ip_address"`
	UserAgent string    `db:"user_agent"`
	Referrer  string    `db:"referrer"`
	CreatedAt time.Time `db:"created_at"`
}

type EventMetadata struct {
	IPAddress string
	UserAgent string
	Referrer  string
}
