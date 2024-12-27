package forum

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode"

	database "forum/database"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

func Hand_register_get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "templates/register.html")
}

func VerifyPassword(password string) error {
	if len(password) < 8 {
		return errors.New("password should be atleast 8 chars long")
	}
	return nil
}

func VerifyName(name string) error {
	if len(name) < 4 || len(name) > 18 {
		return errors.New("Name-lenght should be between 4 and 18")
	}
	for _, char := range name {
		if !unicode.IsGraphic(char) {
			return errors.New("Name should only contain printable characters")
		}
	}
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return err
	}
	query := "SELECT  (username) FROM users WHERE username= ( ?)"
	err = db.QueryRow(query, name).Scan()
	if err != sql.ErrNoRows {
		return errors.New("Name already exists")
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
			return errors.New("Email should only contain printable characters")
		}
	}
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return err
	}

	query := "SELECT  (email) FROM users WHERE email= ( ?)"
	err = db.QueryRow(query, Email).Scan()
	if err != sql.ErrNoRows {
		return errors.New("Email already exists")
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

func Hand_register_post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid form", http.StatusNotAcceptable)
		return
	}
	if err := Verify(user); err != nil {
		fmt.Println(err)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	uuid, err := uuid.NewV7()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = database.CreateUser(user.Username, user.Email, string(hashedPassword), uuid.String())
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("set-cookie", "token="+uuid.String())
}

func Hand_login_get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "templates/login.html")
}

func Hand_login_post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid form", http.StatusNotAcceptable)
		return
	}
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	query := "SELECT password FROM users WHERE username=? AND email=?"
	var password string
	err = db.QueryRow(query, user.Username, user.Email).Scan(&password)
	if err != nil {

		if err == sql.ErrNoRows {
			fmt.Println("No user found")
		}
	}
	err = bcrypt.CompareHashAndPassword( []byte(user.Password),[]byte(password))
	if err != nil {
		w.Write([]byte("oh snap"))
	} else {
		w.Write([]byte("bim"))
	}
}
