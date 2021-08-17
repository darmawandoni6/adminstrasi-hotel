package routes

import (
	"administrasi-hotel/controlers/facilities"
	"administrasi-hotel/controlers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware        middleware.JWTConfig
	UserController       users.UserController
	FacilitiesController facilities.FacilitiesController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	jwt := middleware.JWTWithConfig(cl.JWTMiddleware)

	v1 := "api/v1"

	auth := e.Group(v1)
	auth.POST("/auth/register", cl.UserController.Register)
	auth.POST("/auth/login", cl.UserController.Login)

	user := e.Group(v1 + "/users")
	user.Use(jwt)
	user.GET("", cl.UserController.Find)
	user.GET("/profile", cl.UserController.Profile)
	user.GET("/id/:id", cl.UserController.FindById)
	user.POST("", cl.UserController.Register)
	user.PUT("/id/:id", cl.UserController.Update)
	user.DELETE("/id/:id", cl.UserController.Delete)

	user.POST("/dumy", cl.UserController.Dummy)

	facilities := e.Group(v1 + "/facilities")
	facilities.Use(jwt)
	facilities.GET("", cl.FacilitiesController.Find)
	facilities.GET("/id/:id", cl.FacilitiesController.FindById)
	facilities.POST("", cl.FacilitiesController.Create)
	facilities.PUT("/id/:id", cl.FacilitiesController.Update)
	facilities.DELETE("/id/:id", cl.FacilitiesController.Delete)

}
