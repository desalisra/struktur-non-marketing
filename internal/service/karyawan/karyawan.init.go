package karyawan

import (
	"context"
	entity "struktur-non-marketing/internal/entity/karyawan"
)

type Data interface {
	GetListKaryawanByNip(ctx context.Context, nip string) ([]entity.ListKaryawan, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}