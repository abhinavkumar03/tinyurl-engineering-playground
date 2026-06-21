package model

type RedirectEvent struct {
	URLID     int64
	ShortCode string
	IPAddress string
	UserAgent string
	Referrer  string
}
