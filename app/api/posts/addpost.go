package posts

import (
	"database/sql"
	"encoding/json"
	"fmt"
	db "forum/app/database"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func AddPost(conn *modules.Connection) {
	var request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	// Decode JSON request body
	err := json.NewDecoder(conn.Req.Body).Decode(&request)
	if err != nil {
		http.Error(conn.Resp, "400 - invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if request.Title == "" || request.Description == "" {
		conn.NewError(http.StatusBadRequest, errors.CodeInvalidOrMissingData, "Title and Description cannot be empty", "")
		return
	}

	// Retrieve the user token from cookies
	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Missing or invalid authentication token", "")
		return
	}

	// Get the user ID using the token
	userID, err := db.GetUserIDByToken(cookie.Value)
	if err != nil {
		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Invalid or expired authentication token", "")
		return
	}

	// Save the post to the database
	err = CreatePost(request.Title, request.Description, userID)
	if err != nil {
		fmt.Println(err)
		conn.NewError(http.StatusInternalServerError, errors.CodeInternalServerError, "Internal Server Error", "The server encountered an error, please try again at a later time.")
		return
	}

	// Respond with success
	conn.Resp.WriteHeader(http.StatusOK)
	json.NewEncoder(conn.Resp).Encode(map[string]string{"message": "Post created successfully"})
}
func CreatePost(title, description, userID string) error {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO posts (title, description, user_id, timestamp) VALUES (?, ?, ?, datetime('now'))")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, description, userID)
	return err
}
