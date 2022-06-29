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
		resData entity.ListGrpteri
		reqData entity.AddGrpteri
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
			return resulst, errors.Wrap(err, "[SERVICE][Check Nip registered in the department]")
		}
		if chekNip > 0 {
			resulst.Message = name + " (" + nip + ") already registered in the department "
			return resulst, nil
		}
	}

	// Prepare Data Insert
	reqData.Periode = request.Periode
	reqData.CompanyId = request.CompanyId
	reqData.DepartmentID = request.DepartmentID
	reqData.CdGroup = cdGroup
	reqData.Nip = nip
	reqData.Name = name
	reqData.PositionID = request.PositionID
	reqData.PositionName = request.PositionName
	reqData.DateIn = request.DateIn
	reqData.Dummy = request.Dummy
	reqData.BranchID = request.BranchID
	reqData.CityID = request.CityID
	reqData.CdHead = request.CdHead

	err = s.data.InsertStrukturTeri(ctx, reqData)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][Err when insert table structure]")
	}
	
	resData, err = s.data.GetListStrukturTeriByCodeGroup(ctx, request.Periode, cdGroup, request.CompanyId, request.DepartmentID);


	resulst.Message = "Data added successfully"
	resulst.Data = resData

	return resulst, err
}