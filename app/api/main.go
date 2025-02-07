package api

import (
	"database/sql"
	"net/http"
	"strings"

	"forum/app/api/auth"
	"forum/app/api/comments"
	"forum/app/api/posts"
	reactions "forum/app/api/reaction"
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
		if req.Method == http.MethodGet {
			posts.GetPosts(conn, forumDB)
			return
		}
		if req.Method == http.MethodPost {
			posts.AddPost(conn, forumDB)
		} else {
			conn.Error(errors.HttpMethodNotAllowed)
		}
	case "coments":
		if req.Method == http.MethodPost {
			comments.AddComment(conn, forumDB)
		} else if req.Method == http.MethodGet {
			comments.GetComents(conn, forumDB)
		} else{
			conn.Error(errors.HttpMethodNotAllowed)
			return
		}
	case "reactions":
		if req.Method != http.MethodPost {
			conn.Error(errors.HttpMethodNotAllowed)
			return
		}
		reactions.HandleReactions(conn, forumDB)
	case "like":
		if req.Method != http.MethodGet {
			conn.Error(errors.HttpMethodNotAllowed)
			return
		}
		useractivities.GetUSerLiked(conn, forumDB)
	case "post":
		if req.Method != http.MethodGet {
			conn.Error(errors.HttpMethodNotAllowed)
			return
		}
		useractivities.GetUSerPosts(conn, forumDB)
	case "profile" : 
		auth.GetUserData(conn, forumDB)
	default:
		conn.Error(errors.HttpNotFound)
	}
}
