package api

type ErrorCode string

const (
	ErrNotFound      ErrorCode = "NOT_FOUND"
	ErrAlreadyExists ErrorCode = "ALREADY_EXISTS"
	ErrConflict      ErrorCode = "CONFLICT"
	ErrInvalidInput  ErrorCode = "INVALID_INPUT"
	ErrInternal      ErrorCode = "INTERNAL_ERROR"
	ErrBadRequest    ErrorCode = "BAD_REQUEST"
)

type APIError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e APIError) Error() string { return e.Message }

var (
	SecretNotFoundErr          = APIError{Code: ErrNotFound, Message: "secret not found"}
	SecretVersionConflictErr   = APIError{Code: ErrConflict, Message: "secret version conflict"}
	SecretNameAlreadyExistsErr = APIError{Code: ErrAlreadyExists, Message: "secret name already exists"}
	SecretDeletedErr           = APIError{Code: ErrConflict, Message: "secret is deleted"}
	InvalidCiphertextErr       = APIError{Code: ErrInvalidInput, Message: "invalid encrypted data format"}
	InternalErr                = APIError{Code: ErrInternal, Message: "internal server error"}
)
