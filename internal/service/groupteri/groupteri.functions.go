package groupteri

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/groupteri"
)

func (s Service) GetStrukturTeri(ctx context.Context, ptID string, dptID string) ([]entity.ListGrpteri, error) {

	resulst, err := s.data.GetStrukturTeri(ctx, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Call GetStrukturTeri]")
	}

	return resulst, err
}