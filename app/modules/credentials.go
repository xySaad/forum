package modules

import (
	"database/sql"
	"net/http"
	"regexp"
	"time"

	"forum/app/modules/errors"
	"forum/app/modules/snowflake"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ValidUserNameSyntax(name string) bool {
	if len(name) > 20 || len(name) < 1 {
		return false
	}
	//{  || char <= '0' && char >= '9'
	for _, char := range name {
		if !(char < 'a' || char >'z' || char < 'A' || char > 'Z' ) {
		// if char != '_' {
				return false
			// }
		}
	}
	return true
}

func ValidEmail(email string) bool {
	valid := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return valid.MatchString(email)
}

func ValidPassword(password string) bool {
	thereIsCap := false
	ThereIsNumbers := false
	ThereIsLower := false
	if len(password) <= 6 {
		return false
	} else {
		for _, char := range password {
			if char >= 'a' && char <= 'z' {
				ThereIsLower = true
			}
			if char >= 'A' && char <= 'Z' {
				thereIsCap = true
			}
			if char >= '0' && char <= '9' {
				ThereIsNumbers = true
			}
		}
	}
	return thereIsCap && ThereIsLower && ThereIsNumbers
}

type AuthCredentials struct {
	Username string
	Email    string
	Password string
}

func (User *AuthCredentials) VerifyPassword(db *sql.DB) *errors.HttpError {
	hashedPassWord := ""

	err := db.QueryRow("SELECT password FROM users WHERE username=? ", User.Username).Scan(&hashedPassWord)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.HttpUnauthorized
		}
		return errors.HttpInternalServerError
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassWord), []byte(User.Password))
	if err != nil {
		return errors.HttpUnauthorized
	}
	return nil
}

func (User *AuthCredentials) CreateUser(db *sql.DB, resp http.ResponseWriter) error {
	hashedPassWord, err := bcrypt.GenerateFromPassword([]byte(User.Password), 12)
	if err != nil {
		return err
	}
	token, err := uuid.NewV4()
	if err != nil {
		return err
	}
	userId := snowflake.Default.Generate()

	_, err = db.Exec("INSERT INTO users (id,username,password,email) VALUES (? ,? ,? ,?)", userId, User.Username, hashedPassWord, User.Email)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO sessions (user_id,token,expires_at) VALUES (?, ? ,datetime('now', '+1 hour'))", userId, token.String())
	if err != nil {
		return err
	}
	cookie := http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(resp, &cookie)
	return nil
}

func (User *AuthCredentials) ValidInfo(db *sql.DB) (httpErr *errors.HttpError) {
	httpErr = &errors.HttpError{
		Status:  http.StatusBadRequest,
		Code:    errors.CodeInvalidUsername,
		Message: "invalid username",
	}
	// check username
	if !ValidUserNameSyntax(User.Username) {
		httpErr.Details = "username should only contains (a-zA-Z0-9) and length between 1-20"
		return
	}
	exists := false
	err := db.QueryRow("SELECT 1 FROM users WHERE username = ?", User.Username).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return errors.HttpInternalServerError
	}

	if exists {
		httpErr.Details = "username already taken"
		return
	}

	httpErr.Code = errors.CodeInvalidEmail
	httpErr.Message = "invalid email"

	if !ValidEmail(User.Email) {
		httpErr.Details = "email contains invalid character"
		return
	}

	err = db.QueryRow("SELECT 1 FROM users WHERE email = ?", User.Email).Scan(&exists)
	if err != sql.ErrNoRows {
		if err != nil {
			return errors.HttpInternalServerError
		}
		httpErr.Details = "email already in use"
		return
	}
	if !ValidPassword(User.Password) {
		httpErr.Message = "invalid password"
		httpErr.Code = errors.CodeIncorrectPassword
		httpErr.Details = "password should contain at least 6 character one upper, one lower and one digit"
		return
	}
	return nil
}
