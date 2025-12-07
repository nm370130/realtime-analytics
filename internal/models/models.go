package models

import "time"

// MODULES 

type Module struct {
	ID                  uint       `gorm:"primaryKey" json:"id"`
	ModuleName          string     `gorm:"column:module_name;not null" json:"moduleName"`
	CurrentVersion      string     `gorm:"column:current_version;not null" json:"currentVersion"`
	LastDeployedAt      *time.Time `gorm:"column:last_deployed_at" json:"lastDeployedAt"`
	UpcomingVersion     *string    `gorm:"column:upcoming_version" json:"upcomingVersion"`
	UpcomingReleaseDate *time.Time `gorm:"column:upcoming_release_date" json:"upcomingReleaseDate"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// SENSORS

type Sensor struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ProjectID uint   `gorm:"column:project_id" json:"projectId"`
	Type      string `gorm:"column:type;not null" json:"type"`
	Status    string `gorm:"column:status;not null" json:"status"` // "online" / "offline"

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// METRICS HISTORY

type MetricsHistory struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	MetricType string    `gorm:"column:metric_type;not null" json:"metricType"`
	TS         time.Time `gorm:"column:ts;not null" json:"timestamp"`
	Value      int64     `gorm:"column:value;not null" json:"value"`

	CreatedAt time.Time `json:"createdAt"`
}

// Force custom table name to match schema.sql
func (MetricsHistory) TableName() string {
	return "metrics_history"
}

// PROJECTS

type Project struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"column:name;not null" json:"name"`
	IsLive bool   `gorm:"column:is_live;not null" json:"isLive"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// USERS 

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"column:username;not null" json:"username"`
	Platform string `gorm:"column:platform;not null" json:"platform"` // "web", "mobile", etc.

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
