package department

import (
	"context"
	"struktur-non-marketing/pkg/errors"

	entity "struktur-non-marketing/internal/entity/department"
)

func (s Service) GetDepartments(ctx context.Context) ([]entity.Department, error) {

	resulst, err := s.data.GetDepartments(ctx)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][GET_ALL_DEPARTMENT]")
	}

	return resulst, err
}

func (s Service) GetDepartmentById(ctx context.Context, dptID string) (entity.Department, error) {

	resulst, err := s.data.GetDepartmentById(ctx, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][GET_DEPARTMENT_BY_ID]")
	}

	return resulst, err
}

func (s Service) GetPosition(ctx context.Context, dptID string) ([]entity.Position, error) {
	
	dpt, err := s.data.GetDepartmentById(ctx, dptID)
	divID :=  dpt.DivID

	resulst, err := s.data.GetPosition(ctx, divID, dptID)
	if err != nil {
		return resulst, errors.Wrap(err, "[SERVICE][GET_POSITION_BY_DIVID_AND_DPTID]")
	}

	return resulst, err
}