package handlers

import "net/http"

func post(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(resp, "405 - method not allowed", 405)
		return
	}
}
