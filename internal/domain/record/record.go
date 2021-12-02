package record

import (
	"github.com/ybalcin/storage-api/internal/domain/error"
	"time"
)

// Model struct
type Model struct {
	id         Id
	key        string
	value      string
	createdAt  time.Time
	totalCount int
}

// New initializes new record
func New(id Id, key, value string, createdAt time.Time, totalCount int) (*Model, *error.DomainErrorBase) {
	r := &Model{
		id:         id,
		key:        key,
		value:      value,
		createdAt:  createdAt,
		totalCount: totalCount,
	}

	if err := r.validate(); err != nil {
		return nil, err
	}

	return r, nil
}

// Id returns id of record
func (r *Model) Id() Id {
	return r.id
}

// Key returns key of record
func (r *Model) Key() string {
	return r.key
}

// Value returns value of record
func (r *Model) Value() string {
	return r.value
}

// CreatedAt returns created at date of record
func (r *Model) CreatedAt() time.Time {
	return r.createdAt
}

// TotalCount returns total count of record
func (r *Model) TotalCount() int {
	return r.totalCount
}

// validate validates record
func (r *Model) validate() *error.DomainErrorBase {
	if r.id == "" {
		return error.NewDomainErrorBase(error.Invalid, "record id is empty")
	}
	if r.key == "" {
		return error.NewDomainErrorBase(error.Invalid, "record key is empty")
	}
	if r.value == "" {
		return error.NewDomainErrorBase(error.Invalid, "record value is empty")
	}

	return nil
}
