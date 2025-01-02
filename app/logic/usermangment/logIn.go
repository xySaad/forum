package usermangment

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LogIn(dataReader io.ReadCloser, resp http.ResponseWriter) error {
	var potentialuser User
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

func (User *User) CheckAccount(db *sql.DB) error {
	hashedPassWord := ""
	err := db.QueryRow("SELECT (password) FROM users WHERE username=? AND email=? VALUES (?,?)", User.Username, User.Email).Scan(&hashedPassWord)
	if err != nil {
		if err == sql.ErrNoRows {
			// need to change
			return errors.New("invalid dkxi rak tem")
		}
		return errors.New("internal server error")
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassWord), []byte(User.Password))
	if err != nil {
		// ........
		return errors.New("invalid okda")
	}
	return nil
}
