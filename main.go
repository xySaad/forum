package main

import (
	"fmt"
	"net/http"

	authUtils "forum/authUtils"
	db "forum/database"
)

func home(w http.ResponseWriter, r *http.Request) {
}

func main() {
	db.InitDB()
	http.HandleFunc("GET /login", authUtils.Hand_login_get)
	http.HandleFunc("POST /login", authUtils.Hand_login_post)
	http.HandleFunc("GET /register", authUtils.Hand_register_get)
	http.HandleFunc("POST /register", authUtils.Hand_register_post)
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
