package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	DB, err := sql.Open("sqlite", "api.db")

	if err !=nil {
		panic("Could not initialize database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createShow := `
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 description TEXT NOT NULL,
	 location TEXT NOT NULL,
	 dateTime DATETIME,
	 user_id INTEGER
	`

	_, err := DB.Exec(createShow)

	if err != nil {
		panic("The show table could not be created")
	}	
}

