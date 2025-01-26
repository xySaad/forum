package comments

import (
	"database/sql"
	"net/http"
	"strconv"

	"forum/app/modules"
)

func GetComents(conn *modules.Connection, forumDB *sql.DB) {
	URL := conn.Req.URL
	post_id := URL.Query().Get("p_id")
	Offset := URL.Query().Get("offset")
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
	query := `SELECT id, post_id, user_id, content, created_at FROM comments WHERE post_id = ? ORDER BY updated_at DESC LIMIT 10 OFFSET ?`
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
		if err := rows.Scan(&comment.Item_id, &comment.Post_id, &comment.User_id, &comment.Content, &comment.Created_at); err != nil {
			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
			return
		}
		comments = append(comments, comment)
	}
	conn.Respond(comments)
}
