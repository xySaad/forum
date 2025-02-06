package errors

import "net/http"

func NewError(status, code int, msg, details string) *HttpError {
	return &HttpError{status, code, msg, details}
}

func BadRequestError(details string) *HttpError {
	return NewError(http.StatusBadRequest, http.StatusBadRequest, "400 - bad request", details)
}
