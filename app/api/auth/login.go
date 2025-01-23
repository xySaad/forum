package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"forum/app/modules"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func LogIn(conn *modules.Connection, forumDB *sql.DB) error {
	var potentialUser modules.AuthCredentials
	if err := json.NewDecoder(conn.Req.Body).Decode(&potentialUser); err != nil {
		http.Error(conn.Resp, "Invalid request format", http.StatusBadRequest)
		return err
	}

	if potentialUser.Username == "" && potentialUser.Email == "" || potentialUser.Password == "" {
		http.Error(conn.Resp, "Username/Email and password are required", http.StatusBadRequest)
		return errors.New("missing required fields")
	}

	if err := potentialUser.CheckAccount(forumDB); err != nil {
		http.Error(conn.Resp, "Invalid username/email or password", http.StatusUnauthorized)
		return err
	}

	token, err := uuid.NewV7()
	if err != nil {
		http.Error(conn.Resp, "Internal server error", http.StatusInternalServerError)
		return err
	}

	_, err = forumDB.Exec("UPDATE users SET token = ? WHERE username = ? OR email = ?", token.String(), potentialUser.Username, potentialUser.Email)
	if err != nil {
		http.Error(conn.Resp, "Failed to save token to database", http.StatusInternalServerError)
		return err
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
	return nil
}
