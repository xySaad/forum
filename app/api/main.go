package api

import (
	"forum/app/api/auth"
	"forum/app/api/posts"
	"net/http"
	"strings"
)

func Router(resp http.ResponseWriter, req *http.Request) {
	path := strings.Split(req.URL.Path[5:], "/")
	switch path[0] {
	case "auth":
		auth.Entry(resp, req)
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
	}
}
