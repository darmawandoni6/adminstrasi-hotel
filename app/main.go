package main

import (
	_userUsecase "administrasi-hotel/busieness/users"
	_userController "administrasi-hotel/controlers/users"
	_userRepo "administrasi-hotel/drivers/tables/users"

	_facilitiesUsecase "administrasi-hotel/busieness/facilities"
	_facilitiesController "administrasi-hotel/controlers/facilities"
	_facilitiesRepo "administrasi-hotel/drivers/tables/facilities"

	_typeRoomsUsecase "administrasi-hotel/busieness/typeRooms"
	_typeRoomsController "administrasi-hotel/controlers/typeRooms"
	_typeRoomsRepo "administrasi-hotel/drivers/tables/typeRooms"

	"time"

	"administrasi-hotel/drivers/database"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_middleware "administrasi-hotel/app/middlewares"
	_routes "administrasi-hotel/app/routes"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configDB := database.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()

	configJWT := _middleware.ConfigJWT{
		Secret:    viper.GetString(`jwt.secret`),
		ExpSecret: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	userRepo := _userRepo.UsersRepository(db)
	userUsecase := _userUsecase.UsersUsecase(timeoutContext, userRepo, &configJWT)
	userControler := _userController.NewUserController(userUsecase)

	facilitiesRepo := _facilitiesRepo.UsersRepository(db)
	facilitiesUsecase := _facilitiesUsecase.FacilitiesUsecase(timeoutContext, facilitiesRepo, &configJWT)
	facilitiesControler := _facilitiesController.NewFacilitiesController(facilitiesUsecase)

	typeRoomsRepo := _typeRoomsRepo.UsersRepository(db)
	typeRoomsUsecase := _typeRoomsUsecase.TypeRoomsUsecase(timeoutContext, typeRoomsRepo, &configJWT)
	typeRoomsController := _typeRoomsController.NewTypeRoomsController(typeRoomsUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:        configJWT.Init(),
		UserController:       *userControler,
		FacilitiesController: *facilitiesControler,
		TypeRoomsController:  *typeRoomsController,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
