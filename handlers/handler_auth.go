package handlers

import (
	"database/sql"
	"go-gin-postgres-template/config"
	"go-gin-postgres-template/models"
	"net/http"

	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	// "os"
	"time"
)

var secretKey1 = []byte(config.GetJWTSecret())

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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return secretKey1, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		user_id := claims["sub"].(string)

		c.Set("user_id", user_id)
		c.Next()
	}
}
