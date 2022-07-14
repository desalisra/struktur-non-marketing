package nsm

import (
	"context"
	entity "struktur-non-marketing/internal/entity/nsm"
)

type Service interface {
	GetStrukturNsm(ctx context.Context, periode string, ptID string, dptID string, nip string) ([]entity.ListNsm, error)
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
