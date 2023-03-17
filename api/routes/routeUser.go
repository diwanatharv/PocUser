package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"npm/api/controller"
	"npm/api/logger"
	"npm/environment"
)

func CreateRoutesAndServer() {

	var e = echo.New()
	// basic authentication which includes name and password for all the methods
	Logfile := logger.Createloggerfile()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: Logfile,
	}))
	e.Use(middleware.BasicAuth(func(userName string, password string, c echo.Context) (bool, error) {
		id, pass := environment.EnvVariable(userName, password)
		if userName == id && password == pass {
			return true, nil
		} else {
			return false, nil
		}
	}))
	e.GET("/users", controller.GetAllUser)
	e.GET("/user", controller.GetUser)
	e.POST("/create", controller.CreateUser)
	e.PUT("/user", controller.UpdateUser)

	e.Logger.Fatal(e.Start(":7000"))
}
