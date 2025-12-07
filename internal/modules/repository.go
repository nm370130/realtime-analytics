package modules

import (
	"context"

	"github.com/nm370130/realtime-analytics/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetModules(ctx context.Context) ([]models.Module, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetModules(ctx context.Context) ([]models.Module, error) {
	var modules []models.Module
	err := r.db.WithContext(ctx).Find(&modules).Error
	return modules, err
}
