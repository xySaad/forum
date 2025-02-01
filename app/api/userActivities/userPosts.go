package useractivities

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
)

func GetUSer(conn *modules.Connection, forumDB *sql.DB) {
	SpUrl := strings.Split(conn.Req.URL.String(), "/")
	if len(SpUrl) != 3 {
		conn.NewError(http.StatusNotFound, 404, "not found", "")
		return
	}
	token, err := conn.Req.Cookie("token")
	if err != nil || token.Value == "" {
		conn.NewError(http.StatusUnauthorized, http.StatusUnauthorized, "unauthorized", "")
		return
	}
	userId, httpErr := handlers.GetUserIDByToken(token.Value, forumDB)
	if err != nil {
		conn.Error(httpErr)
		return
	}
	switch SpUrl[2] {
	case "posts":
		GetUSerPosts(conn, userId, forumDB)
	case "like":
		GetUserReactions(conn, userId, "like", forumDB)
	case "dislike":
		GetUserReactions(conn, userId, "dislike", forumDB)
	default:
		conn.NewError(http.StatusNotFound, 404, "not found", "")
	}
}

func GetUserReactions(conn *modules.Connection, uId, reaction string, db *sql.DB) {
	user, err := GetUser(uId, db)
	if err != nil {
		conn.NewError(http.StatusNotFound, errors.CodeUserNotFound, "user not found", "")
	}
	posts := []modules.Post{}
	query := `SELECT SUBSTRING_INDEX(item_id,'_',-1) FROM reactions WHERE user_id=? AND SUBSTRING_INDEX(item_id,'_',1='posts') AND reaction_type=?`
	rows, err := db.Query(query, uId, reaction)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, 500, "internal pointer variable", "")
		return
	}
	for rows.Next() {
		postID := ""
		if err := rows.Scan(&postID); err != nil {

			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
			return
		}
		var post modules.Post
		query = `SELECT content, title,categories, created_at FROM posts WHERE id= ? ORDER BY updated_at DESC`
		if err = db.QueryRow(query, postID).Scan(&post.Text, &post.Title, &post.Categories, &post.CreationTime); err != nil {
			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
			return
		}
		post.Publisher = user
		posts = append(posts, post)
	}
	err = json.NewEncoder(conn.Resp).Encode(posts)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
	}
}

func GetUser(uID string, db *sql.DB) (user modules.User, err error) {
	query := `SELECT username FROM users WHERE id=?`
	err = db.QueryRow(query, uID).Scan(&user.Username)
	return
}

func GetUSerPosts(conn *modules.Connection, userId string, db *sql.DB) {
	userPosts := []modules.Post{}
	user, err := GetUser(userId, db)
	if err != nil {
		conn.NewError(http.StatusNotFound, errors.CodeUserNotFound, "user not found", "")
	}
	query := `SELECT content, title,categories, created_at FROM posts WHERE user_id = ? ORDER BY updated_at DESC LIMIT 10 OFFSET ?`
	rows, err := db.Query(query, userId)
	if err != nil {
	}
	defer rows.Close()

	for rows.Next() {
		var post modules.Post
		if err := rows.Scan(&post.Text, &post.Title, &post.Categories, &post.CreationTime); err != nil {
			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
		}
		post.Publisher = user
		userPosts = append(userPosts, post)
	}
	err = json.NewEncoder(conn.Resp).Encode(userPosts)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
	}
}
