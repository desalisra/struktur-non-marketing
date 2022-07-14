package groupteri

import (
	"context"
	"struktur-non-marketing/internal/entity/client"
	entity "struktur-non-marketing/internal/entity/groupteri"
)

type Data interface {
	GetStrukturAll(ctx context.Context, periode, pt, dept string, nip string) ([]entity.Grpteri, error)
	GetStrukturByCdGroup(ctx context.Context, periode string, cdGroup string, pt string, dept string) (entity.Grpteri, error)
	MaxCodeGroup(ctx context.Context, periode string, pt string, dept string) (string, error)
	ChekNipExistOnDepartment(ctx context.Context, periode string, pt string, dpt string, nip string) (int, error)
	InsertNewStruktur(ctx context.Context, e entity.Grpteri) error
	DeleteStruktur(ctx context.Context, periode string, pt string, dept string, cdGroup string) error
}

type Client interface {
	GetCity(ctx context.Context, _token, code string) (client.City, error)
}

type Service struct {
	data   Data
	client Client
}

func New(data Data, client Client) Service {
	return Service{
		data:   data,
		client: client,
	}
}
