package nsm

import (
	"context"
	"fmt"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/nsm"
)

func (d Data) GetListStrukturNsm(ctx context.Context, periode string, ptID string, dptID string, nip string) ([]entity.ListNsm, error) {
	resulst := []entity.ListNsm{}

	d.UpdateConn()

	qNip := "'%" + nip + "%'"

	query := fmt.Sprintf(`SELECT Nsm_CompanyId, Pt_Name, Nsm_DepartmentId, Dpt_Name, 
							Nsm_CdGroup, Nsm_Nip, Nsm_Name, Nsm_PositionId, Nsm_Position,
							Nsm_In, Nsm_Out, Nsm_DummyYN, Nsm_BranchId, Cab_Nama, Nsm_CityId, Kota_Name,
							Nsm_NipShadow, Nsm_NameShadow, Nsm_InShadow, Nsm_OutShadow, Nsm_DummyShadowYN
						FROM Nm_Rayon_Nsm_%s
						LEFT JOIN M_Pt ON Nsm_CompanyId = Pt_Id
						LEFT JOIN M_Departemen ON Nsm_DepartmentId = Dpt_Id
						LEFT JOIN M_Cabang ON Nsm_BranchId = Cab_Id
						LEFT JOIN M_Kota ON Nsm_CityId = Kota_Id
						WHERE Nsm_ActiveYN = 'Y'
						AND Nsm_CompanyId = %s
						AND Nsm_DepartmentId = %s
						AND Nsm_Nip LIKE %s`, periode, ptID, dptID, qNip)

	rows, err := d.db.QueryxContext(ctx, query)
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
