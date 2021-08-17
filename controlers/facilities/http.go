package facilities

import (
	"administrasi-hotel/busieness/facilities"
	"administrasi-hotel/helpers/alert"
	"administrasi-hotel/helpers/baseResponse"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FacilitiesController struct {
	facilitiesUsecase facilities.Usecase
}

func NewFacilitiesController(uc facilities.Usecase) *FacilitiesController {
	return &FacilitiesController{
		facilitiesUsecase: uc,
	}
}

func (ctrl *FacilitiesController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	req := ReqFacilities{}
	err := c.Bind(&req)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.facilitiesUsecase.Create(ctx, req.ToDomain())

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessInsert, nil)
}
