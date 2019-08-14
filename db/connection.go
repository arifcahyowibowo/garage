package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Conn is Function to return DB Connection
func Conn() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:changeme@localhost/garages?sslmode=disable")
	if err != nil {
		fmt.Println(`Could not connect to db`)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(`Could not connect to db`)
	}
	return db
}
