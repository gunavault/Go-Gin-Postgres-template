package handlers

import (
	"database/sql"
	"my-portofolio/config"
	"my-portofolio/models"
	"net/http"

	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	// "os"
	"time"
)

// Example login handler
func Login(c *gin.Context) {
	var user models.User

	// initiate to usingENV
	config.LoadEnv()

	if err := c.ShouldBindJSON(&user); err != nil {
		switch {
		default:
			// Unexpected error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	// Fetch user from database based on Username (you can costume here base on your login method)
	row := config.Db.QueryRow("SELECT user_id, username, password FROM users WHERE username = $1", user.Username)
	var dbUser models.User
	err := row.Scan(&dbUser.User_id, &dbUser.Username, &dbUser.Password)
	if err != nil {
		// Error handler
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	// validate the password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Create JWT token (you can also make it as new function)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dbUser.User_id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	secretKey1 := []byte(config.GetJWTSecret())

	tokenString, err := token.SignedString(secretKey1)
	if err != nil {
		// Log the error with details
		log.Println("Failed to sign JWT token:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	// Respond with the token
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
