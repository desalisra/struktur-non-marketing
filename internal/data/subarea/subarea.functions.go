package subarea

import (
	"context"
	"fmt"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/subarea"
)

func (d Data) GetListStrukturSubarea(ctx context.Context, periode string, ptID string, dptID string) ([]entity.ListSubarea, error) {
	resulst := []entity.ListSubarea{}

	d.UpdateConn()

	query := fmt.Sprintf(
		`SELECT Sub_CompanyId, Pt_Name, Sub_DepartmentId, Dpt_Name, 
			Sub_CdGroup, Sub_Nip, Sub_Name, Sub_PositionId, Sub_Position,
			Sub_In, Sub_Out, Sub_DummyYN, Sub_BranchId, Cab_Nama, Sub_CityId, Kota_Name,
			Sub_NipShadow, Sub_NameShadow, Sub_InShadow, Sub_OutShadow, Sub_DummyShadowYN,
			Sub_Head, Area_Nip, Area_Name
		FROM Nm_Rayon_Subarea_%s
		LEFT JOIN M_Pt ON Sub_CompanyId = Pt_Id
		LEFT JOIN M_Departemen ON Sub_DepartmentId = Dpt_Id
		LEFT JOIN M_Cabang ON Sub_BranchId = Cab_Id
		LEFT JOIN M_Kota ON Sub_CityId = Kota_Id
		LEFT JOIN Nm_Rayon_Area_%s ON Sub_Head = Area_CdGroup
			AND Sub_CompanyId = Area_CompanyId
			AND Sub_DepartmentId = Area_DepartmentId
		WHERE Sub_ActiveYN = 'Y'
		AND Sub_CompanyId = %s
		AND Sub_DepartmentId = %s`, periode, periode, ptID, dptID)

	rows, err := d.db.QueryxContext(ctx, query)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][Get List Subarea]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.ListSubarea{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][Scan Data Subarea]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}

func (d Data) MaxCodeGroup(ctx context.Context, periode string, dptID string) (string, error) {
	var resulst string

	d.UpdateConn()

	query := fmt.Sprintf(`SELECT IFNULL(MAX(Sub_CdGroup), 0) + 1 AS MaxKode
						FROM Nm_Rayon_Subarea_%s
						WHERE Sub_CdGroup <> '9999'
						AND Sub_DepartmentId = %s`, periode, dptID)

	if err := d.db.QueryRowxContext(ctx, query).Scan(&resulst); err != nil {
		return resulst, errors.Wrap(err, "[DATA][Get Max CodeGroup]")
	}	
	
	return resulst, nil
}