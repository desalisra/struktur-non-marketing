package karyawan

import (
	"context"
	entity "struktur-non-marketing/internal/entity/karyawan"
)

type Service interface {
	GetListKaryawanByNip(ctx context.Context, nip string) ([]entity.ListKaryawan, error)
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