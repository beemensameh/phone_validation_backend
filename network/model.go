package network

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code=%v message=%s", e.Code, e.Message)
}

func InternalServerError(code int, message string) (err *Error) {
	return newError(code, message)
}

func BadRequest(code int, message string) (err *Error) {
	return newError(code, message)
}

func NotFound(code int, message string) (err *Error) {
	return newError(code, message)
}

func Forbidden(code int, message string) (err *Error) {
	return newError(code, message)
}

func Unauthorized(code int, message string) (err *Error) {
	return newError(code, message)
}

func Invalid(code int, message string) (err *Error) {
	return newError(code, message)
}

func TooManyRequests(code int, message string) (err *Error) {
	return newError(code, message)
}

func BadGateway(code int, message string) (err *Error) {
	return newError(code, message)
}

func newError(code int, message string) (err *Error) {
	return &Error{
		Code:    code,
		Message: message,
	}
}
