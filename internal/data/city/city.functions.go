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


func (d Data) GetBranchByCityId(ctx context.Context, cityID string) ([]entity.Bracnh, error) {
	resulst := []entity.Bracnh{}

	d.UpdateConn()

	rows, err := d.stmt[getBranchByCityId].QueryxContext(ctx, cityID)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_BRANCH_BY_CITYID_EXEC_QUERY]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.Bracnh{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][GET_BRANCH_SCAN_DATA]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}