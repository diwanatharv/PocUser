package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"npm/api/service"
	"npm/pkg/models"
)

func CreateUser(c echo.Context) error {
	// creating an instance of the user struct
	var reqBody models.User

	err := c.Bind(&reqBody) // whatever the data is coming we will bind it with the reqbody
	if err != nil {
		return err
	}
	var v = validator.New()
	err2 := v.Struct(&reqBody) // checking validation added in the model of the structure of the user
	if err2 != nil {
		return c.JSON(http.StatusNotFound, "validation failed")
	}
	// this will insert in the mongodb
	res, err4 := service.Create(reqBody)

	if err4 != nil {
		return err4
	}
	// return res if no error
	return c.JSON(http.StatusOK, res)
}
