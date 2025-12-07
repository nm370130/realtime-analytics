package sensors

import (
	"context"
	"time"

	"github.com/nm370130/realtime-analytics/internal/common"
)

type Service interface {
	GetTypeBreakdown(ctx context.Context) (TypeBreakdownResponse, error)
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

// SENSOR TYPE BREAKDOWN 

func (s *service) GetTypeBreakdown(ctx context.Context) (TypeBreakdownResponse, error) {

	// Cache key 
	cacheKey := "sensors:type-breakdown"

	// Check in Cache (30 seconds TTL)
	var cached TypeBreakdownResponse
	found, _ := s.cache.Get(ctx, cacheKey, &cached)
	if found {
		return cached, nil
	}

	// Fetch generic type breakdown from repository
	//    - Includes all sensor types
	//    - Includes types with 0 count
	//    - Counts online + offline
	result, err := s.repo.GetTypeBreakdown(ctx)
	if err != nil {
		return nil, err
	}

	// Store in Redis Cache (30s)
	_ = s.cache.Set(ctx, cacheKey, result, 30*time.Second)

	return result, nil
}
