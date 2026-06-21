package dto

type AnalyticsResponse struct {
	ShortCode      string `json:"short_code"`
	TotalClicks    int64  `json:"total_clicks"`
	UniqueVisitors int64  `json:"unique_visitors"`
}
