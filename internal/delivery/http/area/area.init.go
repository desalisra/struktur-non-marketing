package area

import (
	"context"
	entity "struktur-non-marketing/internal/entity/area"
)

type Service interface {
	GetStrukturArea(ctx context.Context, periode string, ptID string, dptID string, nip string) ([]entity.ListArea, error)
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
