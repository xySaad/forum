package auth

import (
	"net/http"
	"time"
)

func Entry(resp http.ResponseWriter, req *http.Request) {
	switch req.URL.Path[6:] {
	case "register":
		if req.Method != http.MethodPost {
			return
		}
		err := Register(req.Body, resp)
		if err != nil {
			// send the err message to front
			return
		}
	case "login":
		if req.Method != http.MethodGet {
			return
		}
		err := LogIn(req.Body, resp)
		if err != nil {
			// send the err message to front
			return
		}
	case "logout":
		if req.Method != http.MethodPost {
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
	}

}
