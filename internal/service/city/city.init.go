package city

import (
	"context"
	entity "struktur-non-marketing/internal/entity/city"
)

type Data interface {
	GetCitys(ctx context.Context) ([]entity.City, error)
	GetCityById(ctx context.Context, cityID string) (entity.City, error)
	GetCityByName(ctx context.Context, cityName string) ([]entity.City, error)
	GetBranchByCityId(ctx context.Context, cityID string) ([]entity.Bracnh, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}