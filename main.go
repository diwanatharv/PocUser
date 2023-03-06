package main

import (
	"npm/api/routes"
	"npm/pkg/config"
	"npm/pkg/dataAccess"
)

func main() {
	// intializes reddis database
	config.CreateRedisDatabase()
	config.CreateMongoDatabase("UserDb")
	dataAccess.CreateUserCollection("Usercool")
	routes.CreateRoutesAndServer()
}
