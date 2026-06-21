package dto

type URLAnalyticsResponse struct {
	ShortCode      string `json:"short_code"`
	TotalClicks    int64  `json:"total_clicks"`
	UniqueVisitors int64  `json:"unique_visitors"`
}

type TopURLResponse struct {
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
	Clicks      int64  `json:"clicks"`
}

type DailyClicksResponse struct {
	Date   string `json:"date"`
	Clicks int64  `json:"clicks"`
}

type DashboardResponse struct {
	TopURLs []TopURLResponse `json:"top_urls"`
}
