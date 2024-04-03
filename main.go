package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", database)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
}
