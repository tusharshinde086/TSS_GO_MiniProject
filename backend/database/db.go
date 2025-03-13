package database

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is the database connection
var DB *sql.DB

// Connect initializes the database connection
func Connect() {
	var err error
	// Replace with your database connection string
	DB, err = sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
	if err != nil {
		panic(err)
	}
}
