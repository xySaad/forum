package db

import (
	"database/sql"
	"log"
)

func CreateUser(username, email, password string) error {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	_, err = db.Exec(query, username, email, password)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
