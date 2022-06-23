package nsm

import (
	"context"
	"struktur-non-marketing/internal/entity/client"
	entity "struktur-non-marketing/internal/entity/nsm"
)

type Data interface {
	GetListStrukturNsm(ctx context.Context, periode string, ptID string, dptID string) ([]entity.ListNsm, error)
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