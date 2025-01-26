package auth

import (
	"database/sql"
	"encoding/json"

	"forum/app/modules"
	"forum/app/modules/errors"
)

func Register(conn *modules.Connection, forumDB *sql.DB) {
	var potentialuser modules.AuthCredentials
	err := json.NewDecoder(conn.Req.Body).Decode(&potentialuser)
	if err != nil {
		conn.NewError(500, errors.CodeParsingError, "Internal Server Error", "Request is not valid JSON")
		return
	}

	httpErr := potentialuser.ValidInfo(forumDB)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}

	httpErr = potentialuser.CreateUser(forumDB, conn.Resp)
	if err != nil {
		conn.Error(httpErr)
		return
	}

	conn.Resp.Write([]byte("Registration successful"))
}
