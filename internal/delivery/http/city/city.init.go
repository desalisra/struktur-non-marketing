package city

import (
	"context"
	entity "struktur-non-marketing/internal/entity/city"
)

type Service interface {
	GetCityById(ctx context.Context, cityID string) (entity.City, error)
	GetCityByName(ctx context.Context, cityName string) ([]entity.City, error)
	GetCityBranchByID(ctx context.Context, cityID string, branchID string) ([]entity.Branch, error)
	GetCityBranchByName(ctx context.Context, cityID string, branchName string) ([]entity.Branch, error)
}

type (
	Handler struct {
		service Service
	}
)

func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}
