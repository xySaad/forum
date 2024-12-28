package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Hand_login_get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./templates/login.html")
}

func Hand_login_post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		redirect("/login", "method not allowed", 405, w)
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		redirect("/login", "invalid format", 406, w)
		// again putting a real err here is beter
		return
	}
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		redirect("/login", "internal server err", 500, w)
		return
	}
	defer db.Close()
	query := "SELECT password FROM users WHERE username=? AND email=?"
	var password string
	err = db.QueryRow(query, user.Username, user.Email).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			redirect("/login", "no account found", 404, w)
		} else {
			redirect("/login", "internal server err", 500, w)
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		redirect("/login", "password uncorrect", 403, w)
		return
	}
	redirect("/", "loged-in", 200, w)
	// i knw my code aint perfect for now but thats the the best
	// i could do so far.... js is shiiiiiit
}
