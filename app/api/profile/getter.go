package profile

import (
	"database/sql"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
)

type privateUser struct {
	modules.User
	Email string `json:"email"`
}

func GetUserData(conn *modules.Connection, forumDB *sql.DB) {
	if !conn.IsAuthenticated(forumDB) {
		return
	}
	user := privateUser{User: conn.User}

	query := `SELECT username,profile_picture,email FROM users WHERE id=?`
	err := forumDB.QueryRow(query, user.Id).Scan(&user.Username, &user.ProfilePicture, &user.Email)
	if err != nil {
		log.Error(err)
		conn.Error(errors.HttpInternalServerError)
		return

	}
	conn.Respond(user)
}
