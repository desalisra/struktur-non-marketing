package region

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/region"
)

func (s Service) GetStrukturRegion(ctx context.Context, ptID string, dptID string) ([]entity.ListRegion, error) {

	resulst, err := s.data.GetStrukturRegion(ctx, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Call Get List Struktur Subarea]")
	}

	return resulst, err
}