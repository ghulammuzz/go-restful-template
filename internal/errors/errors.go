package errors

import "net/http"

type ErrorCode string

const (
	ErrUsernameRequired ErrorCode = "USERNAME_REQUIRED"
	ErrPasswordRequired ErrorCode = "PASSWORD_REQUIRED"
	ErrPayload          ErrorCode = "PAYLOAD_ERROR"
	ErrUsernameTaken    ErrorCode = "USERNAME_TAKEN_ERROR"
	ErrHashingFailed    ErrorCode = "HASHING_FAILED"
	ErrInternal         ErrorCode = "INTERNAL_ERROR"
)

var ErrorMapping = map[ErrorCode]int{
	ErrUsernameRequired: http.StatusBadRequest,
	ErrPasswordRequired: http.StatusBadRequest,
	ErrPayload:          http.StatusBadRequest,
	ErrUsernameTaken:    http.StatusBadRequest,
	ErrHashingFailed:    http.StatusInternalServerError,
	ErrInternal:         http.StatusInternalServerError,
}

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func GetStatusCode(errorCode ErrorCode) int {
	return ErrorMapping[errorCode]
}

func NewAppError(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Error mengimplementasikan interface error
func (e *AppError) Error() string {
	return e.Message
}
