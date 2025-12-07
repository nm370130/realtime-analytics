package modules

import (
	"context"
)

type Service interface {
	GetModules(ctx context.Context) ([]ModuleResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// List All Modulus

func (s *service) GetModules(ctx context.Context) ([]ModuleResponse, error) {

	rows, err := s.repo.GetModules(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]ModuleResponse, 0, len(rows))

	for _, m := range rows {

		var upcomingVersion *string
		if m.UpcomingVersion != nil {
			upcomingVersion = m.UpcomingVersion
		} else {
			upcomingVersion = nil
		}

		var upcomingReleaseDate *string
		if m.UpcomingReleaseDate != nil {
			str := m.UpcomingReleaseDate.Format("2006-01-02")
			upcomingReleaseDate = &str
		} else {
			upcomingReleaseDate = nil
		}

		// Add clean response item
		resp = append(resp, ModuleResponse{
			ModuleName:          m.ModuleName,
			CurrentVersion:      m.CurrentVersion,
			LastDeployedAt:      m.LastDeployedAt,
			UpcomingVersion:     upcomingVersion,
			UpcomingReleaseDate: upcomingReleaseDate,
		})
	}

	return resp, nil
}
