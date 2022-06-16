package area

import (
	"context"
	entity "struktur-non-marketing/internal/entity/area"
)

type Service interface {
	GetStrukturArea(ctx context.Context, ptID string, dptID string) ([]entity.ListArea, error)
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