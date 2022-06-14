package jabiklan

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/jabiklan"
)

func (d Data) GetJabIklan(ctx context.Context, ptID string, jabID string, dptID string) ([]entity.JabIklan, error) {
	resulst := []entity.JabIklan{}

	d.UpdateConn()

	rows, err := d.stmt[getJabIklan].QueryxContext(ctx, ptID, jabID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_JABIKLAN_EXEC_QUERY]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.JabIklan{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][GET_JABIKLAN_SCAN_DATA]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}
