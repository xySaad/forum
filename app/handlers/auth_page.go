package handlers

import (
	"forum/app/config"
	"net/http"
)

func AuthPage(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "405 - method not allowed", 405)
		return
	}
	if req.URL.Path == "/auth/" || req.URL.Path == "/auth" {
		config.Templates.Exec(resp, "login.html", nil)
	} else {
		http.Error(resp, "404 - page not found", 404)
		return
	}

}
