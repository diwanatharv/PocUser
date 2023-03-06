package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"npm/api/service"
)

func GetUser(c echo.Context) error {
	// will find the user with his unique field which is id in my case
	Id, err := strconv.Atoi(c.QueryParam("id"))

	if err != nil {
		return err
	}
	// this function helps to find the user with the particular id
	ans := service.FindOne(bson.M{"id": Id}, c.QueryParam("id"))
	return c.JSON(http.StatusOK, ans)
}
