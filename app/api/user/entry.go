package user

import (
	"database/sql"

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
