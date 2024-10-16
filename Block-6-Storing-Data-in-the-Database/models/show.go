package models

import (
	"api3/db"
	"time"
)

type Show struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string	`binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID int // user that created the show
}

var shows []Show = []Show{}


func (s Show) Save() error {
	// later: add to database
	query := `INSERT INTO shows(name, description, location, dateTime, user_id) VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err !=nil {
		return err
	}

	defer stmt.Close()
	results, err := stmt.Exec(s.Name, s.Description, s.Location, s.DateTime, s.UserID)

	if err !=nil {
		return err
	}

	id, err := results.LastInsertId()

	s.ID = id

	return err

	//shows = append(shows, s)
}

func GetAllShows() []Show {
	return shows
}
