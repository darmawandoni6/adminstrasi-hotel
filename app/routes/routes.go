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
	v1 := "api/v1"
	auth := e.Group(v1)
	auth.POST("/auth/register", cl.UserController.Register)
	auth.POST("/auth/login", cl.UserController.Login)

}
