package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"unicode"
)

func IsAuth(r *http.Request) (is_auth bool) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return
	}
	uuid := cookie.Value
	if uuid == "" {
		return
	}

	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return
	}
	query := "SELECT  ( uuid) FROM users WHERE uuid= ( ?)"
	_, err = db.Exec(query, uuid)
	if err != nil {
		return
	}
	is_auth = true
	return
}

func VerifyName(name string) error {
	if len(name) < 4 || len(name) > 18 {
		return errors.New("name-lenght should be between 4 and 18")
	}
	for _, char := range name {
		if !unicode.IsGraphic(char) {
			return errors.New("name should only contain printable characters")
		}
	}
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return err
	}
	query := "SELECT  (username) FROM users WHERE username= ( ?)"
	err = db.QueryRow(query, name).Scan()
	if err != sql.ErrNoRows {
		return errors.New("name already exists")
	}
	return nil
}

func VerifyEmail(Email string) error {
	EmailParts := strings.Split(Email, "@")
	if len(EmailParts) != 2 {
		return errors.New("invalid email")
	}
	if len(EmailParts[0]) == 0 {
		return errors.New("invalid email")
	}
	EmailParts = strings.Split(EmailParts[1], ".")
	if len(EmailParts) != 2 {
		return errors.New("invalid email")
	}
	if len(EmailParts[0]) == 0 {
		return errors.New("invalid email")
	}
	for _, char := range Email {
		if !unicode.IsGraphic(char) {
			return errors.New("email should only contain printable characters")
		}
	}
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return err
	}

	query := "SELECT  (email) FROM users WHERE email= ( ?)"
	err = db.QueryRow(query, Email).Scan()
	if err != sql.ErrNoRows {
		return errors.New("email already exists")
	}
	return nil
}

func Verify(user User) error {
	err := VerifyName(user.Username)
	if err != nil {
		return err
	}
	err = VerifyPassword(user.Password)
	if err != nil {
		return err
	}
	err = VerifyEmail(user.Email)
	return err
}

func VerifyPassword(password string) error {
	if len(password) < 8 {
		return errors.New("password should be atleast 8 chars long")
	}
	return nil
}

func redirect(url, message string, status uint16, w http.ResponseWriter) (unexpectedErr error) {
	var response response
	response.Url = url
	response.Message = message
	jrsp, err := json.Marshal(response)
	if err != nil {
		// for the one that wills to create anm err page and handler please put it right here
		return err
		// maybe not here but whenever we call this function (sry ;) )
	}
	w.WriteHeader(int(status))
	w.Write(jrsp)
	return
}
