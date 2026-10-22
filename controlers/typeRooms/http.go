package typeRooms

import (
	"administrasi-hotel/busieness/typeRooms"
	"administrasi-hotel/helpers/alert"
	"administrasi-hotel/helpers/baseResponse"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TypeRoomsController struct {
	typeRoomsUsecase typeRooms.Usecase
}

func NewTypeRoomsController(uc typeRooms.Usecase) *TypeRoomsController {
	return &TypeRoomsController{
		typeRoomsUsecase: uc,
	}
}

func (ctrl *TypeRoomsController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	req := ReqTypeRooms{}
	err := c.Bind(&req)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.typeRoomsUsecase.Create(ctx, req.ToDomain())

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessInsert, nil)
}
func (ctrl *TypeRoomsController) Find(c echo.Context) error {
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

	res, count, LastPage, err := ctrl.typeRoomsUsecase.Find(ctx, page, offset)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	paginition.CurrentPage = page
	paginition.LastPage = LastPage
	paginition.PerPage = offset
	paginition.Total = count

	resArr := []ResTypeRooms{}

	for _, value := range res {
		resArr = append(resArr, FromDomain(value))
	}
	return baseResponse.SuccessResponse(c, resArr, paginition)
}

func (ctrl *TypeRoomsController) FindById(c echo.Context) error {
	ctx := c.Request().Context()
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	res, err := ctrl.typeRoomsUsecase.FindById(ctx, id)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return baseResponse.SuccessResponse(c, FromDomain(res), nil)
}

func (ctrl *TypeRoomsController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	req := ReqTypeRooms{}
	paramId := c.Param("id")
	err := c.Bind(&req)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.typeRoomsUsecase.Update(ctx, id, req.ToDomain())

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessUpdate, nil)
}

func (ctrl *TypeRoomsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.typeRoomsUsecase.Delete(ctx, id)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessDelete, nil)
}
