package auth

import (
	"forum/app/modules"
	"net/http"
	"time"
)

func Entry(conn *modules.Connection) {
	req := conn.Req
	resp := conn.Resp
	switch req.URL.Path[10:] {
	case "register":
		if req.Method != http.MethodPost {
			http.Error(resp, "405 - method not allowed", http.StatusMethodNotAllowed)
			return
		}
		Register(conn)
	case "login":
		if req.Method != http.MethodPost {
			http.Error(resp, "405 - method not allowed", http.StatusMethodNotAllowed)
		}
		err := LogIn(req.Body, resp)
		if err != nil {
			http.Error(resp, "500 - "+err.Error(), 500)
			return
		}
	case "logout":
		if req.Method != http.MethodPost {
			http.Error(resp, "405 - method not allowed", http.StatusMethodNotAllowed)
			return
		}
		cookie := http.Cookie{
			Name:     "ticket",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(resp, &cookie)
		// then redirect to / or whatever
	case "auth":
		cookie, err := req.Cookie("ticket")
		if err != nil {
			// unothorized
			return
		}
		if err := CheckAuth(cookie.Value); err != nil {
			// unothorized
			return
		}
	default:
		http.Error(conn.Resp, "404 - page not found", 404)
		return
	}

}
