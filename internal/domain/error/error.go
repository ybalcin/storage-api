// Package error provides domain error utilities
package error

import "fmt"

type code int

const (
	Success code = iota
	NotExist
	Invalid
)

// DomainErrorBase is domain error
type DomainErrorBase struct {
	Code    code
	Message string
}

// NewDomainErrorBase initializes new error
func NewDomainErrorBase(code code, message string) *DomainErrorBase {
	return &DomainErrorBase{
		Code:    code,
		Message: message,
	}
}

func (e *DomainErrorBase) String() string {
	return fmt.Sprintf("error occured with message: %s and code: %d", e.Message, e.Code)
}

func (e *DomainErrorBase) Int() int {
	return int(e.Code)
}
