package comments

import (
	"database/sql"
	"net/http"
	"strconv"

	"forum/app/api/posts"
	reactions "forum/app/api/reaction"
	"forum/app/handlers"
	"forum/app/modules"
)

func GetComents(conn *modules.Connection, forumDB *sql.DB) {
	URL := conn.Req.URL
	post_id := URL.Query().Get("p_id")
	Offset := URL.Query().Get("offset")
	from := URL.Query().Get("from")
	uid := ""
	cookie, err := conn.Req.Cookie("token")
	if err == nil && cookie.Value != "" {
		uid, _ = handlers.GetUserIDByToken(cookie.Value, forumDB)
	}
	if from == "" {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}
	fromN, err := strconv.Atoi(from)
	if err != nil {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
	}
	if Offset == "" {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}
	offset, err := strconv.Atoi(Offset)
	if err != nil {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}
	if post_id == "" {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}
	p_id, err := strconv.Atoi(post_id)
	if err != nil {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}
	query := `SELECT id, post_id, user_id, content, likes, dislikes, created_at FROM comments WHERE `
	if fromN > 0 {
		query += `id<=` + from
	}
	query += `post_id = ? ORDER BY created_at DESC LIMIT 10 OFFSET ?`
	rows, err := forumDB.Query(query, p_id, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			conn.NewError(http.StatusNotFound, 404, "no post with such id", "")
			return
		}

		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
		return
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ItemID, &comment.PostID, &comment.Publisher.Id, &comment.Content, &comment.Likes, &comment.Dislikes, &comment.CreationTime); err != nil {
			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
			return
		}
		comment.Dislikes, comment.Likes, comment.Reaction = reactions.GetReactions("comment-"+comment.ItemID, uid, forumDB)
		err = posts.GetPublicUser(&comment.Publisher, forumDB)
		if err != nil {
			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
			return
		}
		comments = append(comments, comment)
	}
	conn.Respond(comments)
}
