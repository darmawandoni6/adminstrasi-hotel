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
}

func (ctrl *UserController) Login(c echo.Context) error {

	ctx := c.Request().Context()
	req := ReqLogin{}
	err := c.Bind(&req)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	token, expired, err := ctrl.userUseCase.Login(ctx, req.Email, req.Password)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := struct {
		Token       string `json:"token"`
		ExpiredDate string `json:"expired_date"`
	}{Token: token, ExpiredDate: expired}

	return baseResponse.SuccessResponse(c, response, nil)
}
