package nsm

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/nsm"
)

func (d Data) GetStrukturNsm(ctx context.Context, ptID string, dptID string) ([]entity.ListNsm, error) {
	resulst := []entity.ListNsm{}

	d.UpdateConn()

	rows, err := d.stmt[getNsm].QueryxContext(ctx, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_LIST_GRPTERI]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.ListNsm{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][SCAN_LIST_GRPTERI]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}
