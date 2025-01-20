package posts

import (
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func HandlePosts(conn *modules.Connection, path []string, method string) {
	// Ensure there's a valid path segment
	if len(path) < 1 {
		http.Error(conn.Resp, "400 - bad request", 400)
		return
	}

	switch method {
	case http.MethodPost:
		// Add a new post
		AddPost(conn)
	default:
		conn.NewError(http.StatusMethodNotAllowed, errors.CodeMethodNotAllowed, "Method Not Allowed", "Only Post/Get/Delete/Put are Allowed")
	}
}
