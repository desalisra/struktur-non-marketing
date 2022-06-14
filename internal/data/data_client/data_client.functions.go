package data_client

import (
	"context"
	"fmt"
	"net/http"
	"struktur-non-marketing/internal/entity/client"
	"struktur-non-marketing/pkg/errors"
)

// GetCity ...
func (d Data) GetCity(ctx context.Context, _token, code string) (client.City, error) {
	clientResponse := client.City{}
	endpoint := d.baseURL + "/cities"
	body := map[string]interface{}{
		"code": code,
	}

	_ = body

	headers := make(http.Header)
	headers.Set("Authorization", _token)
	headers.Set("Content-Type", "application/json")

	res, err := d.client.GetJSON(ctx, endpoint, headers, &clientResponse)
	if err != nil {
		return clientResponse, errors.Wrap(err, "[DATA][Call Api GetCity]")
	}

	fmt.Println(res)

	return clientResponse, nil
}
