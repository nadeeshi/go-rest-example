package service

import (
	"net/http"

	data_types "example.com/gin-rest-api-example/data_types" // Import the data_types package
	storage "example.com/gin-rest-api-example/storage"       // Import the storage package

	"github.com/gin-gonic/gin"
)

// Get all users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, storage.GetUsers())
}

// Get user by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, found := storage.GetUserByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Create user
func CreateUser(c *gin.Context) {
	var newUser data_types.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	storage.AddUser(newUser)
	c.JSON(http.StatusCreated, newUser)
}

// Update user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser data_types.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if !storage.UpdateUser(id, updatedUser) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

// Delete user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if !storage.DeleteUser(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
