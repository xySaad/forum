package posts

import (
	"database/sql"
	"encoding/json"

	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
	"forum/app/modules/snowflake"
)

func AddPost(conn *modules.Connection, forumDB *sql.DB) {
	var postContent modules.PostContent
	err := json.NewDecoder(conn.Req.Body).Decode(&postContent)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		return
	}

	httpErr := postContent.ValidatePostContent()
	if httpErr != nil {
		conn.Error(httpErr)
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

	postID, err := CreatePost(&postContent, userID, forumDB)
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

func CreatePost(content *modules.PostContent, userID int, forumDB *sql.DB) (int64, error) {
	postID := snowflake.Default.Generate()

	sqlQuery := "INSERT INTO posts (id, title, content, user_internal_id) VALUES (?, ?, ?, ?)"
	result, err := forumDB.Exec(sqlQuery, postID, content.Title, content.Text, userID)
	if err != nil {
		return 0, err
	}
	internalPostID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}
	for _, category := range content.Categories {
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
