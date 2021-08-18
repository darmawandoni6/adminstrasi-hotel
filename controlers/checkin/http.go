package checkin

import (
	"administrasi-hotel/busieness/checkin"
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

	return baseResponse.SuccessResponse(c, req, nil)
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
