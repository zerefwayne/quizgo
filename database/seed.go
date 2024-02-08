package database

import (
	"database/sql"
	"fmt"
	"time"
)

func createTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			score INTEGER,
			last_game TIMESTAMP
		)
	`)
	return err
}

func createUser(db *sql.DB) error {
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE id = ?", 1).Scan(&userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		stmt, err := db.Prepare("INSERT INTO users (id, name, score, last_game) VALUES (?, ?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		name := ""
		fmt.Printf("Please, tell us your name: ")
		fmt.Scanf("%s", &name)

		_, err = stmt.Exec(1, name, 0, time.Now()) // Assuming user_id is 1
		if err != nil {
			return err
		}

		fmt.Println("User posted successfully!")
	}

	return nil
}

func Initialize(db *sql.DB) error {
	err := createTable(db)
	if err != nil {
		return err
	}

	err = createUser(db)
	if err != nil {
		return err
	}

	fmt.Println("Database initialized")
	return nil
}
