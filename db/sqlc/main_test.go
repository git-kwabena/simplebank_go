package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	//database, err := db.NewMYSQLStorage()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//db.InitStorage(database)

	// Loaded .env with specific path
	// Since default .load could not find the .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	//Open database nb: this does not connect to db
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}
	testQueries = New(db)
	os.Exit(m.Run())
}
