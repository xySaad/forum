package auth

import (
	"database/sql"
	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
	"net/http"
	"time"
)

func Logout(conn *modules.Connection, db *sql.DB) {
	if conn.Req.Method != http.MethodPost {
		conn.Error(errors.HttpMethodNotAllowed)
		return
	}
	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		conn.Error(errors.HttpUnauthorized)
		return
	}

	_, err = db.Exec("UPDATE sessions SET expires_at=0 WHERE token=?", cookie.Value)
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
	http.SetCookie(conn.Resp, &newCookie)
	http.Redirect(conn.Resp, conn.Req, "/", http.StatusSeeOther)
}
