package metrics

type SummaryResponse struct {
	ActiveUsers           int64            `json:"activeUsers"`
	ActiveUsersByPlatform map[string]int64 `json:"activeUsersByPlatform"`
	APIRejectedCount      int64            `json:"apiRejectedCount"`
	NewProjectsLast7Days  int64            `json:"newProjectsLast7Days"`
	TotalLiveProjects     int64            `json:"totalLiveProjects"`
	SensorsOnline         int64            `json:"sensorsOnline"`
	SensorsOffline        int64            `json:"sensorsOffline"`
}
