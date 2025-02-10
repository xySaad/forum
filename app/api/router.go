package api

import (
	"database/sql"
	"forum/app/api/auth"
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
	if conn.Path[len(conn.Path)-1] == "" {
		conn.Path = conn.Path[:len(conn.Path)-1]
	}
	switch conn.Path[1] {
	case "auth":
		auth.Entry(conn, forumDB)
	case "posts":
		posts.Entry(conn, forumDB)
	case "reactions":
		reactions.Entry(conn, forumDB)
	case "profile" :
		auth.GetUserData(conn, forumDB)
	default:
		conn.Error(errors.HttpNotFound)
	}
}
