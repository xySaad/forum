package errors

import "net/http"

type HttpError struct 
{
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func NewError(status, code int, msg, details string) *HttpError {
	return &HttpError{
		Status:  status,
		Code:    code,
		Message: msg,
		Details: details,
	}
}

var HttpNotFound = &HttpError{
	http.StatusNotFound, http.StatusNotFound,
	"404 - page not found",
	"The Page you are trying to access doesn't not exists",
}

var HttpUnauthorized = &HttpError{
	http.StatusUnauthorized, http.StatusUnauthorized,
	"unauthorized",
	"only loged-in members can performe this action",
}

var HttpMethodNotAllowed = &HttpError{
	http.StatusMethodNotAllowed, http.StatusMethodNotAllowed,
	"405 - method not allowed",
	"The Method you are using is not supported in this endpoint",
}

var HttpInternalServerError = &HttpError{
	http.StatusInternalServerError, http.StatusInternalServerError,
	"500 - internal server error",
	"Sorry something went wrong",
}
