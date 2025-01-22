package auth

import (
	"database/sql"
	"errors"
)

func CheckAuth(cookieValue string, forumDB *sql.DB) error {

	var userID string
	err := forumDB.QueryRow("SELECT id FROM users WHERE token=?", cookieValue).Scan(&userID)
	if err != nil {
		return errors.New("Unauthorized access")
	}

	return nil
}
