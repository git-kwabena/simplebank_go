package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// NewMYSQLStorage Create database storage connection
func NewMYSQLStorage() *sql.DB {
	// Load in the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Open database nb: this does not connect to db
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	return db
}

// InitStorage Connect to database
func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
