package models

import (
	"api2/db"
	"api2/utils"
	"errors"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err !=nil {
		return err 
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userid, err := result.LastInsertId()

	u.ID = userid
	return err
}

func (u User) ConfirmCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&retrievedPassword)

	if err !=nil {
		return err
	}

	// check for password validity

	validPassword := utils.IsPasswordValid(u.Password, retrievedPassword)

	if !validPassword {
		return errors.New("invalid credentials");
	}

	return nil
}