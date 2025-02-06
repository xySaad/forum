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
			conn.Error(errors.HttpMethodNotAllowed)
			return
		}
		Register(conn, forumDB)
	case "login":
		if req.Method != http.MethodPost {
			conn.Error(errors.HttpMethodNotAllowed)
			return
		}
		LogIn(conn, forumDB)
	case "logout":
		if req.Method != http.MethodPost {
			conn.Error(errors.HttpMethodNotAllowed)
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
			conn.Error(errors.HttpUnauthorized)
			return
		}

		if CheckAuth(cookie.Value, forumDB) != nil {
			conn.Error(errors.HttpUnauthorized)
			return
		}

	default:
		conn.Error(errors.HttpNotFound)
		return
	}
}
