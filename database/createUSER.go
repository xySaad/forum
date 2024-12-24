package forum

import (
	"database/sql"
	"errors"
	"fmt"
)

func CreateUser(username, email, password string) error {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return errors.New("probleme accesing database")
	}
	defer db.Close()
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	_, err = db.Exec(query, username, email, password)
	if err != nil {
		fmt.Println(err)
		return errors.New("user already exists")
	}

	return nil
}
