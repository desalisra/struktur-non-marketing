package grpteri

import (
	"context"
	"fmt"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/groupteri"
)

func (d Data) GetStrukturAll(ctx context.Context, periode, pt, dept string, nip string) ([]entity.Grpteri, error) {
	resulst := []entity.Grpteri{}

	d.UpdateConn()

	query := `SELECT '` + periode + `' AS Grpt_Periode, Grpt_CompanyId, Pt_Name, Grpt_DepartmentId, Dpt_Name, 
				Grpt_CdGroup, Grpt_Nip, Grpt_Name, Grpt_PositionId, Grpt_Position,
				Grpt_In, IFNULL(Grpt_Out, '') Grpt_Out, Grpt_DummyYN, Grpt_BranchId, Cab_Nama, Grpt_CityId, Kota_Name,
				IFNULL(Grpt_NipShadow, '') Grpt_NipShadow, IFNULL(Grpt_NameShadow, '') Grpt_NameShadow, IFNULL(Grpt_InShadow, '') Grpt_InShadow, 
				IFNULL(Grpt_OutShadow, '') Grpt_OutShadow, Grpt_DummyShadowYN,
				Grpt_Head, Sub_Nip, Sub_Name
			FROM Nm_Rayon_Grpteri_` + periode + `
			LEFT JOIN M_Pt ON Grpt_CompanyId = Pt_Id
			LEFT JOIN M_Departemen ON Grpt_DepartmentId = Dpt_Id
			LEFT JOIN M_Cabang ON Grpt_BranchId = Cab_Id
			LEFT JOIN M_Kota ON Grpt_CityId = Kota_Id
			LEFT JOIN Nm_Rayon_Subarea_` + periode + ` ON Grpt_Head = Sub_CdGroup
				AND Grpt_CompanyId = Sub_CompanyId
				AND Grpt_DepartmentId = Sub_DepartmentId
			WHERE Grpt_ActiveYN = 'Y'
			AND Grpt_CompanyId = '` + pt + `'
			AND Grpt_DepartmentId = '` + dept + `'
			AND Grpt_Nip LIKE '%` + nip + `%'`

	rows, err := d.db.QueryxContext(ctx, query)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][Exec All Struktur Teri]")
	}

	defer rows.Close()

	for rows.Next() {
		row := entity.Grpteri{}
		if err = rows.StructScan(&row); err != nil {
			fmt.Println("Error Disini 1")
			return resulst, errors.Wrap(err, "[DATA][Scan All Struktur Teri]")
		}
		resulst = append(resulst, row)
	}

	return resulst, nil
}

func (d Data) GetStrukturByCdGroup(ctx context.Context, periode string, cdGroup string, pt string, dept string) (entity.Grpteri, error) {
	resulst := entity.Grpteri{}

	d.UpdateConn()

	query := `SELECT '` + periode + `' Grpt_Periode, Grpt_CompanyId, Pt_Name, Grpt_DepartmentId, Dpt_Name, 
				Grpt_CdGroup, Grpt_Nip, Grpt_Name, Grpt_PositionId, Grpt_Position,
				Grpt_In, IFNULL(Grpt_Out, '') Grpt_Out, Grpt_DummyYN, Grpt_BranchId, Cab_Nama, Grpt_CityId, Kota_Name,
				IFNULL(Grpt_NipShadow, '') Grpt_NipShadow, IFNULL(Grpt_NameShadow, '') Grpt_NameShadow, IFNULL(Grpt_InShadow, '') Grpt_InShadow, 
				IFNULL(Grpt_OutShadow, '') Grpt_OutShadow, Grpt_DummyShadowYN,
				Grpt_Head, Sub_Nip, Sub_Name
			FROM Nm_Rayon_Grpteri_` + periode + `
			LEFT JOIN M_Pt ON Grpt_CompanyId = Pt_Id
			LEFT JOIN M_Departemen ON Grpt_DepartmentId = Dpt_Id
			LEFT JOIN M_Cabang ON Grpt_BranchId = Cab_Id
			LEFT JOIN M_Kota ON Grpt_CityId = Kota_Id
			LEFT JOIN Nm_Rayon_Subarea_` + periode + ` ON Grpt_Head = Sub_CdGroup
				AND Grpt_CompanyId = Sub_CompanyId
				AND Grpt_DepartmentId = Sub_DepartmentId
			WHERE Grpt_ActiveYN = 'Y'
			AND Grpt_CdGroup = '` + cdGroup + `'
			AND Grpt_CompanyId = '` + pt + `'
			AND Grpt_DepartmentId = '` + dept + `'`

	rows, err := d.db.QueryxContext(ctx, query)
	if err != nil {
		return resulst, errors.Wrap(err, "[DATA][Exec Struktur Teri By CdGroup]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&resulst); err != nil {
			return resulst, errors.Wrap(err, "[DATA][Scan Struktur Teri By CdGroup]")
		}
	}

	return resulst, nil
}

