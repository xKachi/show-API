package models

import "time"

type Show struct {
	ID          int
	Description string `binding:"required"`
	Location    string	`binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID int // user that created the show
}

var shows []Show = []Show{}

func (s Show) Save() {
	// later: add to database
	shows = append(shows, s)
}

func GetAllShows() []Show {
	return shows
}
