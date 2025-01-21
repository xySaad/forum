package posts

import (
	"database/sql"
	"encoding/json"
	"fmt"
	db "forum/app/database"
	"forum/app/modules"
	"forum/app/modules/errors"
	"log"
	"net/http"
)

func AddPost(conn *modules.Connection) {
	var request struct {
		Title      string   `json:"title"`
		Content    string   `json:"content"`
		Categories []string `json:"categories"`
	}

	err := json.NewDecoder(conn.Req.Body).Decode(&request)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(conn.Resp, "400 - invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Validating title: %s, content: %s", request.Title, request.Content)
	if err := validateTitleContent(request.Title, request.Content); err != nil {
		log.Printf("Validation failed: %v", err)
		conn.NewError(http.StatusBadRequest, errors.CodeInvalidOrMissingData, err.Error(), "")
		return
	}

	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		log.Printf("Authentication token missing or invalid")
		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Missing or invalid authentication token", "")
		return
	}

	userID, err := db.GetUserIDByToken(cookie.Value)
	if err != nil {
		log.Printf("Error getting user ID from token: %v", err)
		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Invalid or expired authentication token", "")
		return
	}

	log.Printf("Creating post for user ID: %s", userID)
	postID, err := CreatePost(request.Title, request.Content, userID, request.Categories)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		conn.NewError(http.StatusInternalServerError, errors.CodeInternalServerError, "Internal Server Error", "The server encountered an error, please try again later.")
		return
	}

	log.Printf("Post created successfully, postID: %d", postID)
	conn.Resp.WriteHeader(http.StatusOK)
	json.NewEncoder(conn.Resp).Encode(map[string]interface{}{
		"message": "Post created successfully",
		"postID":  postID,
	})
}

func validateTitleContent(title, content string) error {
	if title == "" {
		log.Printf("Validation failed: Title is empty")
		return fmt.Errorf("Title cannot be empty")
	}
	if content == "" {
		log.Printf("Validation failed: Content is empty")
		return fmt.Errorf("Content cannot be empty")
	}

	if len(title) < 5 || len(title) > 100 {
		log.Printf("Validation failed: Title length is invalid (title length: %d)", len(title))
		return fmt.Errorf("Title must be between 5 and 100 characters")
	}

	if len(content) < 10 || len(content) > 5000 {
		log.Printf("Validation failed: Content length is invalid (content length: %d)", len(content))
		return fmt.Errorf("Content must be between 10 and 5000 characters")
	}

	log.Printf("Validation successful for title and content")
	return nil
}

func GetCategoryMask(categories []string) string {
	categoryMap := map[string]int{
		"Sport":      0,
		"Technology": 1,
		"Finance":    2,
		"Science":    3,
	}

	mask := [4]rune{'0', '0', '0', '0'}

	log.Printf("Input categories: %v", categories)

	for _, category := range categories {
		if index, exists := categoryMap[category]; exists {
			mask[index] = '1'
		} else {
			log.Printf("Unknown category: %s", category)
		}
	}

	log.Printf("Generated category mask: %s", string(mask[:]))
	return string(mask[:])
}

func CreatePost(title, content, userID string, categories []string) (int64, error) {
	categoryMask := GetCategoryMask(categories)

	log.Printf("Inserting post into database with title: %s, categories: %s", title, categoryMask)

	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return 0, err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO posts (title, content, user_id, categories) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("Error preparing SQL statement: %v", err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(title, content, userID, categoryMask)
	if err != nil {
		log.Printf("Error executing SQL statement: %v", err)
		return 0, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error fetching last insert ID: %v", err)
		return 0, err
	}

	log.Printf("Post created with ID: %d", postID)
	return postID, nil
}
