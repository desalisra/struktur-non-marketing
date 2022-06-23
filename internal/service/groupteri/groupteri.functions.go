package groupteri

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/groupteri"
)

func (s Service) GetStrukturTeri(ctx context.Context, periode string, ptID string, dptID string) ([]entity.ListGrpteri, error) {

	resulst, err := s.data.GetListStrukturTeri(ctx, periode, ptID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Call GetStrukturTeri]")
	}

	return resulst, err
}

func (s Service) AddStrukturTeri(ctx context.Context, request entity.AddGrpteri) (entity.ResMessage, error) {
	var (
		resulst entity.ResMessage
		data entity.AddGrpteri
		err error
		cdGroup string
		nip string
		name string
	)

	// Validasi Data
	
	// Generate CodeGroup
	cdGroup, err = s.data.MaxCodeGroup(ctx, request.Periode, request.DepartmentID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][get max CodeGroup]")
	}
	cdGroup = "0000" + cdGroup;
	cdGroup =  cdGroup[len(cdGroup)-4:]

	if request.Nip == "" {
		// Create Nip & Name Vacant 
		nip = "V" + request.DepartmentID[len(request.DepartmentID) - 2:] + cdGroup
		name = "(VACANT) " + request.PositionName
	
	} else {
		nip = request.Nip
		name = request.Name

		// Chek Nip Exist 
		chekNip, err := s.data.ChekNipExistOnDepartment(ctx, request.Periode, request.CompanyId, request.DepartmentID, nip)
		if err != nil {
			return resulst, errors.Wrap(err, "[SERVICE][get max CodeGroup]")
		}
		if chekNip > 0 {
			resulst.Message = name + " (" + nip + ") already registered in the department "
			return resulst, nil
		}
	}

	// Prepare Data Insert
	data.Periode = request.Periode
	data.CompanyId = request.CompanyId
	data.DepartmentID = request.DepartmentID
	data.CdGroup = cdGroup
	data.Nip = nip
	data.Name = name
	data.PositionID = request.PositionID
	data.PositionName = request.PositionName
	data.DateIn = request.DateIn
	data.Dummy = request.Dummy
	data.BranchID = request.BranchID
	data.CityID = request.CityID
	data.CdHead = request.CdHead

	err = s.data.InsertStrukturTeri(ctx, data)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Err when insert table structure]")
	}

	resulst.Message = "Data added successfully"
	return resulst, err
}