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
	"forum/app/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	err := config.InitLogger()
	if err != nil {
		log.Fatal(err)
	}

	forumDB, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}

	config.CreateTables(forumDB)

	http.HandleFunc("/static/", handlers.Static)
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) { api.Router(w, r, forumDB) })
	http.HandleFunc("/", handlers.Home)

	go func() {
		fmt.Println("server started: http://localhost:8080")
		err = http.ListenAndServe(":8080", nil)
		if err != nil {
			config.Logger.Println("Error starting server: ", err)
			config.CloseLogger()
			forumDB.Close()
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	config.Logger.Println("\nReceived termination signal:", sig)
	config.Logger.Println("Shutting down gracefully...")
	forumDB.Close()
	config.Logger.Println("Database connection closed.")
	config.CloseLogger()
}
