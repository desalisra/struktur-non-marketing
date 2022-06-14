package department

import (
	"context"
	entity "struktur-non-marketing/internal/entity/department"
)

type Service interface {
	GetDepartments(ctx context.Context) ([]entity.Department, error) 
	GetDepartmentById(ctx context.Context, dptID string) (entity.Department, error)
	GetPosition(ctx context.Context, dptID string) ([]entity.Position, error)
}

type (
	Handler struct {
		service Service
	}
)

func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}