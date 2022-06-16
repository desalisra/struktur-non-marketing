package region

import (
	"context"
	"struktur-non-marketing/internal/entity/client"
	entity "struktur-non-marketing/internal/entity/region"
)

type Data interface {
	GetStrukturRegion(ctx context.Context, ptID string, dptID string) ([]entity.ListRegion, error)
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