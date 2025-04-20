package modules

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/app/modules/errors"
	"forum/app/modules/log"
)

type Connection struct {
	Resp http.ResponseWriter
	Req  *http.Request
	Path []string
	User User
}

const tokenQuery = `SELECT u.id FROM users u
	JOIN sessions s ON s.user_id=u.id 
	WHERE s.token=? AND s.expires_at > datetime('now')`

func (conn *Connection) GetUserId(forumDB *sql.DB) bool {
	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		return false
	}

	return forumDB.QueryRow(tokenQuery, cookie.Value).Scan(&conn.User) != nil
}

func (conn *Connection) IsAuthenticated(forumDB *sql.DB) bool {
	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		conn.Error(errors.HttpUnauthorized)
		return false
	}

	err = forumDB.QueryRow(tokenQuery, cookie.Value).Scan(&conn.User.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			conn.Error(errors.HttpUnauthorized)
		} else {
			log.Error(err)
			conn.Error(errors.HttpInternalServerError)
		}
		return false
	}
	return true
}

func (conn *Connection) NewError(httpStatus, code int, message, details string) {
	httpError := errors.HttpError{
		Code:    code,
		Message: message,
		Details: details,
		Status:  httpStatus,
	}

	sendHttpError(conn, &httpError)
}

func (conn *Connection) Error(httpError *errors.HttpError) {
	sendHttpError(conn, httpError)
}

func (conn *Connection) Respond(data any) {
	jsonResult, err := json.Marshal(data)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		return
	}
	conn.Resp.Header().Set("Content-Type", "application/json")
	conn.Resp.Write(jsonResult)
}

func sendHttpError(conn *Connection, httpError *errors.HttpError) {
	conn.Resp.Header().Set("Content-Type", "application/json")
	conn.Resp.WriteHeader(httpError.Status)

	jsonError, err := json.Marshal(httpError)
	if err != nil {
		conn.Resp.Write([]byte(httpError.Message))
		return
	}

	conn.Resp.Write(jsonError)
}
