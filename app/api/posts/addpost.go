package posts

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func AddPost(conn *modules.Connection, forumDB *sql.DB) {
	var request struct {
		Title      string   `json:"title"`
		Content    string   `json:"content"`
		Categories []string `json:"categories"`
	}

	err := json.NewDecoder(conn.Req.Body).Decode(&request)
	if err != nil {

		http.Error(conn.Resp, "400 - invalid request body", http.StatusBadRequest)
		return
	}

	if err := validateTitleContent(request.Title, request.Content); err != nil {

		conn.NewError(http.StatusBadRequest, errors.CodeInvalidOrMissingData, err.Error(), "")
		return
	}

	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {

		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Missing or invalid authentication token", "")
		return
	}

	userID, err := handlers.GetUserIDByToken(cookie.Value, forumDB)
	if err != nil {

		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Invalid or expired authentication token", "")
		return
	}

	postID, err := CreatePost(request.Title, request.Content, userID, request.Categories, forumDB)
	if err != nil {

		conn.NewError(http.StatusInternalServerError, errors.CodeInternalServerError, "Internal Server Error", "The server encountered an error, please try again later.")
		return
	}

	conn.Resp.WriteHeader(http.StatusOK)
	json.NewEncoder(conn.Resp).Encode(map[string]interface{}{
		"message": "Post created successfully",
		"postID":  postID,
	})
}

func validateTitleContent(title, content string) error {
	if title == "" {

		return fmt.Errorf("Title cannot be empty")
	}
	if content == "" {

		return fmt.Errorf("Content cannot be empty")
	}

	if len(title) < 5 || len(title) > 100 {

		return fmt.Errorf("Title must be between 5 and 100 characters")
	}

	if len(content) < 10 || len(content) > 5000 {

		return fmt.Errorf("Content must be between 10 and 5000 characters")
	}

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

	for _, category := range categories {
		if index, exists := categoryMap[category]; exists {
			mask[index] = '1'
		} else {

		}
	}

	return string(mask[:])
}

func CreatePost(title, content, userID string, categories []string, forumDB *sql.DB) (int64, error) {
	categoryMask := GetCategoryMask(categories)

	stmt, err := forumDB.Prepare("INSERT INTO posts (title, content, user_id, categories) VALUES (?, ?, ?, ?)")
	if err != nil {

		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(title, content, userID, categoryMask)
	if err != nil {

		return 0, err
	}

	postID, err := result.LastInsertId()
	if err != nil {

		return 0, err
	}

	return postID, nil
}
