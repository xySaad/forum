package handlers

import (
	"net/http"
	"os"
)

func Static(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "405 - method not allowed", 405)
		return
	}

	fileInfo, err := os.Stat(req.URL.Path[1:])
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(resp, "404 - page not found", http.StatusNotFound)
			return
		}
		http.Error(resp, "500 - internal server error", 500)
		return
	}
	if fileInfo.IsDir() {
		http.Error(resp, "403 - access forbidden", 403)
		return
	}

	http.ServeFile(resp, req, req.URL.Path[1:])
}
