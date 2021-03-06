package area

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/area"
)

func (s Service) GetStrukturArea(ctx context.Context, periode string, ptID string, dptID string) ([]entity.ListArea, error) {

	resulst, err := s.data.GetListStrukturArea(ctx, periode, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Call Get List Struktur Subarea]")
	}

	return resulst, err
}
