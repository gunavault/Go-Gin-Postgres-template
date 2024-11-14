package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Handle all your env here in feature
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Example JWT Secret Key
func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET_KEY")
}
