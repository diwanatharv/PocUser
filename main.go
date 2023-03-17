package main

import (
	"npm/api/logger"
	"npm/api/routes"
	"npm/pkg/config"
)

func main() {
	// intializes reddis database
	config.CreateRedisDatabase()
	config.CreateMongoDatabase("UserDb")
	//dataAccess.CreateUserCollection("Usercool")
	routes.CreateRoutesAndServer()
	logger.Closeloggerfile()
}
