package reactions

import (
	"database/sql"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func HandleReactions(conn *modules.Connection, forumDB *sql.DB) {
	// Ensure there's a valid path segment
	if len(conn.Path) < 2 {
		http.Error(conn.Resp, "400 - bad request", 400)
		return
	}

	switch conn.Req.Method {
	case http.MethodPost:
		AddReaction(conn, forumDB)
	case http.MethodDelete:
		RemoveReaction(conn, forumDB)
	case http.MethodGet:
		if len(conn.Path) < 3 {
			conn.NewError(http.StatusBadRequest, errors.CodeInvalidOrMissingData, "Post ID required", "No Post ID provided")
			return
		}
		GetReaction(conn, forumDB)
	default:
		conn.NewError(http.StatusMethodNotAllowed, errors.CodeMethodNotAllowed, "Method Not Allowed", "Only Post/Get/Delete are Allowed")
	}
}
