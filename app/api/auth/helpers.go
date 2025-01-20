package auth

import (
	"database/sql"
	"errors"
)

func CheckAuth(cookieValue string) error {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return errors.New("Internal server error")
	}
	defer db.Close()

	var userID string
	err = db.QueryRow("SELECT id FROM users WHERE token=?", cookieValue).Scan(&userID)
	if err != nil {
		return errors.New("Unauthorized access")
	}

	return nil
}
