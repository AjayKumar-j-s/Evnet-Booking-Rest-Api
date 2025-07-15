package db

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3","api.db")
	if err != nil {
		panic("Could not connect to DB: " + err.Error())
	} 

	// Connection pool configuration
	DB.SetMaxOpenConns(10)              // Max 10 total open DB connections (active + idle)
	DB.SetMaxIdleConns(5)               // Max 5 idle connections ready for reuse

	// Ping to ensure the database is reachable
	// err = DB.Ping()
	// if err != nil {
	// 	panic("Failed to ping database: " + err.Error())
	// }

	// Create tables if not exist
	createTable()


	//create User 
}


func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS User (
	UserID INTEGER PRIMARY KEY AUTOINCREMENT,
	Email TEXT NOT NULL UNIQUE ,
	Password TEXT NOT NULL

	)`

	_,err := DB.Exec(query)

	if(err != nil){
		panic("Could not Create User")
	}


	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		Description TEXT NOT NULL,
		Location TEXT NOT NULL,
		DateTime DATETIME NOT NULL,
		UserID INTEGER,
		FOREIGN KEY (UserID) REFERENCES User(UserID)
	);`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table: " + err.Error())
	} else {
		fmt.Println("Events table created or already exists.")
	}
}
