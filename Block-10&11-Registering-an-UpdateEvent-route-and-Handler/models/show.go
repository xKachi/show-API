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

func GetAllShows() ([]Show, error) {
	query := "SELECT * FROM  shows"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var shows []Show

	for rows.Next() {
		var show Show

		err := rows.Scan(&show.ID, &show.Name, &show.Description, &show.Location, &show.DateTime, &show.UserID)
		
		defer rows.Close()
		if err != nil {
			return nil, err
		}

		shows = append(shows, show)
	}

	return shows, nil
}

func GetShowID(showid int64) (*Show, error) {
	query := "SELECT * FROM shows WHERE id = ?"
	row := db.DB.QueryRow(query, showid)

	var show Show
	err := row.Scan(&show.ID, &show.Name, &show.Description, &show.Location, &show.DateTime, &show.UserID)

	if err != nil {
		return nil, err
	}

	return &show, nil
}

func (show Show) Update () error {
	query := `
		UPDATE shows
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_ , err = stmt.Exec(show.Name, show.Description, show.Location, show.DateTime, show.ID)

	return err
}
