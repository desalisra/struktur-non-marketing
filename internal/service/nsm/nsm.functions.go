package nsm

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/nsm"
)

func (s Service) GetStrukturNsm(ctx context.Context, ptID string, dptID string) ([]entity.ListNsm, error) {

	resulst, err := s.data.GetStrukturNsm(ctx, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Call Get List Struktur NSM]")
	}

	return resulst, err
}