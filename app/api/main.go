package api

import (
	"database/sql"
	"net/http"
	"strings"

	"forum/app/api/auth"
	"forum/app/api/comments"
	"forum/app/api/posts"
	"forum/app/api/reactions"
	useractivities "forum/app/api/userActivities"
	"forum/app/modules"
	"forum/app/modules/errors"
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
		switch req.Method {
		case http.MethodGet:
			posts.GetPosts(conn, forumDB)
		case http.MethodPost:
			posts.AddPost(conn, forumDB)
		default:
			conn.Error(errors.HttpInternalServerError)
		}
	case "comments":
		switch req.Method {
		case http.MethodGet:
			comments.GetComments(conn, forumDB)
		case http.MethodPost:
			comments.AddComment(conn, forumDB)
		case http.MethodPatch:
			comments.UpdateComment(conn, forumDB)
		default:
			conn.Error(errors.HttpMethodNotAllowed)
		}
	case "reactions":
		reactions.HandleReactions(conn, forumDB)
	case "activities":
		useractivities.GetUSer(conn, forumDB)
	default:
		conn.Error(errors.HttpNotFound)
	}
}
