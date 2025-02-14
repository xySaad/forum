package auth

import (
	"database/sql"
	"net/http"
	"time"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
)

func Entry(conn *modules.Connection, forumDB *sql.DB) {
	req := conn.Req
	resp := conn.Resp
	if len(conn.Path) != 3 {
		conn.Error(errors.HttpNotFound)
		return
	}
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
		cookie, err := conn.Req.Cookie("token")
		if err != nil || cookie.Value == "" {
			conn.Error(errors.HttpUnauthorized)
			return
		}

		_, err = forumDB.Exec("DELETE FROM sessions WHERE token=?", cookie.Value)
		if err != nil {
			log.Error(err)
			conn.Error(errors.HttpInternalServerError)
			return
		}

		newCookie := http.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(resp, &newCookie)
		http.Redirect(resp, req, "/", http.StatusSeeOther)
	case "session":
		if conn.IsAuthenticated(forumDB) {
			conn.Resp.Write([]byte{'o', 'k'})
		}
		return

	default:
		conn.Error(errors.HttpNotFound)
		return
	}
}
