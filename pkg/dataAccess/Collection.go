package dataAccess

import (
	"fmt"

	"npm/pkg/config"
	mongodb "npm/pkg/dataAccess/mongo"
)

// collection can be differeent that 's why using this function to create the collection
var Collection mongodb.Collection

func CreateUserCollection() {
	Collection.Users = config.Db.Collection("Usercool")
	if Collection.Users == nil {
		fmt.Println("not initialised in the user ")
	}
}
