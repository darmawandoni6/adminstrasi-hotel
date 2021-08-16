package main

import (
	_userUsecase "administrasi-hotel/busieness/users"
	_userController "administrasi-hotel/controlers/users"
	_userRepo "administrasi-hotel/drivers/tables/users"
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
	userUsecase := _userUsecase.UsersUsecase(timeoutContext, userRepo)
	userControler := _userController.NewUserController(userUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:  configJWT.Init(),
		UserController: *userControler,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
