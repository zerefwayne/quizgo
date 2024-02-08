package database

import (
	"database/sql"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	// Open a database connection with quizgo.sqlite file
	db, err := sql.Open("sqlite3", "quizgo.sqlite")
	if err != nil {
		t.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Error pinging database: %v", err)
	}

	t.Log("Database Online")
}
