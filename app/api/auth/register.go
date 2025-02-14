package auth

import (
	"database/sql"
	"encoding/json"
	"strings"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
)

func Register(conn *modules.Connection, forumDB *sql.DB) {
	var potentialuser modules.AuthCredentials
	err := json.NewDecoder(conn.Req.Body).Decode(&potentialuser)
	if err != nil {
		conn.Error(errors.BadRequestError("Request is not valid JSON"))
		return
	}
	potentialuser.Username = strings.ToLower(potentialuser.Username)
	httpErr := potentialuser.ValidInfo(forumDB)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}

	err = potentialuser.CreateUser(forumDB, conn.Resp)
	if err != nil {
		log.Error(err)
		conn.Error(errors.HttpInternalServerError)
		return
	}

	conn.Resp.Write([]byte("Registration successful"))
}
