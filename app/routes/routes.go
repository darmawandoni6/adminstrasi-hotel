package routes

import (
	"administrasi-hotel/controlers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	route := e.Group("api/v1")
	route.POST("/auth/register", cl.UserController.Register)

}
