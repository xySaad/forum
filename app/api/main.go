package api

import (
	"database/sql"
	"encoding/json"
	"forum/app/api/auth"
	"forum/app/api/comments"
	"forum/app/api/posts"
	"forum/app/api/reactions"
	"forum/app/config"
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
			err := posts.GetPosts(conn, forumDB)
			if err != nil {
				config.Logger.Println(err)
			}
			return
		}
		if req.Method == http.MethodPost {
			posts.AddPost(conn, forumDB)
			return
		}
		conn.Error(errors.HttpNotFound)
	case "coments":
		if req.Method == http.MethodPost {
			err := comments.AddComment(conn, forumDB)
			if err != nil {
				http.Error(resp, err.Error()+"500 - internal server error", 500)
				return
			}
		} else if req.Method == http.MethodGet {
			coments, err := comments.GetComents(req.URL, forumDB)
			if err != nil {
				http.Error(resp, err.Error()+"500 - internal server error", 500)
				return
			}
			err = json.NewEncoder(resp).Encode(coments)
			if err != nil {
				http.Error(resp, err.Error()+"500 - internal server error", 500)
				return
			}
		} else {
			http.Error(conn.Resp, "405 - method not allowed ", 405)
			return
		}
	case "reactions":
		reactions.HandleReactions(conn, forumDB)
	default:
		conn.Error(errors.HttpNotFound)
	}
}
