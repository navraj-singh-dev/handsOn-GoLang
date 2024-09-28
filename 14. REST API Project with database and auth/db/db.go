package db

import (
	"database/sql"

	// Note:
	//	'sqlite3' will not be used directly by me, as it will be used behind the scenes by database/sql
	//	package. But still i need it, so i will add _ in front of sqlite3 import statement to tell go
	//	that i need this package but it will not be used directly
	_ "github.com/mattn/go-sqlite3"
)

// global variable for the database
var DB *sql.DB

// opens a connection to the local database
func InitDB() {
	// sqlite3 will be used as driver, api.db file is local.. which will act as database for us
	// if api.db is not found then it will be created
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("cannot connect with database...")
	}

	// keep open 10 connection to this database simultaneously
	// upper bound
	DB.SetMaxOpenConns(10)
	// if database is idle then keep 5 connection opened
	// lower bound
	DB.SetMaxIdleConns(5)
}

