package jabiklan

import (
	"context"
	entity "struktur-non-marketing/internal/entity/jabiklan"
)

type Data interface {
	GetJabIklan(ctx context.Context, ptID string, jabID string, dptID string) ([]entity.JabIklan, error) 
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}