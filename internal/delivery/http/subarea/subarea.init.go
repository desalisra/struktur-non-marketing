package subarea

import (
	"context"
	entity "struktur-non-marketing/internal/entity/subarea"
)

type Service interface {
	GetStrukturSubarea(ctx context.Context, ptID string, dptID string) ([]entity.ListSubarea, error)
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