package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDatabase() {

	LoadEnv()

	// Retrieve database configuration from environment variables
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Error converting port to integer: %v", err)
	}
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	// Prepare the PostgreSQL connection string
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)

	// Open the database connection
	db, err := sql.Open("postgres", psqlSetup)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	} else {
		Db = db
		fmt.Println("Successfully connected to the database!")
	}
}
