package groupteri

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/groupteri"
)

func (s Service) GetStruktur(ctx context.Context, periode, pt, dept string) ([]entity.Grpteri, error) {

	resulst, err := s.data.GetStrukturAll(ctx, periode, pt, dept)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Call GetStrukturTeri]")
	}

	return resulst, err
}

func (s Service) AddStrukturTeri(ctx context.Context, r entity.Grpteri) (entity.ResMessage, error) {
	var (
		resulst entity.ResMessage
		data 	entity.Grpteri
		cdGroup, nip, name string
		err error
	)

	// Validasi Data

	// Generate CodeGroup
	cdGroup, err = s.data.MaxCodeGroup(ctx, r.Periode, r.CompanyId, r.DepartmentId)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][get max CodeGroup]")
	}
	cdGroup = "0000" + cdGroup;
	cdGroup =  cdGroup[len(cdGroup)-4:]

	if r.Nip == "" {
		// Create Nip & Name Vacant 
		nip = "V" + r.DepartmentId[len(r.DepartmentId) - 2:] + cdGroup
		name = "(VACANT) " + r.PositionName
	
	} else {
		nip = r.Nip
		name = r.Name

		// Chek Nip Exist 
		chekNip, err := s.data.ChekNipExistOnDepartment(ctx, r.Periode, r.CompanyId, r.DepartmentId, nip)
		if err != nil {
			return resulst, errors.Wrap(err, "[SERVICE][Check Nip registered in the department]")
		}
		if chekNip > 0 {
			resulst.Message = name + " (" + nip + ") already registered in the department "
			return resulst, nil
		}
	}

	// Prepare data for Insert Database
	data.Periode		= r.Periode
	data.CompanyId		= r.CompanyId
	data.CompanyName    = r.CompanyName
	data.DepartmentId   = r.DepartmentId
	data.CdGroup        = cdGroup
	data.Nip            = nip
	data.Name           = name
	data.PositionId     = r.PositionId
	data.PositionName   = r.PositionName
	data.DateIn         = r.DateIn
	data.DateOut        = "NULL"
	data.Dummy          = r.Dummy
	data.BranchId       = r.BranchId
	data.CityId         = r.CityId
	data.NipShadow      = "NULL"
	data.NameShadow     = "NULL"
	data.DateInShadow   = "NULL"
	data.DateOutShadow  = "NULL"
	data.DummyShadow    = "Y"
	data.CdHead         = r.CdHead

	err = s.data.InsertNewStruktur(ctx, data)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Insert Table Struktur]")
	}
	
	data, err = s.data.GetStrukturByCdGroup(ctx, r.Periode, cdGroup, r.CompanyId, r.DepartmentId);

	resulst.Message = "Data added successfully"
	resulst.Data = data

	return resulst, err
}

func (s Service) EditStrukturTeri(ctx context.Context, r entity.Grpteri) (entity.ResMessage, error) {
	var (
		resulst entity.ResMessage
		data entity.Grpteri
		err error
		nip, name string 
		nipShadow, nameShadow string
	)

	// Validasi Data
	if r.Nip == "" || r.Nip[0:1] == "V" {
		// Create Nip & Name Vacant 
		nip = "V" + r.DepartmentId[len(r.DepartmentId) - 2:] + r.CdGroup
		name = "(VACANT) " + r.PositionName
	} else {
		nip = r.Nip
		name = r.Name

		// Check Nip Changed
		data, err = s.data.GetStrukturByCdGroup(ctx, r.Periode, r.CdGroup, r.CompanyId, r.DepartmentId);
		if err != nil {
			return resulst, errors.Wrap(err, "[SERVICE][Get Data by CdGroup]")	
		}

		if data.Nip != nip {
			// Check Nip Exist On Department
			chekNip, err := s.data.ChekNipExistOnDepartment(ctx, r.Periode, r.CompanyId, r.DepartmentId, nip)

			if err != nil {
				return resulst, errors.Wrap(err, "[SERVICE][Check Nip registered in the department]")
			}
			if chekNip > 0 {
				resulst.Message = name + " (" + nip + ") already registered in the department "
				return resulst, nil
			}
		}
	}


	if r.DummyShadow == "N" {
		nipShadow = r.NipShadow
		nameShadow = r.NameShadow

		if nipShadow != "" {
			// Chek Nip Shadow Exist On Department
			chekNip, err := s.data.ChekNipExistOnDepartment(ctx, r.Periode, r.CompanyId, r.DepartmentId, nipShadow)
			if err != nil {
				return resulst, errors.Wrap(err, "[SERVICE][Check Nip Shadow registered in the department]")
			}
			if chekNip > 0 {
				resulst.Message = "Shadow " + nameShadow + " (" + nipShadow + ") already registered in the department "
				return resulst, nil
			}
		}
	}else{
		nipShadow = ""
		nameShadow = ""
	}


	// Non Aktifkan Struktur Lama
	err = s.data.DeleteStruktur(ctx, r.Periode, r.CompanyId, r.DepartmentId, r.CdGroup)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Update ActiveYN => N]")	
	}

	// Prepare data for Insert Database
	data.Periode		= r.Periode
	data.CompanyId		= r.CompanyId
	data.CompanyName    = r.CompanyName
	data.DepartmentId   = r.DepartmentId
	data.CdGroup        = r.CdGroup
	data.Nip            = nip
	data.Name           = name
	data.PositionId     = r.PositionId
	data.PositionName   = r.PositionName
	data.DateIn         = r.DateIn
	data.DateOut        = "NULL"
	data.Dummy          = r.Dummy
	data.BranchId       = r.BranchId
	data.CityId         = r.CityId
	
	if nipShadow != "" {
		data.NipShadow   = "'" + nipShadow + "'"
	}else{
		data.NipShadow   = "NULL"
	}
	
	if nameShadow != "" {
		data.NameShadow   = "'" + nameShadow + "'"
	}else{
		data.NameShadow   = "NULL"
	}
	
	if r.DateInShadow != "" {
		data.DateInShadow   = "'" + r.DateInShadow + "'"
	}else{
		data.DateInShadow   = "NULL"
	}

	data.DateOutShadow  = "NULL"
	data.DummyShadow    = r.DummyShadow
	data.CdHead         = r.CdHead


	err = s.data.InsertNewStruktur(ctx, data)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Insert New Data After Edit]")
	}
	
	data, err = s.data.GetStrukturByCdGroup(ctx, r.Periode, r.CdGroup, r.CompanyId, r.DepartmentId);
	resulst.Message = "Data updated successfully"
	resulst.Data = data

	return resulst, err
}

func (s Service) DeleteStrukturTeri(ctx context.Context, r entity.Grpteri) (entity.ResMessage, error) {
	var (
		resulst entity.ResMessage
		err error
	)

	// Non Aktifkan Struktur 
	err = s.data.DeleteStruktur(ctx, r.Periode, r.CompanyId, r.DepartmentId, r.CdGroup)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Update ActiveYN => N]")	
	}

	return resulst, err
}