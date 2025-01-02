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

func RegisterUSer(dataReader io.ReadCloser, resp http.ResponseWriter) error {
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
	if err := potentialuser.ValidInfo(db); err != nil {
		return err
	}

	if err := potentialuser.CreateUser(db, resp); err != nil {
		return err
	}
	return nil
}

func (User *User) CreateUser(db *sql.DB, resp http.ResponseWriter) error {
	hashedPassWord, err := bcrypt.GenerateFromPassword([]byte(User.Password), 12)
	if err != nil {
		return errors.New("internal server error")
	}
	uuid, err := uuid.NewV7()
	if err != nil {
		return errors.New("internal server error")
	}
	_, err = db.Exec("INSERT INTO users (username,uuid,password,email) VALUES (? ,? ,? ,?)", User.Username, uuid, hashedPassWord, User.Email)
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

func (User *User) ValidInfo(db *sql.DB) error {
	err := ValidUserName(User.Username)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT (*) FROM users WHERE username=(?)", User.Username).Scan()
	if err != sql.ErrNoRows {
		if err != nil {
			return errors.New("internal server error")
		}
		return errors.New("user already exists")
	}
	err = ValidEmail(User.Email)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT (*) FROM users WHERE email=(?)", User.Email).Scan()
	if err != sql.ErrNoRows {
		if err != nil {
			return errors.New("internal server error")
		}
		return errors.New("email taken")
	}
	return ValidPassword(User.Password, User.PasswordConfirm)
}
