package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"forum/app/modules"
	"io"
	"net/http"
)

func Register(dataReader io.ReadCloser, resp http.ResponseWriter) error {
	var potentialuser modules.User
	err := json.NewDecoder(dataReader).Decode(&potentialuser)
	if err != nil {
		return errors.New("invalid format")
	}
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return errors.New("internal server error")
	}
	defer db.Close()
	if err := potentialuser.ValidInfo(db); err != nil {
		return err
	}

	if err := potentialuser.CreateUser(db, resp); err != nil {
		return err
	}
	return nil
}
