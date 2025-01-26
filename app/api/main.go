package api

import (
	"database/sql"
	"forum/app/api/auth"
	"forum/app/api/comments"
	"forum/app/api/posts"
	"forum/app/api/reactions"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
	"strings"
)

func Router(resp http.ResponseWriter, req *http.Request, forumDB *sql.DB) {
	conn := &modules.Connection{
		Resp: resp,
		Req:  req,
		Path: strings.Split(req.URL.Path, "/")[1:],
	}

	switch conn.Path[1] {
	case "auth":
		auth.Entry(conn, forumDB)
	case "posts":
		if req.Method == http.MethodGet {
			posts.GetPosts(conn, forumDB)
		} else if req.Method == http.MethodPost {
			posts.AddPost(conn, forumDB)
		} else {
			conn.Error(errors.HttpNotFound)
		}
	case "coments":
		if req.Method == http.MethodPost {
			comments.AddComment(conn, forumDB)
		} else if req.Method == http.MethodGet {
			comments.GetComents(conn, forumDB)
		} else {
			conn.Error(errors.HttpMethodNotAllowed)
			return
		}
	case "reactions":
		reactions.HandleReactions(conn, forumDB)
	default:
		conn.Error(errors.HttpNotFound)
	}
}
