package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTables(db *sql.DB) {
	statement, err := os.ReadFile("./app/config/schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(string(statement))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Connected successfully!")
}
