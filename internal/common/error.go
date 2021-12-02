package common

import "net/http"

type (
	// Error interface
	Error interface {
		error
		Status() int
	}

	// StatusError struct is status error
	StatusError struct {
		Code int
		Err  error
	}
)

// NewStatusError initializes new status err
func NewStatusError(code int, err error) error {
	return StatusError{
		Code: code,
		Err:  err,
	}
}

// Status returns http status code
func (e StatusError) Status() int {
	return e.Code
}

// Error returns error message
func (e StatusError) Error() string {
	return e.Err.Error()
}

func ThrowBadRequestError(err error) StatusError {
	return StatusError{
		Code: http.StatusBadRequest,
		Err:  err,
	}
}
