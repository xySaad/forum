package handlers

import (
	"forum/app/config"
	"net/http"
)

func Home(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "405 - method not allowed", 405)
		return
	}
	if req.URL.Path != "/" {
		http.Error(resp, "404 - page not found", 404)
		return
	}
	config.Templates.Exec(resp, "home.html", nil)
}
