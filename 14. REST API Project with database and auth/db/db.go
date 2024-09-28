package db

import (
	"database/sql"
	// sqlite3 is being used behind the scenes by sql package, it is not used directly in code
	// _ tells go that this package needs to be imported by not used directly
	_ "github.com/mattn/go-sqlite3"
)

// global variable for the database
var DB *sql.DB

// opens a connection to the local database
func InitDB() {
	var err error
	// sqlite3 will be used as driver, api.db file is local.. which will act as database for us
	// if api.db is not found then it will be created
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("cannot connect to database...")
	}

	// Set database connection pool parameters
	DB.SetMaxOpenConns(10) // upper bound of connections
	DB.SetMaxIdleConns(5)  // lower bound of connections

	// create the tables
	createTables()
}

func createTables() {
	// sql query to create an users table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("cannot create 'users' table...")
	}

	// sql query to create an events table
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	// using global DB variable to execute this sql query and create the table
	_, err = DB.Exec(createEventTable)

	if err != nil {
		panic("cannot create 'events' table...")
	}
}
