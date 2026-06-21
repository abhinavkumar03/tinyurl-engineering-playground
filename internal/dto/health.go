package dto

type HealthResponse struct {
	Status    string `json:"status"`
	Postgres  string `json:"postgres"`
	Redis     string `json:"redis"`
	Version   string `json:"version"`
	UptimeSec int64  `json:"uptime_sec"`
}
