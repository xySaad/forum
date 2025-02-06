package posts

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/app/config"
	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
	"forum/app/modules/snowflake"
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
		conn.Error(errors.HttpUnauthorized)
		return
	}

	userID, httpErr := handlers.GetUserIDByToken(cookie.Value, forumDB)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}

	postID, err := CreatePost(&request.body, userID, forumDB)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			conn.Error(errors.HttpInternalServerError)
		} else {
			conn.Error(errors.BadRequestError("Inavalid categories"))
		}
		return
	}
	conn.Respond(map[string]int64{"postID": postID})
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
	if len(req.body.Categories) > config.MaxCategoriesSize {
		req.Error(errors.BadRequestError("can't select more than 4 categories"))
		return
	}
	return true
}

func CreatePost(body *postRequestBody, userID int, forumDB *sql.DB) (int64, error) {
	postID := snowflake.Default.Generate()

	sqlQuery := "INSERT INTO posts (id, title, content, user_internal_id) VALUES (?, ?, ?, ?)"
	result, err := forumDB.Exec(sqlQuery, postID, body.Title, body.Content, userID)
	if err != nil {
		return 0, err
	}
	internalPostID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}
	for _, category := range body.Categories {
		var categoryID string
		err := forumDB.QueryRow("SELECT id FROM categories WHERE name = ?", category).Scan(&categoryID)
		if err != nil {
			return 0, err
		}
		_, err = forumDB.Exec("INSERT INTO post_categories (post_internal_id, category_id) VALUES (?, ?)", internalPostID, categoryID)
		if err != nil {
			return 0, err
		}
	}
	return postID, nil
}
