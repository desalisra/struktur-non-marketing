package groupteri

import (
	"context"
	"struktur-non-marketing/internal/entity/client"
	entity "struktur-non-marketing/internal/entity/groupteri"
)

type Data interface {
	GetListStrukturTeri(ctx context.Context, periode string, ptID string, dptID string) ([]entity.ListGrpteri, error)
	MaxCodeGroup(ctx context.Context, periode string, dptID string) (string, error)
	ChekNipExistOnDepartment(ctx context.Context, periode string, pt string, dpt string, nip string) (int, error)
	InsertStrukturTeri(ctx context.Context, val entity.AddGrpteri) error
}

type Client interface {
	GetCity(ctx context.Context, _token, code string) (client.City, error)
}

type Service struct {
	data Data
	client Client
}

func New(data Data, client Client) Service {
	return Service{
		data: data,
		client: client,
	}
}