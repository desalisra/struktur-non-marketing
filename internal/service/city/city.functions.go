package city

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/city"
)

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

func (s Service) GetCityBranchByID(ctx context.Context, cityID string, branchID string) ([]entity.Branch, error) {

	result, err := s.data.GetCityBranchByID(ctx, cityID, branchID)
	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][GET_ALL_CITYBRANCH_BY_ID]")
	}

	return result, err
}

func (s Service) GetCityBranchByName(ctx context.Context, cityID string, branchName string) ([]entity.Branch, error) {

	result, err := s.data.GetCityBranchByName(ctx, cityID, branchName)
	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][GET_ALL_CITYBRANCH_BY_Name]")
	}

	return result, err
}
