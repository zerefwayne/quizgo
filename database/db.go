package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	// Open a database connection
	db, err := sql.Open("sqlite3", "quizgo.sqlite")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Test database
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	Initialize(db)
	fmt.Println("Connected to the database!")
	return db, nil

}
