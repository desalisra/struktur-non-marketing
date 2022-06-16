package subarea

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/subarea"
)

func (d Data) GetStrukturSubarea(ctx context.Context, ptID string, dptID string) ([]entity.ListSubarea, error) {
	resulst := []entity.ListSubarea{}

	d.UpdateConn()

	rows, err := d.stmt[getSubarea].QueryxContext(ctx, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_LIST_GRPTERI]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.ListSubarea{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][SCAN_LIST_GRPTERI]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}
