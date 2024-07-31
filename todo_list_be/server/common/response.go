package response

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

type BaseResponse[D any] struct {
	Code    codes.Code `json:"code"`
	Message string     `json:"message"`
	Data    *D         `json:"data"`
}

func (r *BaseResponse[D]) Render(_ http.ResponseWriter, req *http.Request) error {
	return nil
}
