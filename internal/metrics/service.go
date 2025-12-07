package metrics

import (
	"context"
	"time"

	"github.com/nm370130/realtime-analytics/internal/common"
)

type Service interface {
	GetSummary(ctx context.Context) (*SummaryResponse, error)
	GetHistory(ctx context.Context, metric string, interval string) (*HistoryResponse, error)
}

type service struct {
	repo  Repository
	cache *common.Cache
}

func NewService(repo Repository, cache *common.Cache) Service {
	return &service{
		repo:  repo,
		cache: cache,
	}
}

// SUMMARY METRICS 

func (s *service) GetSummary(ctx context.Context) (*SummaryResponse, error) {
	cacheKey := "metrics:summary"

	// 1. Check cache (5 sec TTL)
	var cached SummaryResponse
	found, _ := s.cache.Get(ctx, cacheKey, &cached)
	if found {
		return &cached, nil
	}

	// 2. Get Active Users (TTL-based keys, generic platform support)
	totalActive, platformCounts, err := s.repo.GetActiveUsers(ctx)
	if err != nil {
		return nil, err
	}

	// 3. Get API rejected count (4xx only, last 5 min)
	apiRejected, err := s.repo.GetAPIRejectedCount(ctx)
	if err != nil {
		return nil, err
	}

	// 4. New projects in last 7 days
	newProjCount, err := s.repo.GetNewProjectsLast7Days(ctx)
	if err != nil {
		return nil, err
	}

	// 5. Live projects count
	liveProjects, err := s.repo.GetTotalLiveProjects(ctx)
	if err != nil {
		return nil, err
	}

	// 6. Sensor online/offline counts (Redis → fallback → DB)
	sOnline, sOffline, err := s.repo.GetSensorCounts(ctx)
	if err != nil {
		return nil, err
	}

	resp := &SummaryResponse{
		ActiveUsers:           totalActive,
		ActiveUsersByPlatform: platformCounts,
		APIRejectedCount:      apiRejected,
		NewProjectsLast7Days:  newProjCount,
		TotalLiveProjects:     liveProjects,
		SensorsOnline:         sOnline,
		SensorsOffline:        sOffline,
	}

	// 7. Cache for 5 seconds
	_ = s.cache.Set(ctx, cacheKey, resp, 5*time.Second)

	return resp, nil
}

// METRICS HISTORY 

func (s *service) GetHistory(ctx context.Context, metric string, interval string) (*HistoryResponse, error) {

	// 1. Parse interval 
	dur, err := time.ParseDuration(interval)
	if err != nil {
		return nil, err
	}

	since := time.Now().Add(-dur)

	// 2. Fetch raw data points from DB, sorted oldest → newest
	rows, err := s.repo.GetMetricHistory(ctx, metric, since)
	if err != nil {
		return nil, err
	}

	points := make([]HistoryPoint, 0, len(rows))

	for _, row := range rows {
		points = append(points, HistoryPoint{
			Timestamp: row.TS.Format(time.RFC3339),
			Value:     row.Value,
		})
	}

	return &HistoryResponse{
		MetricType: metric,
		Interval:   interval,
		Points:     points,
	}, nil
}
