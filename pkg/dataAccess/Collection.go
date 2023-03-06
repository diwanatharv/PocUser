package dataAccess

import (
	"fmt"

	"npm/pkg/config"
	mongodb "npm/pkg/dataAccess/mongo"
)

// we can insert different user in  the collection
var Collection mongodb.Collection

func CreateUserCollection(dbName string) {
	Collection.Users = config.Db.Collection(dbName) // making a collection in mongodb
	if Collection.Users == nil {
		fmt.Println("not initialised in the user ")
	}
}
