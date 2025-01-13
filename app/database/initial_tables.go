package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTables(db)
}

func createTables(db *sql.DB) {
	statement, err := os.ReadFile("./app/config/schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(string(statement))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database and tables created successfully!")
}
