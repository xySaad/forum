package posts

import (
	"database/sql"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func HandlePosts(conn *modules.Connection, path []string, method string, forumDB *sql.DB) {
	if len(path) < 1 {
		http.Error(conn.Resp, "400 - bad request", 400)
		return
	}

	switch method {
	case http.MethodPost:
		AddPost(conn, forumDB)
	default:
		conn.NewError(http.StatusMethodNotAllowed, errors.CodeMethodNotAllowed, "Method Not Allowed", "Only Post/Get/Delete/Put are Allowed")
	}
}
