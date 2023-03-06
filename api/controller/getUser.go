package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"npm/api/service"
)

func GetUser(c echo.Context) error {
	Id, err := strconv.Atoi(c.QueryParam("id")) // accessing the query param and converting to int

	if err != nil {
		return err
	}
	ans := service.FindOne(bson.M{"id": Id}, c.QueryParam("id"))
	return c.JSON(http.StatusOK, ans)
}
