package region

import (
	"context"
	entity "struktur-non-marketing/internal/entity/region"
)

type Service interface {
	GetStrukturRegion(ctx context.Context, ptID string, dptID string) ([]entity.ListRegion, error)
}

type (
	Handler struct {
		service Service
	}
)

func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}