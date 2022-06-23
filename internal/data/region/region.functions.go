package region

import (
	"context"
	"fmt"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/region"
)

func (d Data) GetListStrukturRegion(ctx context.Context, periode string, ptID string, dptID string) ([]entity.ListRegion, error) {
	resulst := []entity.ListRegion{}

	d.UpdateConn()

	query := fmt.Sprintf(
		`SELECT Reg_CompanyId, Pt_Name, Reg_DepartmentId, Dpt_Name, 
			Reg_CdGroup, Reg_Nip, Reg_Name, Reg_PositionId, Reg_Position,
			Reg_In, Reg_Out, Reg_DummyYN, Reg_BranchId, Cab_Nama, Reg_CityId, Kota_Name,
			Reg_NipShadow, Reg_NameShadow, Reg_InShadow, Reg_OutShadow, Reg_DummyShadowYN,
			Reg_Head, Nsm_Nip, Nsm_Name
		FROM Nm_Rayon_Region_%s
		LEFT JOIN M_Pt ON Reg_CompanyId = Pt_Id
		LEFT JOIN M_Departemen ON Reg_DepartmentId = Dpt_Id
		LEFT JOIN M_Cabang ON Reg_BranchId = Cab_Id
		LEFT JOIN M_Kota ON Reg_CityId = Kota_Id
		LEFT JOIN Nm_Rayon_Nsm_%s ON Reg_Head = Nsm_CdGroup
			AND Reg_CompanyId = Nsm_CompanyId
			AND Reg_DepartmentId = Nsm_DepartmentId
		WHERE Reg_ActiveYN = 'Y'
		AND Reg_CompanyId = %s
		AND Reg_DepartmentId = %s`, periode, periode, ptID, dptID)

	rows, err := d.db.QueryxContext(ctx, query)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][Get List Region]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.ListRegion{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][Scan List Region]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}
