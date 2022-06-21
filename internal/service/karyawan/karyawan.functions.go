package karyawan

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/karyawan"
)

func (s Service) GetListKaryawanByNip(ctx context.Context, nip string) ([]entity.ListKaryawan, error) {

	resulst, err := s.data.GetListKaryawanByNip(ctx, nip)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Get List Karyawan]")
	}

	return resulst, err
}