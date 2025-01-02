package auth

import (
	"database/sql"
	"errors"
)

func CheckAuth(cookieValue string) error {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return errors.New("internal server error")
	}
	defer db.Close()
	err = db.QueryRow("SELECT (*) FROM users WHERE uuid=?", cookieValue).Scan()
	if err != nil {

	}
	return nil
}
