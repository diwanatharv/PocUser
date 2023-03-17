package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"npm/api/service"
	"npm/pkg/models"
)

func GetAllUser(c echo.Context) error {
	// this is the  empty slice
	var allUser []models.User

	// this will find all those whose id is true which mean it exists
	elementFilter := bson.M{
		"id": bson.M{"$exists": true},
	}
	//function for finding all the user
	allUser = service.FindAll(elementFilter)

	return c.JSON(http.StatusOK, allUser)
}
