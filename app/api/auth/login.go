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

func LogIn(dataReader io.ReadCloser, resp http.ResponseWriter, forumDB *sql.DB) error {
	var potentialUser modules.AuthCredentials
	if err := json.NewDecoder(dataReader).Decode(&potentialUser); err != nil {
		http.Error(resp, "Invalid request format", http.StatusBadRequest)
		return err
	}

	if potentialUser.Username == "" && potentialUser.Email == "" || potentialUser.Password == "" {
		http.Error(resp, "Username/Email and password are required", http.StatusBadRequest)
		return errors.New("missing required fields")
	}

	if err := potentialUser.CheckAccount(forumDB); err != nil {
		http.Error(resp, "Invalid username/email or password", http.StatusUnauthorized)
		return err
	}

	token, err := uuid.NewV7()
	if err != nil {
		http.Error(resp, "Internal server error", http.StatusInternalServerError)
		return err
	}

	_, err = forumDB.Exec("UPDATE users SET token = ? WHERE username = ? OR email = ?", token.String(), potentialUser.Username, potentialUser.Email)
	if err != nil {
		http.Error(resp, "Failed to save token to database", http.StatusInternalServerError)
		return err
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(resp, &cookie)

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"message": "Login successful"}`))
	return nil
}
