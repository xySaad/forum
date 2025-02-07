package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"

	"github.com/gofrs/uuid"
)

func LogIn(conn *modules.Connection, forumDB *sql.DB) {
	var potentialUser modules.AuthCredentials
	if err := json.NewDecoder(conn.Req.Body).Decode(&potentialUser); err != nil {
		conn.NewError(http.StatusBadRequest, 405, "invalid format", "")
		return
	}

	if potentialUser.Username == "" && potentialUser.Email == "" || potentialUser.Password == "" {
		conn.NewError(http.StatusBadRequest, 405, "missing required fields", "")
		return
	}

	if err := potentialUser.VerifyPassword(forumDB); err != nil {
		conn.Error(err)
		return
	}
	token, err := uuid.NewV7()
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		return
	}

	_, err = forumDB.Exec("UPDATE users SET token = ? WHERE username = ? OR email = ?", token.String(), potentialUser.Username, potentialUser.Email)
	if err != nil {
		log.Error("internal server error: ", err)
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(conn.Resp, &cookie)

	conn.Resp.WriteHeader(http.StatusOK)
	conn.Resp.Write([]byte(`{"message": "Login successful"}`))
	return
}
