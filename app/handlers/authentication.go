package handlers

import "net/http"

func Auth(resp http.ResponseWriter, req *http.Request) {
	switch req.URL.Path[6:] {
	case "register":
	case "login":
	}
}
