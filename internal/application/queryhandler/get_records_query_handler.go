package queryhandler

import (
	"context"
	"github.com/ybalcin/storage-api/internal/common"
	"github.com/ybalcin/storage-api/internal/domain/error"
	"github.com/ybalcin/storage-api/internal/domain/record"
	"time"
)

// GetRecordsQuery struct
type GetRecordsQuery struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type GetRecordQueryResult struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

// GetRecordsQueryHandler handles get record query
type GetRecordsQueryHandler struct {
	repository record.Repository
}

// NewGetRecordQueryHandler initializes new get record query handler
func NewGetRecordQueryHandler(repository record.Repository) *GetRecordsQueryHandler {
	return &GetRecordsQueryHandler{repository: repository}
}

// Handle handles query
func (h *GetRecordsQueryHandler) Handle(ctx context.Context, q *GetRecordsQuery) ([]GetRecordQueryResult, *error.DomainErrorBase) {
	if err := q.validate(); err != nil {
		return nil, err
	}

	records, err := h.repository.GetRecords(ctx, q.StartDate, q.EndDate, q.MinCount, q.MaxCount)
	if err != nil {
		return nil, err
	}

	var queryResult []GetRecordQueryResult

	for _, m := range records {
		queryResult = append(queryResult, GetRecordQueryResult{
			Key:        m.Key(),
			CreatedAt:  m.CreatedAt(),
			TotalCount: m.TotalCount(),
		})
	}

	return queryResult, nil
}

func (q *GetRecordsQuery) validate() *error.DomainErrorBase {
	if q.StartDate == "" {
		return error.NewDomainErrorBase(error.Invalid, "queryhandler: start date is empty")
	}
	if _, err := common.StringToTime(q.StartDate); err != nil {
		return error.NewDomainErrorBase(error.Invalid, "queryhandler: start date is invalid format")
	}
	if q.EndDate == "" {
		return error.NewDomainErrorBase(error.Invalid, "queryhandler: end date is empty")
	}
	if _, err := common.StringToTime(q.EndDate); err != nil {
		return error.NewDomainErrorBase(error.Invalid, "queryhandler: end date is invalid format")
	}
	if q.MinCount <= 0 {
		return error.NewDomainErrorBase(error.Invalid, "queryhandler: min count must be bigger than zero")

	}
	if q.MaxCount <= 0 {
		return error.NewDomainErrorBase(error.Invalid, "queryhandler: max count must be bigger than zero")

	}
	if q.MinCount > q.MaxCount {
		return error.NewDomainErrorBase(error.Invalid, "queryhandler: min count cannot be smaller than max count")

	}

	return nil
}
