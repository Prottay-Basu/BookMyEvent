package models

import (
	"errors"
	"prottaybasu/book-my-event/db"
	"prottaybasu/book-my-event/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {

	query := `INSERT INTO users(email, password) VALUES(?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashedPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {

	query := `SELECT id, password FROM users where email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !isPasswordValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
