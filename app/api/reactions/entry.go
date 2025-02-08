package reactions

import (
	"database/sql"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func Entry(conn *modules.Connection, forumDB *sql.DB) {
	switch conn.Req.Method {
	case http.MethodPost:
		AddReaction(conn, forumDB)
	case http.MethodDelete:
		RemoveReaction(conn, forumDB)
	default:
		conn.NewError(http.StatusMethodNotAllowed, errors.CodeMethodNotAllowed, "Method Not Allowed", "Only Post/Delete are Allowed")
	}
}
