package routes

import (
	"go-gin-postgres-template/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes
func SetupRoutes(router *gin.Engine) {

	// User-related routes
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	auth := router.Group("/")
	auth.Use(handlers.AuthMiddleware())
	{
		auth.GET("/user/:username", handlers.GetUser)
		auth.GET("/users", handlers.GetAllUsers)
		auth.POST("/projects", handlers.PostProject)
		auth.PUT("/projects/:project_id", handlers.UpdateProject)
		auth.DELETE("/projects/:project_id", handlers.DeleteProject)
	}
	router.GET("/projects", handlers.GetAllProject)

}
