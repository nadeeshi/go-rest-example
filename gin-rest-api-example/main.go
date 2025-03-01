package main

import (
	_ "example.com/gin-rest-api-example/docs" // Import generated docs
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	service "example.com/gin-rest-api-example/service" // Import the service package
)

// @title Gin REST API Example
// @version 1.0
// @description This is a sample Gin REST API with Swagger documentation
// @host localhost:8080
// @BasePath /api/v1
func main() { // Entry point
	router := gin.Default()

	/*router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from API 1!"})
	})*/

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", service.GetUsers)
		v1.GET("/users/:id", service.GetUserByID)
		v1.POST("/users", service.CreateUser)
		v1.PUT("/users/:id", service.UpdateUser)
		v1.DELETE("/users/:id", service.DeleteUser)
	}

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	router.Run(":8080") // API 1 runs on port 8081
}
