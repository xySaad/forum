package auth

import (
	"database/sql"
	"net/http"
	"time"

	"forum/app/modules"
	"forum/app/modules/errors"
)

func Entry(conn *modules.Connection, forumDB *sql.DB) {
	req := conn.Req
	resp := conn.Resp

	switch conn.Path[2] {
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
		LogIn(conn, forumDB)
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

		if CheckAuth(cookie.Value, forumDB) != nil {
			http.Error(resp, "Unauthorized", http.StatusUnauthorized)
			return
		}

	default:
		conn.Error(errors.HttpNotFound)
		return
	}
}
