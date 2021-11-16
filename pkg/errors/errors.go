package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error implements the error interface.
type Error struct {
	Code int32
	Err  error
}

// ViewError : The error view.
type ViewError struct {
	Code    int32  `json:"faultCode"`
	Message string `json:"faultString"`
}

// MarshalJSON for error.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(ViewError{
		e.Code,
		e.Err.Error(),
	})
}

// Error : Error to string.
func (e *Error) Error() string {
	return e.Err.Error()
}

// New : Generates a custom error.
func New(code int32, format string, a ...interface{}) *Error {
	return &Error{
		Code: code,
		Err:  fmt.Errorf(format, a...),
	}
}

// BadRequest : Generates a 400 error.
func BadRequest(err error) *Error {
	return &Error{
		Code: http.StatusBadRequest,
		Err:  err,
	}
}

// BadRequestNew : Generates a 400 error from the message.
func BadRequestNew(format string, a ...interface{}) *Error {
	return &Error{
		Code: http.StatusBadRequest,
		Err:  fmt.Errorf(format, a...),
	}
}

// Unauthorized : Generates a 401 error.
func Unauthorized(err error) *Error {
	return &Error{
		Code: http.StatusUnauthorized,
		Err:  err,
	}
}

// UnauthorizedNew : Generates a 401 error from the message.
func UnauthorizedNew(format string, a ...interface{}) *Error {
	return &Error{
		Code: http.StatusUnauthorized,
		Err:  fmt.Errorf(format, a...),
	}
}

// Forbidden : Generates a 403 error.
func Forbidden(err error) *Error {
	return &Error{
		Code: http.StatusForbidden,
		Err:  err,
	}
}

// ForbiddenNew : Generates a 403 error from the message.
func ForbiddenNew(format string, a ...interface{}) *Error {
	return &Error{
		Code: http.StatusForbidden,
		Err:  fmt.Errorf(format, a...),
	}
}

// NotFound : Generates a 404 error.
func NotFound(err error) *Error {
	return &Error{
		Code: http.StatusNotFound,
		Err:  err,
	}
}

// NotFoundNew : Generates a 404 error from the message.
func NotFoundNew(format string, a ...interface{}) *Error {
	return &Error{
		Code: http.StatusNotFound,
		Err:  fmt.Errorf(format, a...),
	}
}

// MethodNotAllowed : Generates a 405 error.
func MethodNotAllowed(err error) *Error {
	return &Error{
		Code: http.StatusMethodNotAllowed,
		Err:  err,
	}
}

// MethodNotAllowedNew : Generates a 405 error from the message.
func MethodNotAllowedNew(format string, a ...interface{}) *Error {
	return &Error{
		Code: http.StatusMethodNotAllowed,
		Err:  fmt.Errorf(format, a...),
	}
}

// InternalServer : Generates a 500 error.
func InternalServer(err error) *Error {
	return &Error{
		Code: http.StatusInternalServerError,
		Err:  err,
	}
}

// InternalServerNew : Generates a 500 error from the message.
func InternalServerNew(format string, a ...interface{}) *Error {
	return &Error{
		Code: http.StatusInternalServerError,
		Err:  fmt.Errorf(format, a...),
	}
}
