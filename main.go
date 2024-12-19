package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	db "forum/database"

	"golang.org/x/crypto/bcrypt"
)

func hand_register_get(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/register.html")
}

func hand_login_get(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/login.html")
}

func hand_login_post(w http.ResponseWriter, r *http.Request) {
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

func hand_register_post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form data: ", err)
		return
	}
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	i, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println("Error creating user: ", err)
	}
	// Create the user
	err = db.CreateUser(username, email, string(i))
	if err != nil {
		fmt.Println("Error creating user: ", err)
		return
	}
	// Redirect to the login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func home(w http.ResponseWriter, r *http.Request) {
}

func main() {
	db.InitDB()
	http.HandleFunc("GET /login", hand_login_get)
	http.HandleFunc("POST /api/login", hand_login_post)
	http.HandleFunc("GET /register", hand_register_get)
	http.HandleFunc("POST /api/register", hand_register_post)
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
