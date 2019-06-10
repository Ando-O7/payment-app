package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Conn : sql connection handler
var Conn *sql.DB

// NewSQLHandler : init sql handler
func init() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")

	dbconf := user + ":" + pass + "@/" + name
	conn, err := sql.Open("mysql", dbconf)
	if err != nil {
		panic(err.Error)
	}
	Conn = conn
}
