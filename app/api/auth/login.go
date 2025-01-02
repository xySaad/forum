package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"forum/app/modules"
	"io"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func LogIn(dataReader io.ReadCloser, resp http.ResponseWriter) error {
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
	if err := potentialuser.CheckAccount(db); err != nil {
		return err
	}
	uuid, err := uuid.NewV7()
	if err != nil {
		return errors.New("internal server error")
	}
	cookie := http.Cookie{
		Name:     "ticket",
		Value:    uuid.String(),
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true, // Makes the cookie inaccessible to JavaScript
		Path:     "/",
	}
	http.SetCookie(resp, &cookie)
	return nil
}
