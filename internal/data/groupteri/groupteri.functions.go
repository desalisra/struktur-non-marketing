package grpteri

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/groupteri"
)

func (d Data) GetStrukturTeri(ctx context.Context, ptID string, dptID string) ([]entity.ListGrpteri, error) {
	resulst := []entity.ListGrpteri{}

	d.UpdateConn()

	rows, err := d.stmt[getGrpteri].QueryxContext(ctx, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_LIST_GRPTERI]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.ListGrpteri{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][SCAN_LIST_GRPTERI]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}