package errorh

import "errors"

// standard repository errors contain implementation details or potentially sensitive information.
// never returned to clients, but can be logged for debugging purposes.
var (
	ErrNilPointer error = errors.New("nil pointer passed as argument")
)

type ErrorCode int

const (
	ErrorCodeUnknown ErrorCode = iota
	ErrorCodeValidation
	ErrorCodeNotFound
	ErrorCodeUnauthorized
	ErrorCodeForbidden
	ErrorCodeBadRequest
	ErrorCodeInternal
)

// services, handlers validators package their errors using error struct. logging.
type Error struct {
	source  error
	Code    ErrorCode
	Message string
	Fields  map[string][]string
}

func New(code ErrorCode) *Error {
	return &Error{
		Code:    code,
		Message: getDefaultMessage(code),
	}
}

func getDefaultMessage(code ErrorCode) string {
	switch code {
	case ErrorCodeUnknown:
		return "An unknown error occurred"
	case ErrorCodeValidation:
		return "The provided data is invalid"
	case ErrorCodeNotFound:
		return "The requested resource was not found"
	case ErrorCodeUnauthorized:
		return "Authentication is required to access this resource"
	case ErrorCodeForbidden:
		return "You don't have permission to access this resource"
	case ErrorCodeBadRequest:
		return "The request could not be processed"
	case ErrorCodeInternal:
		return "An internal server error occurred"
	default:
		return "An unexpected error occurred"
	}
}

