package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"npm/api/service"
	"npm/pkg/models"
)

func UpdateUser(c echo.Context) error {

	// paramter passed by the user with termed id
	key := c.QueryParam("id")

	var reqBody models.User
	err1 := c.Bind(&reqBody) // binding the data(sent by user) with reqBody
	if err1 != nil {
		return err1
	}

	var v = validator.New()
	err2 := v.Struct(&reqBody) // checking validation
	if err2 != nil {
		return err2
	}

	// updating the user  instance provided
	ans, err3 := service.UpdateOne(reqBody, key)
	if err3 != nil {
		return err3
	}
	return c.JSON(http.StatusOK, ans)
}
