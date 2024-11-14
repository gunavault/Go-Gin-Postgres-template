package routes

import (
	"my-portofolio/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes
func SetupRoutes(router *gin.Engine) {

	// User-related routes
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
}
