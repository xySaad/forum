package main

import (
	"fmt"
	"net/http"

	"forum/app/config"
	db "forum/app/database"
	"forum/app/handlers"
	"forum/app/handlers/api"
)

func main() {
	db.InitDB()
	config.InitTemplates("templates/*.html")
	config.InitTemplates("templates/components/*.html")

	http.HandleFunc("/static/", handlers.Static)
	http.HandleFunc("/auth/", handlers.Auth)
	http.HandleFunc("/api/", api.Entry)
	http.HandleFunc("/", handlers.Home)

	fmt.Println("server started: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("Error starting server: ", err)
}
