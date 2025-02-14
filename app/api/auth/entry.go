package auth

import (
	"database/sql"
	"net/http"

	"forum/app/modules"
	"forum/app/modules/errors"
)

func Entry(conn *modules.Connection, forumDB *sql.DB) {
	req := conn.Req
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
		Logout(conn, forumDB)
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
