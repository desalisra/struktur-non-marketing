package grpteri

import (
	"context"
	"fmt"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/groupteri"
)

func (d Data) GetListStrukturTeri(ctx context.Context, periode string, ptID string, dptID string) ([]entity.ListGrpteri, error) {
	resulst := []entity.ListGrpteri{}

	d.UpdateConn()

	query := fmt.Sprintf(`SELECT Grpt_CompanyId, Pt_Name, Grpt_DepartmentId, Dpt_Name, 
						Grpt_CdGroup, Grpt_Nip, Grpt_Name, Grpt_PositionId, Grpt_Position,
						Grpt_In, Grpt_Out, Grpt_DummyYN, Grpt_BranchId, Cab_Nama, Grpt_CityId, Kota_Name,
						Grpt_NipShadow, Grpt_NameShadow, Grpt_InShadow, Grpt_OutShadow, Grpt_DummyShadowYN,
						Grpt_Head, Sub_Nip, Sub_Name
					FROM Nm_Rayon_Grpteri_%s
					LEFT JOIN M_Pt ON Grpt_CompanyId = Pt_Id
					LEFT JOIN M_Departemen ON Grpt_DepartmentId = Dpt_Id
					LEFT JOIN M_Cabang ON Grpt_BranchId = Cab_Id
					LEFT JOIN M_Kota ON Grpt_CityId = Kota_Id
					LEFT JOIN Nm_Rayon_Subarea_%s ON Grpt_Head = Sub_CdGroup
						AND Grpt_CompanyId = Sub_CompanyId
						AND Grpt_DepartmentId = Sub_DepartmentId
					WHERE Grpt_ActiveYN = 'Y'
					AND Grpt_CompanyId = %s
					AND Grpt_DepartmentId = %s`, periode, periode, ptID, dptID)
	
	rows, err := d.db.QueryxContext(ctx, query)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][GET_LIST_GRPTERI]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.ListGrpteri{}
		if err = rows.StructScan(&row); err != nil {
			fmt.Println("Error Disini 1")
			return resulst, errors.Wrap(err, "[DATA][SCAN_LIST_GRPTERI]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}

func (d Data) GetListStrukturTeriByCodeGroup(ctx context.Context, periode string, cdGroup string, ptID string, dptID string) (entity.ListGrpteri, error) {
	resulst := entity.ListGrpteri{}

	d.UpdateConn()

	query := fmt.Sprintf(`SELECT Grpt_CompanyId, Pt_Name, Grpt_DepartmentId, Dpt_Name, 
						Grpt_CdGroup, Grpt_Nip, Grpt_Name, Grpt_PositionId, Grpt_Position,
						Grpt_In, Grpt_Out, Grpt_DummyYN, Grpt_BranchId, Cab_Nama, Grpt_CityId, Kota_Name,
						Grpt_NipShadow, Grpt_NameShadow, Grpt_InShadow, Grpt_OutShadow, Grpt_DummyShadowYN,
						Grpt_Head, Sub_Nip, Sub_Name
					FROM Nm_Rayon_Grpteri_%s
					LEFT JOIN M_Pt ON Grpt_CompanyId = Pt_Id
					LEFT JOIN M_Departemen ON Grpt_DepartmentId = Dpt_Id
					LEFT JOIN M_Cabang ON Grpt_BranchId = Cab_Id
					LEFT JOIN M_Kota ON Grpt_CityId = Kota_Id
					LEFT JOIN Nm_Rayon_Subarea_%s ON Grpt_Head = Sub_CdGroup
						AND Grpt_CompanyId = Sub_CompanyId
						AND Grpt_DepartmentId = Sub_DepartmentId
					WHERE Grpt_ActiveYN = 'Y'
					AND Grpt_CdGroup = %s
					AND Grpt_CompanyId = %s
					AND Grpt_DepartmentId = %s`, periode, periode, cdGroup, ptID, dptID)
	
	rows, err := d.db.QueryxContext(ctx, query)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][Get Struktur]")
	}

	defer rows.Close()
	
	for rows.Next() {
		if err = rows.StructScan(&resulst); err != nil {
			return resulst, errors.Wrap(err, "[DATA][Err Scan Struct]")
		}
	}
	
	return resulst, nil
}

