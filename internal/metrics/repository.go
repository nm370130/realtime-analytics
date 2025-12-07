package metrics

import (
	"context"
	"strings"
	"time"

	"github.com/nm370130/realtime-analytics/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repository interface {
	GetActiveUsers(ctx context.Context) (int64, map[string]int64, error)
	GetAPIRejectedCount(ctx context.Context) (int64, error)
	GetNewProjectsLast7Days(ctx context.Context) (int64, error)
	GetTotalLiveProjects(ctx context.Context) (int64, error)
	GetSensorCounts(ctx context.Context) (int64, int64, error)
	GetMetricHistory(ctx context.Context, metric string, since time.Time) ([]models.MetricsHistory, error)
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) Repository {
	return &repository{db: db, redis: redis}
}

// Active Users (TTL-based)

func (r *repository) GetActiveUsers(ctx context.Context) (int64, map[string]int64, error) {
	keys, err := r.redis.Keys(ctx, "active_users:*").Result()
	if err != nil {
		return 0, nil, err
	}

	platformCounts := make(map[string]int64)
	var total int64 = 0

	for _, key := range keys {
		val, err := r.redis.Get(ctx, key).Int64()
		if err != nil {
			continue
		}

		platform := strings.TrimPrefix(key, "active_users:")
		platformCounts[platform] = val
		total += val
	}

	return total, platformCounts, nil
}

// API Rejected Count (last 5 minutes)

func (r *repository) GetAPIRejectedCount(ctx context.Context) (int64, error) {
	val, err := r.redis.Get(ctx, "api_rejected:5min").Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return val, err
}

// MySQL Queries

func (r *repository) GetNewProjectsLast7Days(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Table("projects").
		Where("created_at >= ?", time.Now().Add(-7*24*time.Hour)).
		Count(&count).Error
	return count, err
}

func (r *repository) GetTotalLiveProjects(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Table("projects").
		Where("is_live = ?", true).
		Count(&count).Error
	return count, err
}

// Sensors Online/Offline (Redis → fallback to DB) 

func (r *repository) GetSensorCounts(ctx context.Context) (int64, int64, error) {
	online, err1 := r.redis.Get(ctx, "sensors:online_count").Int64()
	offline, err2 := r.redis.Get(ctx, "sensors:offline_count").Int64()

	if err1 == nil && err2 == nil {
		return online, offline, nil
	}

	// fallback to DB
	var onlineDB int64
	var offlineDB int64

	r.db.WithContext(ctx).Table("sensors").Where("status = 'online'").Count(&onlineDB)
	r.db.WithContext(ctx).Table("sensors").Where("status = 'offline'").Count(&offlineDB)

	return onlineDB, offlineDB, nil
}

// Metric History

func (r *repository) GetMetricHistory(ctx context.Context, metric string, since time.Time) ([]models.MetricsHistory, error) {
	var rows []models.MetricsHistory
	err := r.db.WithContext(ctx).
		Where("metric_type = ? AND ts >= ?", metric, since).
		Order("ts ASC").
		Find(&rows).Error

	return rows, err
}
