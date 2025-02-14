package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
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
	query := `UPDATE sessions 
    SET token = ?, expires_at = datetime('now', '+1 hour')
    WHERE user_id = (SELECT id FROM users WHERE username = ?)`

	_, err = forumDB.Exec(query, token.String(), strings.ToLower(potentialUser.Username))
	if err != nil {
		log.Error("internal server error: ", err)
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(conn.Resp, &cookie)

	conn.Resp.WriteHeader(http.StatusOK)
	conn.Resp.Write([]byte(`{"message": "Login successful"}`))
	return
}
