package jabiklan

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/jabiklan"
)

func (s Service) GetJabIklan(ctx context.Context, ptID string, jabID string, dptID string) ([]entity.JabIklan, error) {

	resulst, err := s.data.GetJabIklan(ctx, ptID, jabID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][GET_JAB_IKLAN]")
	}

	return resulst, err
}