package repositories

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Score    int       `json:"score"`
	LastGame time.Time `json:"last_game"`
}

// Save this method, could be used later
func Get(db *sql.DB, userID int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, score, last_game FROM users WHERE id = ?", userID).
		Scan(&user.ID, &user.Name, &user.Score, &user.LastGame)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func Update(db *sql.DB, score int, userID int) error {
	stmt, err := db.Prepare("UPDATE users SET score = ?, last_game = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(score, time.Now(), userID)
	if err != nil {
		return err
	}

	fmt.Println("Score updated successfully!")
	return nil
}

const (
	selectScoreQuery = "SELECT name, score FROM users WHERE id = ?"
)
