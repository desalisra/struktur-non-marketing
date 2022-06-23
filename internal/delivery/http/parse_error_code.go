package http

import (
	"strings"
	"struktur-non-marketing/pkg/response"
)

// ParseErrorCode ...
func ParseErrorCode(err string) response.Response {
	errResp := response.Error{}

	switch {
	case strings.Contains(err, "401"):
		errResp = response.Error{
			Status: true,
			Msg:    "Unauthorized",
			Code:   401,
		}
	case strings.Contains(err, "10001"):
		errResp = response.Error{
			Status: true,
			Msg:    "Failed to fetch data",
			Code:   10001,
		}
	case strings.Contains(err, "10002"):
		errResp = response.Error{
			Status: true,
			Msg:    "Failed to insert data",
			Code:   10001,
		}
	case strings.Contains(err, "1146"):
		errResp = response.Error{
			Status: true,
			Msg:    "Table Dosn't Exsist",
			Code:   1146,
		}
	default:
		errResp = response.Error{
			Status: true,
			Msg:    err,
			Code:   0,
		}
	}

	return response.Response{
		Error: errResp,
	}
}
