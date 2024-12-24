package forum

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	database "forum/database"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Auth(jwt string) (is_auth bool) {
	if jwt == "" {
		return
	}

	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return
	}
	query := "SELECT  ( password) FROM users WHERE password= ( ?)"
	_, err = db.Exec(query, jwt)
	if err != nil {
		return
	}
	is_auth = true
	return
}

func Hand_register_get(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/register.html")
}

func VerifyPassword(password string) (verified bool) {
	if 8 <= len(password) || len(password) <= 64 {
		verified = true
	}
	return
}
func VerifyName(name string) (verified bool) {
	for _, char := range name {
		if char<32 {
			
		}
	}
	return
}
func Verify(user User) (verified bool) {

	return
}

func Hand_register_post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	jwt, err := r.Cookie("token")
	if err == nil {
		if Auth(string(jwt.Value)) {
			fmt.Println("user")
			return
		}
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid form", http.StatusNotAcceptable)
		return
	}
	if user.Email == "" || user.Password == "" || user.Username == "" {
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Println("Error creating user: ", err)
	}

	err = database.CreateUser(user.Username, user.Email, string(hashedPassword))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	Cookie := http.Cookie{
		Name:  "token",
		Value: string(hashedPassword),
	}
	http.SetCookie(w, &Cookie)
}

func Hand_login_get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "templates/login.html")
}

func Hand_login_post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form data: ", err)
		return
	}
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password1 := r.Form.Get("password")
	query := "SELECT password FROM users WHERE username=? AND email=?"
	row := db.QueryRow(query, username, email)
	var password string
	// Scan the result into a variable
	err = row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found")
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(password1))
	if err != nil {
		w.Write([]byte("oh snap"))
	} else {
		w.Write([]byte("bim"))
	}
}
