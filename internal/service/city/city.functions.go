package city

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/city"
)

func (s Service) GetCitys(ctx context.Context) ([]entity.City, error) {

	resulst, err := s.data.GetCitys(ctx)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][GET_ALL_City]")
	}

	return resulst, err
}

func (s Service) GetCityById(ctx context.Context, cityID string) (entity.City, error) {

	resulst, err := s.data.GetCityById(ctx, cityID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][GET_ALL_City_BYID]")
	}

	return resulst, err
}

func (s Service) GetCityByName(ctx context.Context, cityName string) ([]entity.City, error) {

	resulst, err := s.data.GetCityByName(ctx, cityName)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][GET_ALL_City_BY_Name]")
	}

	return resulst, err
}

func (s Service) GetBranchByCityId(ctx context.Context, cityID string) ([]entity.Bracnh, error) {

	resulst, err := s.data.GetBranchByCityId(ctx, cityID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][GET_ALL_City_BYID]")
	}

	return resulst, err
}