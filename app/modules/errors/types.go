package errors

import "net/http"

type HttpError struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

var HttpNotFound = &HttpError{
	http.StatusNotFound, http.StatusNotFound,
	"404 - page not found",
	"The Page you are trying to access doesn't not exists",
}

var HttpMethodNotAllowed = &HttpError{
	http.StatusMethodNotAllowed, http.StatusMethodNotAllowed,
	"405 - method not allowed",
	"The Method you are using is not supported in this endpoint",
}