func (d Data) MaxCodeGroup(ctx context.Context, periode string, pt string, dept string) (string, error) {
	var resulst string

	d.UpdateConn()

	query := `SELECT IFNULL(MAX(Grpt_CdGroup), 0) + 1 AS MaxKode
			FROM Nm_Rayon_Grpteri_` + periode + `
			WHERE Grpt_CdGroup <> '9999'
			AND Grpt_CompanyId = '` + pt + `'
			AND Grpt_DepartmentId = '` + dept + `'`

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

func (d Data) InsertNewStruktur(ctx context.Context, e entity.Grpteri) error {
	d.UpdateConn()

	query := `INSERT INTO struktur_rayon.Nm_Rayon_Grpteri_` + e.Periode + ` (
		Grpt_CompanyId, Grpt_DepartmentId, Grpt_CdGroup, 
		Grpt_Nip, Grpt_Name, Grpt_PositionId, Grpt_Position, 
		Grpt_In, Grpt_Out, Grpt_DummyYN, Grpt_NipShadow, 
		Grpt_NameShadow, Grpt_InShadow, Grpt_OutShadow, Grpt_DummyShadowYN, 
		Grpt_BranchId, Grpt_CityId, Grpt_Head, Grpt_ActiveYN, 
		Grpt_UpdateId, Grpt_UpdateTime
	) VALUES (
		'` + e.CompanyId + `', '` + e.DepartmentId + `', '` + e.CdGroup + `', 
		'` + e.Nip + `', '` + e.Name + `', '` + e.PositionId + `', '` + e.PositionName + `', 
		'` + e.DateIn + `', ` + e.DateOut + `, 
		'` + e.Dummy + `', ` + e.NipShadow + `, ` + e.NameShadow + `, 
		` + e.DateInShadow + `, ` + e.DateOutShadow + `, '` + e.DummyShadow + `', 
		'` + e.BranchId + `', '` + e.CityId + `',  '` + e.CdHead + `', 		
		'Y', '0000', NOW()
	)`

	_, err := d.db.ExecContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "[DATA][Exec Query Insert]")
	}

	return nil
}

func (d Data) DeleteStruktur(ctx context.Context, periode string, pt string, dept string, cdGroup string) error {
	d.UpdateConn()

	query := `UPDATE struktur_rayon.Nm_Rayon_Grpteri_` + periode + `
			SET Grpt_ActiveYN = 'N',
				Grpt_UpdateId = '0000',
				Grpt_DeleteTime = NOW()
			WHERE Grpt_CompanyId = '` + pt + `'
			AND Grpt_DepartmentId = '` + dept + `'
			AND Grpt_CdGroup = '` + cdGroup + `'`

	_, err := d.db.ExecContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "[DATA][Exec Query Update ActiveYN => N]")
	}

	return nil
}
