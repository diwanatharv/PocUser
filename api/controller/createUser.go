package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"npm/api/service"
	"npm/pkg/models"
)

func CreateUser(c echo.Context) error {

	var reqBody models.User

	err := c.Bind(&reqBody) //whatever the data is coming, bind with reqBody
	if err != nil {
		return err
	}
	var v = validator.New()
	err2 := v.Struct(&reqBody) //checking validation
	if err2 != nil {
		return c.JSON(http.StatusNotFound, "validation match nhi huye")
	}

	//insert the data into the collection
	//res, err4 := collection.InsertOne(reqBody)
	res, err4 := service.Create(reqBody)

	if err4 != nil {
		return err4
	}
	return c.JSON(http.StatusOK, res)
}
