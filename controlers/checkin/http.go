package checkin

import (
	"administrasi-hotel/busieness/checkin"
	"administrasi-hotel/helpers/alert"
	"administrasi-hotel/helpers/baseResponse"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CheckinController struct {
	checkinUsecase checkin.Usecase
}

func NewCheckinController(uc checkin.Usecase) *CheckinController {
	return &CheckinController{
		checkinUsecase: uc,
	}
}

func (ctrl *CheckinController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	req := ReqCheckin{}
	err := c.Bind(&req)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.checkinUsecase.Create(ctx, req.ToDomain())

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessInsert, nil)
}

func (ctrl *CheckinController) Find(c echo.Context) error {
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

	res, count, LastPage, err := ctrl.checkinUsecase.Find(ctx, page, offset)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	paginition.CurrentPage = page
	paginition.LastPage = LastPage
	paginition.PerPage = offset
	paginition.Total = count

	resArr := []ResCheckin{}

	for _, value := range res {
		resArr = append(resArr, FromDomain(value))
	}
	return baseResponse.SuccessResponse(c, resArr, paginition)
}

func (ctrl *CheckinController) FindById(c echo.Context) error {
	ctx := c.Request().Context()
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	res, err := ctrl.checkinUsecase.FindById(ctx, id)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return baseResponse.SuccessResponse(c, FromDomain(res), nil)
}
