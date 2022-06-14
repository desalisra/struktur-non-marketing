package city

import (
	"context"
	entity "struktur-non-marketing/internal/entity/city"
)

type Service interface {
	GetCitys(ctx context.Context) ([]entity.City, error)
	GetCityById(ctx context.Context, cityID string) (entity.City, error)
	GetCityByName(ctx context.Context, cityName string) ([]entity.City, error)
	GetBranchByCityId(ctx context.Context, cityID string) ([]entity.Bracnh, error)
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