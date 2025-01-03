package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	var err error

	//Opens the database connection
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Error ! Could not open connection to DB")
	}

	//Sets the maximum number of open connections to the database
	DB.SetMaxOpenConns(10)

	// Sets the maximum number of connections to be available in the connection pool when idle
	DB.SetMaxIdleConns(5)

	if err = DB.Ping(); err != nil {
		panic("Error! Could not establish a connection to the database: " + err.Error())
	}

	createTables()

}

func createTables() {

	createTableUsers :=
		` CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createTableUsers)

	if err != nil {
		panic("Sorry, could not create the users table: " + err.Error())

	}

	createTableEvent := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id))
	`
	_, err = DB.Exec(createTableEvent)

	if err != nil {
		panic("Sorry, could not create the event table: " + err.Error())

	}

	createTableRegistrations := `
	CREATE TABLE IF NOT EXISTS registrations(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		event_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(event_id) REFERENCES events(id))
	`

	_, err = DB.Exec(createTableRegistrations)

	if err != nil {
		panic("Sorry, could not create the registrations table: " + err.Error())

	}

}
