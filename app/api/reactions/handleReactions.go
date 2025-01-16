package reactions

import (
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func HandleReactions(conn *modules.Connection, path []string, method string) {
	// Ensure there's a valid path segment
	if len(path) < 1 {
		http.Error(conn.Resp, "400 - bad request", 400)
		return
	}

	switch method {
	case http.MethodPost:
		AddReaction(conn)
	case http.MethodDelete:
		RemoveReaction(conn)
	case http.MethodGet:
		if len(path) < 2 {
			conn.NewError(http.StatusBadRequest, errors.CodeInvalidOrMissingData, "Post ID required", "No Post ID provided")
			return
		}
		GetReaction(conn, path[1])
	default:
		conn.NewError(http.StatusMethodNotAllowed, errors.CodeMethodNotAllowed, "Method Not Allowed", "Only Post/Get/Delete are Allowed")
	}
}
