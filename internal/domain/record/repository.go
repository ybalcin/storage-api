package record

import (
	"context"
	"github.com/ybalcin/storage-api/internal/domain/error"
)

// Repository interface wraps record operations
type Repository interface {
	GetRecords(ctx context.Context, startDate, endDate string, minCount, maxCount int) ([]Model, *error.DomainErrorBase)
}
