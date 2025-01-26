package config

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTables(db *sql.DB) error {
	statement, err := os.ReadFile("./app/config/schema.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(statement))
	if err != nil {
		return err
	}
	return nil
}
