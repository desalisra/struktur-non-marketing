package groupteri

import (
	"context"
	entity "struktur-non-marketing/internal/entity/groupteri"
)

type Service interface {
	GetStrukturTeri(ctx context.Context, periode string, ptID string, dptID string) ([]entity.ListGrpteri, error)
	AddStrukturTeri(ctx context.Context, request entity.AddGrpteri) (entity.ResMessage, error)
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