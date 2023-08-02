package httpservice

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseJSONResponse struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Latency string      `json:"latency,omitempty"`
}

type JSONResponse struct {
	BaseJSONResponse
	Data    interface{} `json:"data,omitempty"`
	HasNext *bool       `json:"has_next,omitempty"`
}

func NewJsonResponse() *JSONResponse {
	return &JSONResponse{
		BaseJSONResponse: BaseJSONResponse{Code: http.StatusOK},
	}
}

func (r *JSONResponse) SetData(data interface{}) *JSONResponse {
	r.Data = data
	return r
}

func (r *JSONResponse) SetMessage(msg string) *JSONResponse {
	r.Message = msg
	return r
}

func (r *JSONResponse) SetStatusCode(code int) *JSONResponse {
	r.Code = code
	return r
}

func (r *JSONResponse) SetLatency(latency float64) *JSONResponse {
	r.Latency = fmt.Sprintf("%.2f ms", latency)
	return r
}

func (r *JSONResponse) Send(e echo.Context) (err error) {
	return e.JSON(r.Code, r)
}
