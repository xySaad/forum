package useractivities

import (
	"database/sql"
	"encoding/json"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
	"strings"
)

func GetUSer(conn modules.Connection, url string) {
	SpUrl := strings.Split(url, "/")
	username := SpUrl[0]
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "error opening db")

	}
	defer db.Close()
	query := `SELECT id FROM users WHERE username=?`
	row, err := db.Query(query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			conn.NewError(http.StatusNotFound, errors.CodeUserNotFound, "user doesn't exists", "")
			return
		}
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
		return
	}
	userId := ""
	err = row.Scan(userId)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")

	}
	switch SpUrl[1] {
	case "posts":
		GetUSerPosts(conn, userId, db)
	case "like":
		GetUserReactions(conn, userId, "like", db)
	case "dislike":
		GetUserReactions(conn, userId, "dislike", db)
	default:
		conn.NewError(http.StatusNotFound, 404, "not found", "")
	}
}
func GetUserReactions(conn modules.Connection, uId, reaction string, db *sql.DB) {
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
func GetUSerPosts(conn modules.Connection, userId string, db *sql.DB) {
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
