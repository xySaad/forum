package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/app/modules"
	"forum/app/modules/errors"
)

func Register(conn *modules.Connection) {
	var potentialuser modules.User
	err := json.NewDecoder(conn.Req.Body).Decode(&potentialuser)
	if err != nil {
		conn.NewError(500, errors.CodeParsingError, "internal server error", "request is not a valid JSON")
		return
	}
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		conn.NewError(500, errors.CodeInternalServerError, "internal server error", "")
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
		fmt.Println(err)
		conn.NewError(500, errors.CodeUserCreationError, "can't register", "cannot create that specific user")
		return
	}
	conn.Resp.Write([]byte("success"))
}
