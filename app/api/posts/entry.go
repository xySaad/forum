package posts

import (
	"database/sql"
	"forum/app/api/posts/comments"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func Entry(conn *modules.Connection, forumDB *sql.DB) {
	switch len(conn.Path) {
	case 2:
		switch conn.Req.Method {
		case http.MethodGet:
			GetBulkPosts(conn, forumDB)
		case http.MethodPost:
			AddPost(conn, forumDB)
		default:
			conn.Error(errors.HttpMethodNotAllowed)
		}
	case 3:
		if conn.Req.Method != http.MethodGet {
			conn.Error(errors.HttpMethodNotAllowed)
			return
		}
		GetSinglePost(conn, forumDB)

	case 4:
		nestedRoutes(conn, forumDB)
	default:
		conn.Error(errors.HttpNotFound)
	}
}

func nestedRoutes(conn *modules.Connection, forumDB *sql.DB) {
	if conn.Path[3] != "comments" {
		conn.Error(errors.HttpNotFound)
		return
	}

	switch conn.Req.Method {
	case http.MethodGet:
		comments.GetPostComments(conn, forumDB)
	case http.MethodPost:
		comments.AddComment(conn, forumDB)
	case http.MethodPatch:
		comments.UpdateComment(conn, forumDB)
	default:
		conn.Error(errors.HttpMethodNotAllowed)
	}
}
