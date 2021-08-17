package routes

import (
	"administrasi-hotel/controlers/facilities"
	"administrasi-hotel/controlers/typeRooms"
	"administrasi-hotel/controlers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware        middleware.JWTConfig
	UserController       users.UserController
	FacilitiesController facilities.FacilitiesController
	TypeRoomsController  typeRooms.TypeRoomsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	jwt := middleware.JWTWithConfig(cl.JWTMiddleware)

	v1 := e.Group("api/v1")

	auth := v1.Group("/auth")
	auth.POST("/register", cl.UserController.Register)
	auth.POST("/login", cl.UserController.Login)

	user := v1.Group("/users")
	user.Use(jwt)
	user.GET("", cl.UserController.Find)
	user.GET("/profile", cl.UserController.Profile)
	user.GET("/id/:id", cl.UserController.FindById)
	user.POST("", cl.UserController.Register)
	user.PUT("/id/:id", cl.UserController.Update)
	user.DELETE("/id/:id", cl.UserController.Delete)

	user.POST("/dumy", cl.UserController.Dummy)

	facilities := v1.Group("/facilities")
	facilities.Use(jwt)
	facilities.GET("", cl.FacilitiesController.Find)
	facilities.GET("/id/:id", cl.FacilitiesController.FindById)
	facilities.POST("", cl.FacilitiesController.Create)
	facilities.PUT("/id/:id", cl.FacilitiesController.Update)
	facilities.DELETE("/id/:id", cl.FacilitiesController.Delete)

	typeRooms := v1.Group("/type-rooms")
	typeRooms.Use(jwt)
	typeRooms.GET("", cl.TypeRoomsController.Find)
	typeRooms.GET("/id/:id", cl.TypeRoomsController.FindById)
	typeRooms.POST("", cl.TypeRoomsController.Create)
	typeRooms.PUT("/id/:id", cl.TypeRoomsController.Update)
	typeRooms.DELETE("/id/:id", cl.TypeRoomsController.Delete)

}
