package forum

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	database "forum/database"

	"golang.org/x/crypto/bcrypt"
)

func Hand_register_get(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/register.html")
}

func Hand_register_post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form data: ", err)
		return
	}
	username := r.Form.Get("username")
	if username=="" {
		w.Write([]byte("username cant be empty"))
		return
	}
	email := r.Form.Get("email")
	if email=="" {
		w.Write([]byte("email cant be empty"))
		return
	}
	password := r.Form.Get("password")
	if password=="" {
		w.Write([]byte("password cant be empty"))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println("Error creating user: ", err)
	}
	// Create the user
	err = database.CreateUser(username, email, string(hashedPassword))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte("FFFF"))
	http.Redirect(w, r, "/", 200)
}

func Hand_login_get(w http.ResponseWriter, r *http.Request) {
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
