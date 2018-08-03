package status

import (
	"fmt"
	"net/http"
)

// ErrInternal represents internal server error.
var ErrInternal = ErrServiceStatus{
	ServiceStatus{Code: http.StatusInternalServerError, Message: "Internal Server Error"},
}

// ErrNotFound represents an error when a domain artifact was not found.
var ErrNotFound = ErrServiceStatus{
	ServiceStatus{Code: http.StatusNotFound, Message: "Not Found"},
}

// ErrBadRequest represents an invalid request error.
var ErrBadRequest = ErrServiceStatus{
	ServiceStatus{Code: http.StatusBadRequest, Message: "Bad Request"},
}

// ErrUnauhtorized represents an unauthorized request error.
var ErrUnauhtorized = ErrServiceStatus{
	ServiceStatus{Code: http.StatusUnauthorized, Message: "Unauthorized"},
}

// Success represents a generic success.
var Success = ServiceStatus{Code: http.StatusOK, Message: "OK"}

// ServiceStatus captures basic information about a status construct.
type ServiceStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrServiceStatus captures basic information about an error.
type ErrServiceStatus struct {
	ServiceStatus
}

// WithMessage returns an error status with given message.
func (e ErrServiceStatus) WithMessage(msg string) ErrServiceStatus {
	return ErrServiceStatus{ServiceStatus{Code: e.Code, Message: msg}}
}

// WithError returns an error status with given err.Error().
func (e ErrServiceStatus) WithError(err error) ErrServiceStatus {
	return ErrServiceStatus{ServiceStatus{Code: e.Code, Message: err.Error()}}
}

// New returns a new status with given status instance.
func New(ss ServiceStatus) ServiceStatus {
	return ServiceStatus{ss.Code, ss.Message}
}

// NewError returns a new status with given status instance.
func NewErrorStatus(err ErrServiceStatus) ServiceStatus {
	return ServiceStatus{err.Code, err.Message}
}

// NewUserDefined returns a new status with given code and message.
func NewUserDefined(code int, msg string) ServiceStatus {
	return ServiceStatus{Code: code, Message: msg}
}

func (e ErrServiceStatus) Error() string {
	return fmt.Sprintf(string(e.Code), ": ", e.Message)
}
