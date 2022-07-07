package groupteri

import (
	"context"
	entity "struktur-non-marketing/internal/entity/groupteri"
)

type Service interface {
	GetStruktur(ctx context.Context, periode, pt, dept string) ([]entity.Grpteri, error)
	AddStrukturTeri(ctx context.Context, request entity.Grpteri) (entity.ResMessage, error)
	EditStrukturTeri(ctx context.Context, r entity.Grpteri) (entity.ResMessage, error)
	DeleteStrukturTeri(ctx context.Context, r entity.Grpteri) (entity.ResMessage, error)
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