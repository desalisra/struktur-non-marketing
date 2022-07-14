package nsm

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/nsm"
)

func (s Service) GetStrukturNsm(ctx context.Context, periode string, ptID string, dptID string, nip string) ([]entity.ListNsm, error) {

	resulst, err := s.data.GetListStrukturNsm(ctx, periode, ptID, dptID, nip)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Call Get List Struktur NSM]")
	}

	return resulst, err
}
