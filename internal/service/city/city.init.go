package city

import (
	"context"
	entity "struktur-non-marketing/internal/entity/city"
)

type Data interface {
	GetCityById(ctx context.Context, cityID string) (entity.City, error)
	GetCityByName(ctx context.Context, cityName string) ([]entity.City, error)
	GetCityBranchByID(ctx context.Context, cityID string, branchID string) ([]entity.Branch, error)
	GetCityBranchByName(ctx context.Context, cityID string, branchName string) ([]entity.Branch, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
