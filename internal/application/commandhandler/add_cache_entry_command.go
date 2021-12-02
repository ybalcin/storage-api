package commandhandler

import (
	"github.com/ybalcin/storage-api/internal/domain/cache"
	"github.com/ybalcin/storage-api/internal/domain/error"
)

type AddCacheEntryCommand struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AddCacheEntryCommandHandler struct {
	repo cache.Repository
}

// NewAddCacheEntryCommandHandler initializes new AddCacheEntryCommandHandler
func NewAddCacheEntryCommandHandler(repo cache.Repository) *AddCacheEntryCommandHandler {
	return &AddCacheEntryCommandHandler{repo: repo}
}

// Handle handles command
func (h *AddCacheEntryCommandHandler) Handle(c *AddCacheEntryCommand) *error.DomainErrorBase {
	entry, err := cache.NewEntry(c.Key, c.Value)
	if err != nil {
		return err
	}

	if err = h.repo.AddEntry(entry); err != nil {
		return err
	}

	return nil
}
