package user

import (
	"database/sql"

	"forum/app/api/posts"
	"forum/app/modules"
	"forum/app/modules/errors"
)

func GetLikedPosts(conn *modules.Connection, db *sql.DB) {
	if !conn.IsAuthenticated(db) {
		return
	}

	lastId := conn.Req.URL.Query().Get("lastId")
	query := `SELECT p.id,p.user_id,title,content,p.created_at FROM posts p
	JOIN item_reactions ir ON ir.item_id=p.id WHERE ir.user_id=? AND ir.reaction_id=1 `
	params := []any{conn.User}

	if lastId != "" {
		params = append(params, any(lastId))
		query += "AND p.id > ? "
	}
	query += "ORDER BY p.id DESC LIMIT 10;"

	posts, err := posts.FetchPosts(query, params, conn.User.Id, db)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		return
	}
	conn.Respond(posts)
}
