package handlers

import (
	database "my-portofolio/config"
	"my-portofolio/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register user and store hashed password
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Insert the user into the database
	_, err = database.Db.Exec("INSERT INTO users(username, password) VALUES($1, $2)", user.Username, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting user into database"})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
