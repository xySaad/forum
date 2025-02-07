package comments

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/app/modules"
)

func AddComment(conn *modules.Connection, forumDB *sql.DB) {
	if !conn.IsAuthenticated(forumDB) {
		return
	}

	var comment Comment
	err := json.NewDecoder(conn.Req.Body).Decode(&comment)
	if err != nil {
		conn.NewError(http.StatusBadRequest, 400, "ivalid format", "")
		return
	}
	if comment.PostID == "" || comment.Content == "" {
		conn.NewError(http.StatusBadRequest, 400, ",issing data", "")
		return
	}

	query := `INSERT INTO comments (post_id, user_id, content, likes, dislikes) VALUES (?, ?, ?, ?, ?)`
	_, err = forumDB.Exec(query, comment.PostID, conn.InternalUserId, comment.Content, 0, 0)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
		return
	}
}
