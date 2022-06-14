package department

import (
	"context"
	entity "struktur-non-marketing/internal/entity/department"
)

type Data interface {
	GetDepartments(ctx context.Context) ([]entity.Department, error)
	GetDepartmentById(ctx context.Context, dptID string) (entity.Department, error)
	GetPosition(ctx context.Context, divID int, dptID string) ([]entity.Position, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}