package user

import (
	"database/sql"

	"forum/app/api/ws"
	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
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
	if !conn.IsAuthenticated(db) {
		return
	}
	var users []modules.User
	query := `SELECT id, username, profile_picture FROM users`

	rows, err := db.Query(query)
	if err != nil {
		log.Error("Error executing query:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user modules.User

		err := rows.Scan(&user.Id, &user.Username, &user.ProfilePicture)
		if err != nil {
			log.Error("Error scanning row:", err)
			continue
		}
		if ws.IsActive(user.Id) {
			user.Status = "online"
		} else {
			user.Status = "offline"
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Error("Rows error:", err)
		return
	}
	conn.Respond(users)
}
