package sensors

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllSensorTypes(ctx context.Context) ([]string, error)
	GetTypeBreakdown(ctx context.Context) (map[string]int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Fetch All Distinct Sensor Types

func (r *repository) GetAllSensorTypes(ctx context.Context) ([]string, error) {
	var types []string

	err := r.db.WithContext(ctx).
		Table("sensors").
		Select("DISTINCT type").
		Find(&types).Error

	if err != nil {
		return nil, err
	}

	return types, nil
}

// Get Sensor Count by Type

func (r *repository) GetTypeBreakdown(ctx context.Context) (map[string]int64, error) {

	type Row struct {
		Type  string
		Count int64
	}

	var rows []Row

	// Count all sensors grouped by type (include online + offline)
	// Example result:
	//  type          count
	//  temperature   10
	//  humidity      5
	//
	err := r.db.WithContext(ctx).
		Table("sensors").
		Select("type, COUNT(*) AS count").
		Group("type").
		Find(&rows).Error

	if err != nil {
		return nil, err
	}

	result := make(map[string]int64)

	// First: fetch ALL sensor types 
	allTypes, err := r.GetAllSensorTypes(ctx)
	if err != nil {
		return nil, err
	}

	// Initialize map with ZERO values for each type
	for _, t := range allTypes {
		result[t] = 0
	}

	// Fill actual counts
	for _, row := range rows {
		result[row.Type] = row.Count
	}

	return result, nil
}
