package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewNotFound(message string, v ...any) *Error {
	return &Error{Code: http.StatusNotFound, Message: fmt.Sprintf(message, v...)}
}

func NewBadRequest(message string, v ...any) *Error {
	return &Error{Code: http.StatusBadRequest, Message: fmt.Sprintf(message, v...)}
}

func NewInternalServerError(message string, v ...any) *Error {
	return &Error{Code: http.StatusInternalServerError, Message: fmt.Sprintf(message, v...)}
}

func NewUnauthorized(message string, v ...any) *Error {
	return &Error{Code: http.StatusUnauthorized, Message: fmt.Sprintf(message, v...)}
}

func NewForbidden(message string, v ...any) *Error {
	return &Error{Code: http.StatusForbidden, Message: fmt.Sprintf(message, v...)}
}

func NewConflict(message string, v ...any) *Error {
	return &Error{Code: http.StatusConflict, Message: fmt.Sprintf(message, v...)}
}

func NewUnprocessableEntity(message string, v ...any) *Error {
	return &Error{Code: http.StatusUnprocessableEntity, Message: fmt.Sprintf(message, v...)}
}

func NewServiceUnavailable(message string, v ...any) *Error {
	return &Error{Code: http.StatusServiceUnavailable, Message: fmt.Sprintf(message, v...)}
}

func NewNotImplemented(message string, v ...any) *Error {
	return &Error{Code: http.StatusNotImplemented, Message: fmt.Sprintf(message, v...)}
}

func NewBadGateway(message string, v ...any) *Error {
	return &Error{Code: http.StatusBadGateway, Message: fmt.Sprintf(message, v...)}
}

func NewGatewayTimeout(message string, v ...any) *Error {
	return &Error{Code: http.StatusGatewayTimeout, Message: fmt.Sprintf(message, v...)}
}

func NewHTTPError(code int, message string, v ...any) *Error {
	return &Error{Code: code, Message: fmt.Sprintf(message, v...)}
}
