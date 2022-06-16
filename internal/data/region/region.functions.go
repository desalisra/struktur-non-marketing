package region

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/region"
)

func (d Data) GetStrukturRegion(ctx context.Context, ptID string, dptID string) ([]entity.ListRegion, error) {
	resulst := []entity.ListRegion{}

	d.UpdateConn()

	rows, err := d.stmt[getRegion].QueryxContext(ctx, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_LIST_GRPTERI]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.ListRegion{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][SCAN_LIST_GRPTERI]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}
