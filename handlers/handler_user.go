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

func GetUser(c *gin.Context) {
	var user models.User
	username := c.Param("username") // Get username from the URL parameter

	// Query the user data from the database
	err := database.Db.QueryRow("SELECT user_id, username, password FROM users WHERE username = $1", username).
		Scan(&user.User_id, &user.Username, &user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	// Respond with the user data (omit password for security)
	c.JSON(http.StatusOK, gin.H{
		"id":       user.User_id,
		"username": user.Username,
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	// Query all users from the database
	rows, err := database.Db.Query("SELECT user_id, username FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users from database"})
		return
	}
	defer rows.Close()

	// Iterate over each row and scan the data into the users slice
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.User_id, &user.Username); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning user data"})
			return
		}
		users = append(users, user)
	}

	// Check for any error encountered during iteration
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing user data"})
		return
	}

	// Respond with the list of users
	c.JSON(http.StatusOK, gin.H{"users": users})
}
