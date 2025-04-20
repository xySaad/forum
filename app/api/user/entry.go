package user

import (
	"database/sql"
	"log"

	"forum/app/modules"
	"forum/app/modules/errors"
)

func Entry(conn *modules.Connection, db *sql.DB) {
	if len(conn.Path) != 3 {
		conn.Error(errors.HttpNotFound)
		return
	}
	switch conn.Path[2] {
	case "liked":
		GetLikedPosts(conn, db)
	case "created":
		GetUserCreatedPosts(conn, db)
	}
}

func GetAllUsers(conn *modules.Connection, db *sql.DB) {
	// efer db.Close()
	var users []string

	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	users = append(users, names...)
	conn.Respond(users)
}
