package main

import (
	"fmt"
	"net/http"

	"forum/app/api"
	"forum/app/config"
	db "forum/app/database"
	"forum/app/handlers"
)

func main() {
	db.InitDB()
	// posts.GetPosts(conn * modules.Connection)
	config.InitTemplates("templates/*.html")
	config.InitTemplates("templates/components/*.html")

	http.HandleFunc("/static/", handlers.Static)
	http.HandleFunc("/auth/", handlers.AuthPage)
	http.HandleFunc("/api/", api.Router)
	http.HandleFunc("/", handlers.Home)

	fmt.Println("server started: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
