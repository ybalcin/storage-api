package adapter

import (
	"github.com/ybalcin/storage-api/internal/domain/cache"
	"github.com/ybalcin/storage-api/internal/domain/error"
	"github.com/ybalcin/storage-api/pkg/inmemorystore"
)

// cacheInMemoryAdapter implements cache.Repository
type cacheInMemoryAdapter struct {
	cli inmemorystore.Client
}

// NewCacheInMemoryAdapter initializes cache inmemory adapter
func NewCacheInMemoryAdapter(cli inmemorystore.Client) *cacheInMemoryAdapter {
	return &cacheInMemoryAdapter{cli: cli}
}

// AddEntry adds cache entry to inmemory store
func (c *cacheInMemoryAdapter) AddEntry(entry *cache.Entry) *error.DomainErrorBase {
	if err := c.cli.AddToMemory(entry.Key(), entry.Value()); err != nil {
		return error.NewDomainErrorBase(error.Invalid, err.Error())
	}

	return nil
}

// GetEntry gets entry from in memory store
func (c *cacheInMemoryAdapter) GetEntry(key string) (*cache.Entry, *error.DomainErrorBase) {
	val, err := c.cli.GetFromMemory(key)
	if err != nil {
		return nil, error.NewDomainErrorBase(error.Invalid, err.Error())
	}

	entry, domainErr := cache.NewEntry(key, val)
	if domainErr != nil {
		return nil, domainErr
	}

	return entry, nil
}
