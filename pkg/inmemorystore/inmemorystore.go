// Package inmemorystore is a simple library for storing key-value pair in memory
package inmemorystore

import (
	"fmt"
	"strings"
	"sync"
)

type (
	storage map[string]string

	client struct {
	}

	// Client is interface that wraps the in memory cache operations
	Client interface {
		// AddToMemory sets a key-value pair.
		AddToMemory(key string, value string) error

		// GetFromMemory retrieves a value by key.
		GetFromMemory(key string) (string, error)
	}

	cacheItem struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	cacheItems struct {
		Data []cacheItem `json:"data"`
	}
)

var (
	lock  = sync.Mutex{}
	cache = storage{}

	ErrEmptyKey   = newPackageError("key is empty")
	ErrEmptyValue = newPackageError("value is empty")
)

func newPackageError(message string) error {
	return fmt.Errorf("inmemorystore: %s", message)
}

// NewClient initializes new inmemorystore client.
// param interval: cache file saving interval time as minutes
func NewClient() Client {
	return &client{}
}

// AddToMemory sets a key-value pair.
func (c *client) AddToMemory(key string, value string) error {
	if key == "" {
		return ErrEmptyKey
	}
	if value == "" {
		return ErrEmptyValue
	}

	lock.Lock()
	defer lock.Unlock()

	if cache == nil {
		cache = storage{}
	}

	key = strings.ReplaceAll(key, " ", "")
	cache[key] = value
	return nil
}

// GetFromMemory retrieves a value by key.
func (c *client) GetFromMemory(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}

	lock.Lock()
	defer lock.Unlock()

	val, ok := cache[key]
	if !ok {
		return "", nil
	}

	return val, nil
}
