package modules

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/app/modules/errors"
)

type Connection struct {
	Resp   http.ResponseWriter
	Req    *http.Request
	Path   []string
	UserId int
}

func (conn *Connection) GetUserId(forumDB *sql.DB) bool {
	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		return false
	}

	err = forumDB.QueryRow("SELECT id FROM users WHERE token=?", cookie.Value).Scan(&conn.UserId)
	if err != nil {
		return false
	}
	return true
}

func (conn *Connection) IsAuthenticated(forumDB *sql.DB) bool {
	validToken := conn.GetUserId(forumDB)
	if !validToken {
		conn.Error(errors.HttpUnauthorized)
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
