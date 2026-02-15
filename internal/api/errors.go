package api

type ErrorCode string

const (
	ErrNotFound      ErrorCode = "NOT_FOUND"
	ErrAlreadyExists ErrorCode = "ALREADY_EXISTS"
	ErrInvalidInput  ErrorCode = "INVALID_INPUT"
	ErrInternal      ErrorCode = "INTERNAL_ERROR"
	ErrorBadRequest  ErrorCode = "BAD_REQUEST"
)

type APIError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e APIError) Error() string { return e.Message }

var (
	SecretNotFoundErr = APIError{Code: ErrNotFound, Message: "secret not found"}
	SecretExistsErr   = APIError{Code: ErrAlreadyExists, Message: "secret already exists"}
	InternalErr       = APIError{Code: ErrInternal, Message: "internal server error"}
)
