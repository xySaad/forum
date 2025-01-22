package auth

import (
	"database/sql"
	"forum/app/modules"
	"net/http"
	"time"
)

func Entry(conn *modules.Connection, forumDB *sql.DB) {
	req := conn.Req
	resp := conn.Resp

	switch req.URL.Path[10:] {
	case "register":
		if req.Method != http.MethodPost {
			http.Error(resp, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		Register(conn, forumDB)

	case "login":
		if req.Method != http.MethodPost {
			http.Error(resp, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		err := LogIn(req.Body, resp, forumDB)
		if err != nil {
			http.Error(resp, "500 - "+err.Error(), http.StatusInternalServerError)
			return
		}

	case "logout":
		if req.Method != http.MethodPost {
			http.Error(resp, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		cookie := http.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(resp, &cookie)
		http.Redirect(resp, req, "/", http.StatusSeeOther)

	case "session":
		cookie, err := req.Cookie("token")
		if err != nil || cookie.Value == "" {
			http.Error(resp, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if err := CheckAuth(cookie.Value, forumDB); err != nil {
			http.Error(resp, "Unauthorized", http.StatusUnauthorized)
			return
		}

	default:
		http.Error(resp, "404 - Page Not Found", http.StatusNotFound)
		return
	}
}
