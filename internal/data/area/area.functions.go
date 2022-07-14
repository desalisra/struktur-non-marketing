package area

import (
	"context"
	"fmt"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/area"
)

func (d Data) GetListStrukturArea(ctx context.Context, periode string, ptID string, dptID string, nip string) ([]entity.ListArea, error) {
	resulst := []entity.ListArea{}

	d.UpdateConn()

	qNip := "'%" + nip + "%'"

	query := fmt.Sprintf(`SELECT Area_CompanyId, Pt_Name, Area_DepartmentId, Dpt_Name, 
							Area_CdGroup, Area_Nip, Area_Name, Area_PositionId, Area_Position,
							Area_In, Area_Out, Area_DummyYN, Area_BranchId, Cab_Nama, Area_CityId, Kota_Name,
							Area_NipShadow, Area_NameShadow, Area_InShadow, Area_OutShadow, Area_DummyShadowYN,
							Area_Head, Reg_Nip, Reg_Name
						FROM Nm_Rayon_Area_%s
						LEFT JOIN M_Pt ON Area_CompanyId = Pt_Id
						LEFT JOIN M_Departemen ON Area_DepartmentId = Dpt_Id
						LEFT JOIN M_Cabang ON Area_BranchId = Cab_Id
						LEFT JOIN M_Kota ON Area_CityId = Kota_Id
						LEFT JOIN Nm_Rayon_Region_%s ON Area_Head = Reg_CdGroup
							AND Area_CompanyId = Reg_CompanyId
							AND Area_DepartmentId = Reg_DepartmentId
						WHERE Area_ActiveYN = 'Y'
						AND Area_CompanyId = %s
						AND Area_DepartmentId = %s
						AND Area_Nip LIKE %s`,
		periode, periode, ptID, dptID, qNip)

	rows, err := d.db.QueryxContext(ctx, query)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][Exec Query Get Struktur Area]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.ListArea{}
		if err = rows.StructScan(&row); err != nil {
			return resulst, errors.Wrap(err, "[DATA][Scan Data Struktur Area]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}
