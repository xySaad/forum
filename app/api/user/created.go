package user

import (
	"database/sql"

	"forum/app/api/posts"
	"forum/app/modules"
	"forum/app/modules/errors"
)

func GetUserCreatedPosts(conn *modules.Connection, db *sql.DB) {
	if !conn.IsAuthenticated(db) {
		return
	}
	lastId := conn.Req.URL.Query().Get("lastId")
	query := "SELECT p.id,user_id,title,content,created_at FROM posts p WHERE user_id=? "
	params := []any{conn.User.Id}

	if lastId != "" {
		params = append(params, any(lastId))
		query += "AND p.id < ? "
	}
	query += "ORDER BY p.id DESC LIMIT 10;"

	posts, err := posts.FetchPosts(query, params, conn.User.Id, db)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		return
	}
	conn.Respond(posts)
}
