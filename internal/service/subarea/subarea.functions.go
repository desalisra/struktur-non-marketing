package subarea

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/subarea"
)

func (s Service) GetStrukturSubarea(ctx context.Context, ptID string, dptID string) ([]entity.ListSubarea, error) {

	resulst, err := s.data.GetStrukturSubarea(ctx, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Call Get List Struktur Subarea]")
	}

	return resulst, err
}