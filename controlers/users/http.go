package users

import (
	"administrasi-hotel/busieness/users"
	"administrasi-hotel/helpers/alert"
	"administrasi-hotel/helpers/baseResponse"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

// func (ctrl *UserController) Register(c echo.Context) error {
// ctx := c.Request().Context()

// 	// return controller.NewSuccessResponse(c, "Successfully inserted")
// 	return baseResponse.SuccessResponse(c, alert.SuccessInsert)
// }
func (ctrl *UserController) Register(c echo.Context) error {
	ctx := c.Request().Context()
	req := ReqUsers{}
	err := c.Bind(&req)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.userUseCase.Create(ctx, req.Email, req.ToDomain())

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessInsert, nil)
	// return baseResponse.SuccessResponse(c, req.Email, nil)

}
