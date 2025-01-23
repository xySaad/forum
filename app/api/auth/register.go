package auth

import (
	"database/sql"
	"encoding/json"
	"forum/app/modules"
	"forum/app/modules/errors"
)

func Register(conn *modules.Connection, forumDB *sql.DB) error {
	var potentialuser modules.AuthCredentials
	err := json.NewDecoder(conn.Req.Body).Decode(&potentialuser)
	if err != nil {
		conn.NewError(500, errors.CodeParsingError, "Internal Server Error", "Request is not valid JSON")
		return err
	}

	httpErr := potentialuser.ValidInfo(forumDB)
	if httpErr != nil {
		conn.Error(httpErr)
		return err
	}

	err = potentialuser.CreateUser(forumDB, conn.Resp)
	if err != nil {
		conn.NewError(500, errors.CodeUserCreationError, "User creation failed", "Unable to create user")
		return err
	}

	conn.Resp.Write([]byte("Registration successful"))
	return nil
}
