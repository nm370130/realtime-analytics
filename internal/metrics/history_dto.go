package metrics

type HistoryPoint struct {
	Timestamp string `json:"timestamp"`
	Value     int64  `json:"value"`
}

type HistoryResponse struct {
	MetricType string         `json:"metricType"`
	Interval   string         `json:"interval"`
	Points     []HistoryPoint `json:"points"`
}
