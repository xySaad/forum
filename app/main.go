package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"forum/app/api"
	"forum/app/config"
	db "forum/app/database"
	"forum/app/handlers"
	logs "forum/app/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	logs.InitLogger()

	forumDB, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}

	db.CreateTables(forumDB)

	config.InitTemplates("templates/*.html")
	config.InitTemplates("templates/components/*.html")

	http.HandleFunc("/static/", handlers.Static)
	http.HandleFunc("/auth/", handlers.AuthPage)
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) { api.Router(w, r, forumDB) })
	http.HandleFunc("/", handlers.Home)

	go func() {
		fmt.Println("server started: http://localhost:8080")
		err = http.ListenAndServe(":8080", nil)
		if err != nil {
			logs.Logger.Println("Error starting server: ", err)
			logs.CloseLogger()
			forumDB.Close()
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	logs.Logger.Println("\nReceived termination signal:", sig)
	logs.Logger.Println("Shutting down gracefully...")
	forumDB.Close()
	logs.Logger.Println("Database connection closed.")
	logs.CloseLogger()
}
