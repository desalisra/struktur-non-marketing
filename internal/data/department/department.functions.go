package department

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/department"
)

func (d Data) GetDepartments(ctx context.Context) ([]entity.Department, error) {
	resulst := []entity.Department{}

	d.UpdateConn()

	rows, err := d.stmt[getDepartments].QueryxContext(ctx)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_DEPARTMENT_EXEC_QUERY]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.Department{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][GET_DEPARTMENT_SCAN_DATA]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}

func (d Data) GetDepartmentById(ctx context.Context, dptID string) (entity.Department, error) {
	resulst := entity.Department{}

	d.UpdateConn()

	if err := d.stmt[getDepartmentById].QueryRowxContext(ctx, dptID).StructScan(&resulst); err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_DEPARTMENT_BYID_EXEC_QUERY]")
	}

	return resulst, nil
}

func (d Data) GetPosition(ctx context.Context, divID int, dptID string) ([]entity.Position, error) {
	resulst := []entity.Position{}

	d.UpdateConn()

	rows, err := d.stmt[getPosition].QueryxContext(ctx, divID, dptID, divID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_POSITION_EXEC_QUERY]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.Position{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][GET_POSITION_SCAN_DATA]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}