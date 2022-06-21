package karyawan

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/karyawan"
)

func (d Data) GetListKaryawanByNip(ctx context.Context, nip string) ([]entity.ListKaryawan, error) {
	resulst := []entity.ListKaryawan{}

	d.UpdateConn()

	nip = "%" + nip + "%"
	rows, err := d.stmt[getKaryawan].QueryxContext(ctx, nip)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][Get List Karyawan]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.ListKaryawan{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][GET_JABIKLAN_SCAN_DATA]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}
