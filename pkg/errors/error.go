package errors

import "net/http"

type BaseError struct {
	msg    string // errors description
	Offset int64  // where the errors occurred
	Code   int
}

func (e *BaseError) Error() string { return e.msg }

func InternalError(message string) *BaseError {
	return &BaseError{msg: message, Code: http.StatusInternalServerError}
}
