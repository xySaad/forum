package auth

import (
	"database/sql"
	"encoding/json"
	"forum/app/modules"
	"forum/app/modules/errors"
)

func Register(conn *modules.Connection) {
	var potentialuser modules.AuthCredentials
	err := json.NewDecoder(conn.Req.Body).Decode(&potentialuser)
	if err != nil {
		conn.NewError(500, errors.CodeParsingError, "Internal Server Error", "Request is not valid JSON")
		return
	}

	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		conn.NewError(500, errors.CodeInternalServerError, "Internal Server Error", "")
		return
	}
	defer db.Close()

	httpErr := potentialuser.ValidInfo(db)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}

	err = potentialuser.CreateUser(db, conn.Resp)
	if err != nil {
		conn.NewError(500, errors.CodeUserCreationError, "User creation failed", "Unable to create user")
		return
	}

	conn.Resp.Write([]byte("Registration successful"))
}
