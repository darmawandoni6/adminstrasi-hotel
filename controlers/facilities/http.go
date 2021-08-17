package facilities

import (
	"administrasi-hotel/busieness/facilities"
	"administrasi-hotel/helpers/alert"
	"administrasi-hotel/helpers/baseResponse"
	"net/http"
	"strconv"

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
func (ctrl *FacilitiesController) Find(c echo.Context) error {
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

	res, count, LastPage, err := ctrl.facilitiesUsecase.Find(ctx, page, offset)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	paginition.CurrentPage = page
	paginition.LastPage = LastPage
	paginition.PerPage = offset
	paginition.Total = count

	resArr := []ResFacilities{}

	for _, value := range res {
		resArr = append(resArr, FromDomain(value))
	}
	return baseResponse.SuccessResponse(c, resArr, paginition)
}

func (ctrl *FacilitiesController) FindById(c echo.Context) error {

	return baseResponse.SuccessResponse(c, alert.SuccessInsert, nil)
}

func (ctrl *FacilitiesController) Update(c echo.Context) error {
	return baseResponse.SuccessResponse(c, alert.SuccessInsert, nil)
}

func (ctrl *FacilitiesController) Delete(c echo.Context) error {
	return baseResponse.SuccessResponse(c, alert.SuccessInsert, nil)
}
