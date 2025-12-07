package modules

import "time"

type ModuleResponse struct {
	ModuleName          string     `json:"moduleName"`
	CurrentVersion      string     `json:"currentVersion"`
	LastDeployedAt      *time.Time `json:"lastDeployedAt"`
	UpcomingVersion     *string    `json:"upcomingVersion"`
	UpcomingReleaseDate *string    `json:"upcomingReleaseDate"`
}
