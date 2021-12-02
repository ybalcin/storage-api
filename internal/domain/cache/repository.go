package cache

import (
	"github.com/ybalcin/storage-api/internal/domain/error"
)

// Repository interface wraps entry operations
type Repository interface {
	// AddEntry adds cache entry
	AddEntry(entry *Entry) *error.DomainErrorBase

	// GetEntry gets cache entry
	GetEntry(key string) (*Entry, *error.DomainErrorBase)
}
