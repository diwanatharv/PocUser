package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"npm/api/controller"
)

func CreateRoutesAndServer() {

	var e = echo.New()
	e.Use(middleware.BasicAuth(func(userName string, password string, c echo.Context) (bool, error) {
		if userName == "atharv" && password == "ath123" {
			return true, nil
		} else {
			return false, nil
		}
	}))
	//apis
	e.GET("/users", controller.GetAllUser)
	e.GET("/user", controller.GetUser)
	e.POST("/create", controller.CreateUser)
	e.PUT("/user", controller.UpdateUser)

	e.Start(":7000")
	fmt.Println("printing after routing")
}
