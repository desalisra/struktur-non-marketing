package city

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/city"
)

func (d Data) GetCitys(ctx context.Context) ([]entity.City, error) {
	resulst := []entity.City{}

	d.UpdateConn()

	rows, err := d.stmt[getCitys].QueryxContext(ctx)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_CITYS_EXEC_QUERY]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.City{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][GET_CITYS_SCAN_DATA]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}

func (d Data) GetCityById(ctx context.Context, cityID string) (entity.City, error) {
	resulst := entity.City{}

	d.UpdateConn()

	if err := d.stmt[getCityById].QueryRowxContext(ctx, cityID).StructScan(&resulst); err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_CITY_BYID_EXEC_QUERY]")
	}

	return resulst, nil
}

func (d Data) GetCityByName(ctx context.Context, cityName string) ([]entity.City, error) {
	resulst := []entity.City{}

	d.UpdateConn()

	qCityName := "%" + cityName + "%"

	rows, err := d.stmt[getCityByName].QueryxContext(ctx, qCityName)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_CITY_BY_NAME_EXEC_QUERY]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.City{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][GET_POSITION_SCAN_DATA]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}

func (d Data) GetCityBranchByID(ctx context.Context, cityID string, branchID string) ([]entity.Branch, error) {
	result := []entity.Branch{}

	d.UpdateConn()

	qCityName := cityID
	qBranchID := "%" + branchID + "%"

	rows, err := d.stmt[getCityBranchById].QueryxContext(ctx, qCityName, qBranchID)
	if err != nil {
		return result, errors.Wrap(err, "[DATA][GET_CITYBRANCH_BY_NAME_EXEC_QUERY]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.Branch{}
		if err = rows.StructScan(&row); err != nil {
			return result, errors.Wrap(err, "[DATA][GET_BRANCH_SCAN_DATA]")
		}
		result = append(result, row)
	}

	return result, nil
}

func (d Data) GetCityBranchByName(ctx context.Context, cityID string, branchName string) ([]entity.Branch, error) {
	result := []entity.Branch{}

	d.UpdateConn()

	qCityName := cityID
	qBranchName := "%" + branchName + "%"

	rows, err := d.stmt[getCityBranchByName].QueryxContext(ctx, qCityName, qBranchName)
	if err != nil {
		return result, errors.Wrap(err, "[DATA][GET_CITYBRANCH_BY_ID_EXEC_QUERY]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.Branch{}
		if err = rows.StructScan(&row); err != nil {
			return result, errors.Wrap(err, "[DATA][GET_BRANCH_SCAN_DATA]")
		}
		result = append(result, row)
	}

	return result, nil
}
