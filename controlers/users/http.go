package users

import (
	"administrasi-hotel/busieness/users"
	"administrasi-hotel/helpers/alert"
	"administrasi-hotel/helpers/baseResponse"
	"net/http"
	"strconv"

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

func (ctrl *UserController) Find(c echo.Context) error {
	ctx := c.Request().Context()
	page := c.QueryParam("page")
	offset := c.QueryParam("limit")
	paginition := baseResponse.Pagination{}

	p, err := strconv.Atoi(page)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	o, err := strconv.Atoi(offset)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	res, count, LastPage, err := ctrl.userUseCase.Find(ctx, p, o)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	paginition.CurrentPage = p
	paginition.LastPage = LastPage
	paginition.PerPage = o
	paginition.Total = count

	resArr := []ResUsers{}

	for _, value := range res {
		// resArr = append(responseController, response.FromDomain(value))
		resArr = append(resArr, FromDomain(value))
	}

	return baseResponse.SuccessResponse(c, resArr, nil)
}
