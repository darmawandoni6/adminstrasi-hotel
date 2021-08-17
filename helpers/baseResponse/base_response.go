package baseResponse

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}
type BaseResponseAuth struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
}

func SuccessResponseAuth(c echo.Context, data interface{}) error {
	response := BaseResponseAuth{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = data

	return c.JSON(http.StatusOK, response)
}

func SuccessResponse(c echo.Context, data interface{}, pagination interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = data
	response.Pagination = pagination

	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something not right"
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(status, response)
}
