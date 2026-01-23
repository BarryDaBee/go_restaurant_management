package utility

import "net/http"

type AppError struct {
	message    string
	statusCode int
}

var (
	ErrRecordNotFound = &AppError{message: "record not found", statusCode: http.StatusNotFound}
	ErrUnauthorized   = &AppError{message: "unauthorized", statusCode: http.StatusUnauthorized}
	ErrBadRequest     = &AppError{message: "bad request", statusCode: http.StatusBadRequest}
)

func NewAppError(message string, statusCode ...int) error {
	defaultStatusCode := http.StatusInternalServerError
	if len(statusCode) > 0 {
		defaultStatusCode = statusCode[0]
	}
	return &AppError{message: message, statusCode: defaultStatusCode}
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) StatusCode() int {
	return e.statusCode
}
