package modules

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username        string `json:"Username"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	PasswordConfirm string ` json:"ConfirmPassword"`
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
