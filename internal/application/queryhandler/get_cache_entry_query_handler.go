package queryhandler

import (
	"github.com/ybalcin/storage-api/internal/domain/cache"
	"github.com/ybalcin/storage-api/internal/domain/error"
)

type GetCacheEntryQuery struct {
	Key string
}

type GetCacheEntryQueryResult struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetCacheEntryQueryHandler struct {
	repo cache.Repository
}

func NewGetCacheEntryQueryHandler(repo cache.Repository) *GetCacheEntryQueryHandler {
	return &GetCacheEntryQueryHandler{repo: repo}
}

func (h *GetCacheEntryQueryHandler) Handle(q *GetCacheEntryQuery) (*GetCacheEntryQueryResult, *error.DomainErrorBase) {
	if q.Key == "" {
		return nil, error.NewDomainErrorBase(error.Invalid, "queryhandler: cache entry key is empty")
	}

	entry, err := h.repo.GetEntry(q.Key)
	if err != nil {
		return nil, err
	}

	return &GetCacheEntryQueryResult{
		Key:   entry.Key(),
		Value: entry.Value(),
	}, nil
}
