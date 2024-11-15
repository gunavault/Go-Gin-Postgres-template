package main

import (
	"go-gin-postgres-template/config"
	"go-gin-postgres-template/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// Connect to the database
	config.ConnectDatabase()

	// Set up routes
	routes.SetupRoutes(route)

	// Start the server
	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}
