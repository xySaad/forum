package modules

import (
	"encoding/json"
	"forum/app/modules/errors"
	"net/http"
)

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

type Connection struct {
	Resp http.ResponseWriter
	Req  *http.Request
	Path []string
}
