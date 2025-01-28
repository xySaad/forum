package comments

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/app/handlers"
	"forum/app/modules"
)

func AddComment(conn *modules.Connection, forumDB *sql.DB) {
	var comment Comment
	err := json.NewDecoder(conn.Req.Body).Decode(&comment)
	if err != nil {
		conn.NewError(http.StatusBadRequest, 400, "ivalid format", "")
		return
	}
	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		conn.NewError(http.StatusForbidden, 403, "unothorized", "")
		return
	}
	uId, httpErr := handlers.GetUserIDByToken(cookie.Value, forumDB)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}

	query := `INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)`
	_, err = forumDB.Exec(query, comment.PostID, uId, comment.Content)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
		return
	}
}
