package api

import (
	"forum/app/api/auth"
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
	default:
		http.Error(conn.Resp, "404 - page not found", 404)
		return
	}
}
