package errors

import "net/http"

type ResponseStatus int

const (
	Success             = ResponseStatus(http.StatusOK)
	BadRequest          = ResponseStatus(http.StatusBadRequest)
	InternalServerError = ResponseStatus(http.StatusInternalServerError)
)
