package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitDB() {
	var DB *sql.DB

	DB, err := sql.Open("sqlite", "api.db")

	if err !=nil {
		panic("Could not initialize database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}