package main

import (
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"forum/app/api"
	"forum/app/config"
	"forum/app/handlers"
	"forum/app/modules/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	forumDB, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("error opening database:", err)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigChan)

	defer func() {
		err = forumDB.Close()
		if err != nil {
			log.Error("error closinging database:", err)
		} else {
			log.Info("database closed successfully")
		}
	}()

	err = config.CreateTables(forumDB)
	if err != nil {
		log.Fatal("error creating tables:", err)
	}

	http.HandleFunc("/static/", handlers.Static)
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) { api.Router(w, r, forumDB) })
	http.HandleFunc("/", handlers.Home)

	server := &http.Server{Addr: ":8080"}

	go func() {
		log.Info("server started: http://localhost:8080")
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error("error starting server:", err)
			sigChan <- syscall.SIGTERM
		}
	}()

	log.Info("shuting down the server", <-sigChan)
	err = server.Close()
	if err != nil {
		
		log.Info("server shutdown successfully")
	}
}
