// Package cache domain
package cache

import (
	"github.com/ybalcin/storage-api/internal/domain/error"
)

// Entry struct is cache entry
type Entry struct {
	key   string
	value string
}

// NewEntry initializes new cache entry
func NewEntry(key, value string) (*Entry, *error.DomainErrorBase) {
	e := &Entry{
		key:   key,
		value: value,
	}

	if err := e.validate(); err != nil {
		return nil, err
	}

	return e, nil
}

// Key returns key of entry
func (e *Entry) Key() string {
	return e.key
}

// Value returns value of entry
func (e *Entry) Value() string {
	return e.value
}

// validate validates entry
func (e *Entry) validate() *error.DomainErrorBase {
	if e.key == "" {
		return error.NewDomainErrorBase(error.Invalid, "entry key is empty")
	}
	if e.value == "" {
		return error.NewDomainErrorBase(error.Invalid, "entry value is empty")
	}

	return nil
}
