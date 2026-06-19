package dto

type CreateURLRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type CreateURLResponse struct {
	ID          int64  `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	ShortURL    string `json:"short_url"`
}

type URLResponse struct {
	ID          int64  `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	ClickCount  int64  `json:"click_count"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type HealthResponse struct {
	Status   string `json:"status"`
	Postgres string `json:"postgres"`
	Redis    string `json:"redis"`
}
