package main

import (
	"baobab/internal/config"
	"baobab/internal/echohandler"
	"baobab/internal/preserver"
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		errMessage := fmt.Sprintf("Error : %v", err)
		fmt.Println(errMessage)
		return
	}

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	err = InitializeApp(cfg, e)
	if err != nil {
		errMessage := fmt.Sprintf("Error : %v", err)
		e.Logger.Error(errMessage)
		return
	}

	err = e.Start(":8080")
	if err != nil {
		errMessage := fmt.Sprintf("Error : %v", err)
		e.Logger.Error(errMessage)
	}
}

func InitializeApp(cfg *config.Config, e *echo.Echo) error {
	db, err := preserver.NewDB(cfg)
	if err != nil {
		return err
	}

	userRepository := preserver.NewUserRepository(db)
	userHandler := echohandler.NewUserHandler(userRepository)
	userHandler.RegisterRoutes(e)

	placeRepository := preserver.NewPlaceRepository(db)
	placeHandler := echohandler.NewPlaceHandler(placeRepository)
	placeHandler.RegisterRoutes(e)

	return nil
}
