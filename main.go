package main

import (
	"npm/api/routes"
	"npm/pkg/config"
	"npm/pkg/dataAccess"
)

func main() {
	config.CreateRedisDatabase()
	config.CreateDatabase("UserDb")
	dataAccess.CreateUserCollection()
	routes.CreateRoutesAndServer()
}
