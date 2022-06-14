package jabiklan

import (
	"context"
	entity "struktur-non-marketing/internal/entity/jabiklan"
)

type Service interface {
	GetJabIklan(ctx context.Context, ptID string, jabID string, dptID string) ([]entity.JabIklan, error)
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