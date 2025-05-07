package modules

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode"

	"forum/app/modules/errors"
	"forum/app/modules/log"
	"forum/app/modules/snowflake"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9.]+@([a-zA-Z0-9]+\.)+[a-zA-Z0-9]{2,24}$`)

func ValidFirstname(name string) (bool, string) {
	if len(name) < 3 || len(name) > 20 {
		return false, "fstname must be between 3 and 20 characters"
	}
	for _, char := range name {
		if !(char >= 'a' && char <= 'z') &&
			!(char >= 'A' && char <= 'Z') &&
			!(char >= '0' && char <= '9') {
			return false, "firstname must contain only alphanumeric characters "
		}
	}
	return true, ""
}
func ValidAge(name string) (bool, string) {
	for _, r := range name {
		if !unicode.IsDigit(r) {
			return false, "age not allowed"
		}
	}
	return true, ""
}
func ValidLastname(name string) (bool, string) {
	if len(name) < 3 || len(name) > 20 {
		return false, "lastname must be between 3 and 20 characters"
	}
	for _, char := range name {
		if !(char >= 'a' && char <= 'z') &&
			!(char >= 'A' && char <= 'Z') &&
			!(char >= '0' && char <= '9') {
			return false, "lastname must contain only alphanumeric characters"
		}
	}
	return true, ""
}
func ValidGender(name string) (bool, string) {
	if name != "male" && name != "female" && name != "other" && name != "prefer not to say" {
		return false, "gender not allowed"

	}
	return true, ""
}
func ValidUsername(name string) (bool, string) {
	if len(name) < 3 || len(name) > 20 {
		return false, "username must be between 3 and 20 characters"
	}
	for _, char := range name {
		if !(char >= 'a' && char <= 'z') &&
			!(char >= 'A' && char <= 'Z') &&
			!(char >= '0' && char <= '9') &&
			char != '.' {
			return false, "username must contain only alphanumeric characters and dot (.)"
		}
	}
	return true, ""
}

func ValidEmail(email string) bool {
	return emailPattern.MatchString(email)
}

func ValidPassword(password string) bool {
	if len(password) < 12 || len(password) > 256 {
		return false
	}

	return true
}

type AuthCredentials struct {
	Username  string
	Age       string
	Gender    string
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

func (User *AuthCredentials) VerifyPassword(db *sql.DB) *errors.HttpError {
	hashedPassWord := ""

	err := db.QueryRow("SELECT password FROM users WHERE username = ? OR email = ?", User.Username, User.Username).Scan(&hashedPassWord)
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

func (AC *AuthCredentials) CreateUser(db *sql.DB, resp http.ResponseWriter) (*User, error) {
	hashedPassWord, err := bcrypt.GenerateFromPassword([]byte(AC.Password), 12)
	if err != nil {
		return nil, err
	}
	token, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	pfp := generateAvatarURL(AC.Username)
	user := &User{
		Id:             snowflake.Generate(),
		Username:       strings.ToLower(AC.Username),
		ProfilePicture: &pfp,
	}

	query := "INSERT INTO users (id,username,age,gender,firstname,lastname,password,email,profile_picture) VALUES (? ,? ,? ,? ,? ,? ,? ,? ,?)"
	_, err = db.Exec(query, user.Id, user.Username, AC.Age, strings.ToLower(AC.Gender), strings.ToLower(AC.Firstname), strings.ToLower(AC.Lastname), hashedPassWord, strings.ToLower(AC.Email), user.ProfilePicture)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	_, err = db.Exec("INSERT INTO sessions (user_id,token,expires_at) VALUES (?, ? ,datetime('now', '+1 hour'))", user.Id, token.String())
	if err != nil {
		return nil, err
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
	return user, nil
}

func (User *AuthCredentials) ValidInfo(db *sql.DB) (httpErr *errors.HttpError) {
	httpErr = &errors.HttpError{
		Status:  http.StatusBadRequest,
		Code:    errors.CodeInvalidUsername,
		Message: "invalid username",
	}
	// check username
	if valid, msg := ValidUsername(User.Username); !valid {
		httpErr.Details = msg
		return
	}
	if valid, msg := ValidFirstname(User.Firstname); !valid {
		httpErr.Details = msg
		return
	}
	if valid, msg := ValidLastname(User.Lastname); !valid {
		httpErr.Details = msg
		return
	}
	if valid, msg := ValidAge(User.Age); !valid {
		httpErr.Details = msg
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
		httpErr.Details = "invalid email"
		return
	}

	err = db.QueryRow("SELECT 1 FROM users WHERE email = ?", strings.ToLower(User.Email)).Scan(&exists)
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
		httpErr.Details = "password must be between 12 and 256 characters"
		return
	}
	return nil
}

func generateAvatarURL(username string) string {
	escaped := url.PathEscape(username)
	return fmt.Sprintf("https://api.dicebear.com/7.x/bottts/svg?seed=%s", escaped)
}
