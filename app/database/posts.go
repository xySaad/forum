package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if post.Title == "" || post.Description == "" {
		http.Error(w, "Title and Description cannot be empty", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO posts (title, description, timestamp) VALUES (?, ?, datetime('now'))")
	if err != nil {
		http.Error(w, "Failed to prepare database statement", http.StatusInternalServerError)
		log.Println("Statement preparation error:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Description)
	if err != nil {
		http.Error(w, "Failed to save post to database", http.StatusInternalServerError)
		log.Println("Database execution error:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Post created successfully")
}