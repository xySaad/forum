package posts

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
)

type postRequestBody struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
}

type postRequest struct {
	*modules.Connection
	body postRequestBody
}

func AddPost(conn *modules.Connection, forumDB *sql.DB) {
	request := postRequest{
		Connection: conn,
	}

	err := json.NewDecoder(conn.Req.Body).Decode(&request.body)
	if err != nil {
		http.Error(conn.Resp, "400 - invalid request body", http.StatusBadRequest)
		return
	}

	if !ValidatePostContent(&request) {
		return
	}

	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Missing or invalid authentication token", "")
		return
	}

	userID, httpErr := handlers.GetUserIDByToken(cookie.Value, forumDB)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}

	postID, herr := CreatePost(&request.body, userID, forumDB)
	if herr != nil {
		fmt.Printf("err: %v\n", err)
		conn.Error(herr)
		return
	}

	conn.Resp.WriteHeader(http.StatusOK)
	json.NewEncoder(conn.Resp).Encode(map[string]int64{
		"postID": postID,
	})
}

func ValidatePostContent(req *postRequest) (isValid bool) {
	if len(req.body.Title) == 0 || len([]rune(req.body.Title)) > 50 {
		req.NewError(http.StatusBadRequest, errors.CodeInvalidRequestFormat, "Title can't be empty or more than 50 character", "Post title too long")
		return
	}
	if len(req.body.Content) == 0 || len([]rune(req.body.Content)) > 5000 {
		req.NewError(http.StatusBadRequest, errors.CodeInvalidRequestFormat, "Content can't be empty or more than 5000 character", "Post content too long")
		return
	}
	return true
}

func CreatePost(body *postRequestBody, userID string, forumDB *sql.DB) (int64, *errors.HttpError) {
	result, err := forumDB.Exec("INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)", body.Title, body.Content, userID)
	if err != nil {
		return 0, errors.HttpInternalServerError
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return 0, errors.HttpInternalServerError
	}

	for _, category := range body.Categories {
		categoryID := ""
		err := forumDB.QueryRow("SELECT id FROM categories WHERE name = ?", category).Scan(&categoryID)
		if err != nil {
			if err == sql.ErrNoRows {
				return 0, &errors.HttpError{
					Code:    404,
					Status:  http.StatusNotFound,
					Message: "not found",
					Details: "no such category as " + category,
				}
			}
			return 0, errors.HttpInternalServerError
		}
		_, err = forumDB.Exec("INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)", postID, categoryID)
		if err != nil {
			return 0, errors.HttpInternalServerError
		}
	}
	return postID, nil
}
