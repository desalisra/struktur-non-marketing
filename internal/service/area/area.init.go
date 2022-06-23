package area

import (
	"context"
	entity "struktur-non-marketing/internal/entity/area"
	"struktur-non-marketing/internal/entity/client"
)

type Data interface {
	GetListStrukturArea(ctx context.Context, periode string, ptID string, dptID string) ([]entity.ListArea, error)
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