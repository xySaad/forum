package modules

import (
	"database/sql"
	"net/http"
	"regexp"
	"time"

	"forum/app/modules/errors"
	"forum/app/modules/log"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ValidUserName(name string) bool {
	for _, char := range name {
		if !(char < 127 && char > 32) {
			return false
		}
	}
	return true
}

func ValidEmail(email string) bool {
	valid := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return valid.MatchString(email)
}

func ValidPassword(password string) bool {
	return true
}

type AuthCredentials struct {
	Username string
	Email    string
	Password string
}

func (User *AuthCredentials) CheckAccount(db *sql.DB) *errors.HttpError {
	hashedPassWord := ""

	err := db.QueryRow("SELECT (password) FROM users WHERE username=? OR email=? VALUES (?,?)", User.Username, User.Email).Scan(&hashedPassWord)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.NewError(404, 404, "account not found", "")
			// need to change
		}
		// return errors.New("internal server error")
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassWord), []byte(User.Password))
	if err != nil {
		return errors.NewError(403, 403, "wrong paasword", "")

		// ........
	}
	return nil
}

func (User *AuthCredentials) CreateUser(db *sql.DB, resp http.ResponseWriter) *errors.HttpError {
	hashedPassWord, err := bcrypt.GenerateFromPassword([]byte(User.Password), 12)
	if err != nil {
		return errors.NewError(500, 500, "internal server error", "")
	}
	token, err := uuid.NewV7()
	if err != nil {
		return errors.NewError(500, 500, "internal server error", "")
	}
	id, err := uuid.NewV6()
	if err != nil {
		return errors.NewError(500, 500, "internal server error", "")
	}
	_, err = db.Exec("INSERT INTO users (id,username,token,password,email) VALUES (?, ? ,? ,? ,?)", id.String(), User.Username, token, hashedPassWord, User.Email)
	if err != nil {
		//handle later
	}
	cookie := http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true, // Makes the cookie inaccessible to JavaScript
		Path:     "/",
	}
	http.SetCookie(resp, &cookie)
	return nil
}

func (User *AuthCredentials) ValidInfo(db *sql.DB) (httpErr *errors.HttpError) {
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
			log.Error("internal server error: " + err.Error())
		}
	}()

	err = db.QueryRow("SELECT * FROM users WHERE username=(?)", User.Username).Scan()
	if err != sql.ErrNoRows {
		if err != nil {
			httpErr.Status = 500
			log.Error("internal server error", err)
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
			log.Error("internal server error", err)
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
