package area

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/area"
)

func (d Data) GetStrukturArea(ctx context.Context, ptID string, dptID string) ([]entity.ListArea, error) {
	resulst := []entity.ListArea{}

	d.UpdateConn()

	rows, err := d.stmt[getArea].QueryxContext(ctx, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][Exec Query Get Struktur Area]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.ListArea{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][Scan Data Struktur Area]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}
