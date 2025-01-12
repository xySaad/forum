package modules

import (
	"database/sql"
	"fmt"
	"forum/app/modules/errors"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username        string
	Email           string
	Password        string
	PasswordConfirm string
}

func (User *User) CheckAccount(db *sql.DB) error {
	hashedPassWord := ""
	err := db.QueryRow("SELECT (password) FROM users WHERE username=? AND email=? VALUES (?,?)", User.Username, User.Email).Scan(&hashedPassWord)
	if err != nil {
		if err == sql.ErrNoRows {
			// need to change
			// return errors.New("invalid dkxi rak tem")
		}
		// return errors.New("internal server error")
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassWord), []byte(User.Password))
	if err != nil {
		// ........
		// return errors.New("invalid okda")
	}
	return nil
}

func (User *User) CreateUser(db *sql.DB, resp http.ResponseWriter) error {
	hashedPassWord, err := bcrypt.GenerateFromPassword([]byte(User.Password), 12)
	if err != nil {
		return err
	}
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users (username,uuid,password,email) VALUES (? ,? ,? ,?)", User.Username, uuid, hashedPassWord, User.Email)
	if err != nil {
		return err
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

func (User *User) ValidInfo(db *sql.DB) (httpErr *errors.HttpError) {
	httpErr = &errors.HttpError{
		Status:  400,
		Code:    errors.CodeInvalidUsername,
		Message: "invalid username",
	}

	if !ValidUserName(User.Username) {
		httpErr.Details = "username contains invalid character"
		return
	}

	var err error
	defer func() {
		if httpErr != nil && httpErr.Status == 500 {
			fmt.Println(err)
		}
	}()

	err = db.QueryRow("SELECT * FROM users WHERE username=(?)", User.Username).Scan()
	if err != sql.ErrNoRows {
		if err != nil {
			httpErr.Status = 500
			httpErr.Code = errors.CodeInternalServerError
			httpErr.Message = "internal server error"
			return
		}

		httpErr.Details = "username already taken"
		return
	}

	httpErr.Code = errors.CodeInvalidEmail
	httpErr.Message = "invalid email"

	if !ValidEmail(User.Email) {
		httpErr.Details = "email contains invalid character"
		return
	}

	err = db.QueryRow("SELECT * FROM users WHERE email=(?)", User.Email).Scan()
	if err != sql.ErrNoRows {
		if err != nil {
			httpErr.Status = 500
			httpErr.Code = errors.CodeInternalServerError
			httpErr.Message = "internal server error"
			return
		}

		httpErr.Details = "email already in use"
		return
	}
	if !ValidPassword(User.Password) {
		httpErr.Details = "password contains invalid character"
		return
	}
	return nil
}
