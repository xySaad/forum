package main

import (
	"fmt"
	db "forum/database"
	"net/http"
)

func main() {
	db.InitDB()
	http.HandleFunc("GET /register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/register.html")
	})
	http.HandleFunc("POST /api/register", func(w http.ResponseWriter, r *http.Request) {	
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing form data: ", err)
			return
		}
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		// Create the user
		err = db.CreateUser(username, email, password)
		if err != nil {
			fmt.Println("Error creating user: ", err)
			return
		}
		// Redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/roadmap.html")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
