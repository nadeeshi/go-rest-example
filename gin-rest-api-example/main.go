package main

import (
	"github.com/gin-gonic/gin"

	service "example.com/gin-rest-api-example/service" // Import the service package
)

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

	router.Run(":8080") // API 1 runs on port 8081
}
