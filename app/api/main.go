package api

import (
	"encoding/json"
	"forum/app/api/auth"
	"forum/app/api/comments"
	"forum/app/api/posts"
	"forum/app/api/reactions"
	"forum/app/modules"
	"net/http"
	"strings"
)

func Router(resp http.ResponseWriter, req *http.Request) {
	conn := &modules.Connection{
		Resp: resp,
		Req:  req,
	}

	path := strings.Split(req.URL.Path[5:], "/")
	switch path[0] {
	case "auth":
		auth.Entry(conn)
	case "posts":
		if req.Method == http.MethodGet {
			data, err := posts.GetPost(req.URL.Path)
			if err != nil {
				http.Error(resp, "500 - internal server error", 500)
				return
			}
			resp.Header().Set("Content-Type", "application/json")
			resp.Write(data)
		}
	case "coments":
		if req.Method == http.MethodPost {
			err := comments.AddComment(req.Body)
			if err != nil {
				http.Error(resp, err.Error()+"500 - internal server error", 500)
				return
			}
		} else if req.Method == http.MethodGet {
			coments, err := comments.GetComents(req.URL)
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
		reactions.HandleReactions(conn, path, conn.Req.Method)
	default:
		http.Error(conn.Resp, "404 - page not found", 404)
		return
	}
}
