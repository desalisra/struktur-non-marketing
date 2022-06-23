package region

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/region"
)

func (s Service) GetStrukturRegion(ctx context.Context, periode string, ptID string, dptID string) ([]entity.ListRegion, error) {

	resulst, err := s.data.GetListStrukturRegion(ctx, periode, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Call Get List Struktur Subarea]")
	}

	return resulst, err
}