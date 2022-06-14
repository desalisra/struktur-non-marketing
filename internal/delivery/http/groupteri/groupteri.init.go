package groupteri

import (
	"context"
	entity "struktur-non-marketing/internal/entity/groupteri"
)

type Service interface {
	GetStrukturTeri(ctx context.Context, ptID string, dptID string) ([]entity.ListGrpteri, error)
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