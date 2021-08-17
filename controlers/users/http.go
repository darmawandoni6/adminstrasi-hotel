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
	qureyPage := c.QueryParam("page")
	qureyOffset := c.QueryParam("limit")
	paginition := baseResponse.Pagination{}

	page, err := strconv.Atoi(qureyPage)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	offset, err := strconv.Atoi(qureyOffset)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	res, count, LastPage, err := ctrl.userUseCase.Find(ctx, page, offset)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	paginition.CurrentPage = page
	paginition.LastPage = LastPage
	paginition.PerPage = offset
	paginition.Total = count

	resArr := []ResUsers{}

	for _, value := range res {
		resArr = append(resArr, FromDomain(value))
	}

	return baseResponse.SuccessResponse(c, resArr, paginition)
}

func (ctrl *UserController) FindById(c echo.Context) error {

	ctx := c.Request().Context()
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := ctrl.userUseCase.FindById(ctx, id)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)

	}

	return baseResponse.SuccessResponse(c, FromDomain(res), nil)
}

func (ctrl *UserController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	req := ReqUsers{}
	paramId := c.Param("id")
	err := c.Bind(&req)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.userUseCase.Update(ctx, id, req.ToDomain())

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessUpdate, nil)

}

func (ctrl *UserController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	req := ReqUsers{}
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.userUseCase.Delete(ctx, id, req.ToDomain())

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessDelete, nil)

}
