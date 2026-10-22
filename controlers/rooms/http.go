package rooms

import (
	"administrasi-hotel/busieness/rooms"
	"administrasi-hotel/helpers/alert"
	"administrasi-hotel/helpers/baseResponse"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomsController struct {
	roomsUsecase rooms.Usecase
}

func NewRoomsController(uc rooms.Usecase) *RoomsController {
	return &RoomsController{
		roomsUsecase: uc,
	}
}

func (ctrl *RoomsController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	req := ReqRooms{}
	err := c.Bind(&req)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.roomsUsecase.Create(ctx, req.ToDomain())

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessInsert, nil)
}

func (ctrl *RoomsController) Find(c echo.Context) error {
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

	res, count, LastPage, err := ctrl.roomsUsecase.Find(ctx, page, offset)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	paginition.CurrentPage = page
	paginition.LastPage = LastPage
	paginition.PerPage = offset
	paginition.Total = count

	resArr := []ResRooms{}

	for _, value := range res {
		resArr = append(resArr, FromDomain(value))
	}
	return baseResponse.SuccessResponse(c, resArr, paginition)
}

func (ctrl *RoomsController) FindById(c echo.Context) error {
	ctx := c.Request().Context()
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}
	res, err := ctrl.roomsUsecase.FindById(ctx, id)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return baseResponse.SuccessResponse(c, FromDomain(res), nil)
}

func (ctrl *RoomsController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	req := ReqRooms{}
	paramId := c.Param("id")
	err := c.Bind(&req)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.roomsUsecase.Update(ctx, id, req.ToDomain())

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessUpdate, nil)
}

func (ctrl *RoomsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = ctrl.roomsUsecase.Delete(ctx, id)

	if err != nil {
		return baseResponse.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return baseResponse.SuccessResponse(c, alert.SuccessDelete, nil)
}
