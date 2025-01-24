package main

import (
	"database/sql"
	"fmt"
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
		config.Logger.Println(err)
		return
	}
	defer config.CloseLogger()

	forumDB, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		config.Logger.Println(err)
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigChan)

	defer func() {
		err = forumDB.Close()
		if err != nil {
			config.Logger.Println("Error closing database connection:", err)
		} else {
			config.Logger.Println("Database connection closed.")
		}
	}()

	err = config.CreateTables(forumDB)
	if err != nil {
		config.Logger.Println("Error creating tables:", err)
		return
	}

	http.HandleFunc("/static/", handlers.Static)
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) { api.Router(w, r, forumDB) })
	http.HandleFunc("/", handlers.Home)

	server := &http.Server{Addr: ":8080"}

	go func() {
		fmt.Println("server started: http://localhost:8080")
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			config.Logger.Println("Error starting server: ", err)
			sigChan <- syscall.SIGTERM
		}
	}()

	config.Logger.Print("Shutting down... signal: ", <-sigChan)

	err = server.Close()
	if err != nil {
		config.Logger.Println("Error during graceful shutdown: ", err)
	} else {
		config.Logger.Println("Server gracefully stopped.")
	}
}