func (d Data) MaxCodeGroup(ctx context.Context, periode string, dptID string) (string, error) {
	var resulst string

	d.UpdateConn()

	query := fmt.Sprintf(`SELECT IFNULL(MAX(Grpt_CdGroup), 0) + 1 AS MaxKode
						FROM Nm_Rayon_Grpteri_%s
						WHERE Grpt_CdGroup <> '9999'
						AND Grpt_DepartmentId = %s`, periode, dptID)

	if err := d.db.QueryRowxContext(ctx, query).Scan(&resulst); err != nil {
		return resulst, errors.Wrap(err, "[DATA][Get Max CodeGroup]")
	}	
	
	return resulst, nil
}

// Cek Nip Sudah terdaftar di Department Tersebut
func (d Data) ChekNipExistOnDepartment(ctx context.Context, periode string, pt string, dpt string, nip string) (int, error) {
	var resulst int

	d.UpdateConn()

	query := fmt.Sprintf(`SELECT COUNT(*) total
	FROM Nm_Rayon_Grpteri_%s
	LEFT JOIN Nm_Rayon_Subarea_%s 
		ON Grpt_Head = Sub_CdGroup
		AND Grpt_CompanyId = Sub_CompanyId
		AND Grpt_DepartmentId = Sub_DepartmentId
	LEFT JOIN Nm_Rayon_Area_%s 
		ON Sub_Head = Area_CdGroup
		AND Sub_CompanyId = Area_CompanyId
		AND Sub_DepartmentId = Area_DepartmentId
	LEFT JOIN Nm_Rayon_Region_%s 
		ON Area_Head = Reg_CdGroup
		AND Area_CompanyId = Reg_CompanyId
		AND Area_DepartmentId = Reg_DepartmentId
	LEFT JOIN Nm_Rayon_Nsm_%s 
		ON Reg_Head = Nsm_CdGroup
		AND Reg_CompanyId = Nsm_CompanyId
		AND Reg_DepartmentId = Nsm_DepartmentId
	WHERE Grpt_CompanyId = %s
	AND Grpt_DepartmentId = %s
	AND (Grpt_Nip = '%s' OR Sub_Nip = '%s' OR Area_Nip = '%s' OR Reg_Nip = '%s' OR Nsm_Nip = '%s')`, 
	periode, periode, periode, periode, periode, pt, dpt, nip, nip, nip, nip, nip)

	if err := d.db.QueryRowxContext(ctx, query).Scan(&resulst); err != nil {
		return resulst, errors.Wrap(err, "[DATA][Chek Nip Exist on Same Department]")
	}	
	
	return resulst, nil
}

func (d Data) InsertStrukturTeri(ctx context.Context, val entity.AddGrpteri) error {
	d.UpdateConn()

	query := fmt.Sprintf(`INSERT INTO Nm_Rayon_Grpteri_%s
	(Grpt_CompanyId, Grpt_DepartmentId, Grpt_CdGroup, Grpt_PositionId, Grpt_Position,
	 Grpt_Nip, Grpt_Name, Grpt_In, Grpt_DummyYN, Grpt_DummyShadowYN, 
	 Grpt_BranchId, Grpt_CityId, Grpt_Head, Grpt_ActiveYN, Grpt_UpdateId,
	 Grpt_UpdateTime
	) VALUES (
		'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %s
	)`, val.Periode, val.CompanyId, val.DepartmentID, val.CdGroup, val.PositionID, val.PositionName,
		val.Nip, val.Name, val.DateIn.String, val.Dummy.String, "Y",
		val.BranchID, val.CityID, val.CdHead.String, "Y", "0000", "NOW()")

	_, err := d.db.ExecContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "[DATA][Exec Query Insert]")
	}	
	
	return nil
}