package models

import "time"

type Show struct {
	ID          int
	Description string
	Location    string
	DateTime    time.Time
	UserID int // user that created the show
}

var shows []Show = []Show{}

func (s Show) Save() {
	// later: add to database
	shows = append(shows, s)
}