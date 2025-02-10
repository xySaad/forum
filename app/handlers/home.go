package handlers

import (
	"net/http"
)

func Home(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "405 - method not allowed", 405)
		return
	}
	http.ServeFile(resp, req, "./static/index.html")
}
