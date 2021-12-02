package adapter

import (
	"context"
	"github.com/ybalcin/storage-api/internal/domain/error"
	"github.com/ybalcin/storage-api/internal/domain/record"
	"github.com/ybalcin/storage-api/pkg/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// mongoRecordAdapter struct implements record.Repository
type mongoRecordAdapter struct {
	records *mgo.Collection
}

type recordBson struct {
	Id        primitive.ObjectID `bson:"_id"`
	Key       string             `bson:"key"`
	Value     string             `bson:"value"`
	CreatedAt time.Time          `bson:"createdAt"`
	Counts    []int              `bson:"counts"`
}

const recordCollection = "records"

// NewRecordMongoAdapter initializes new record mongo adapter
func NewRecordMongoAdapter(store *mgo.Store) (*mongoRecordAdapter, *error.DomainErrorBase) {
	if store == nil {
		return nil, error.NewDomainErrorBase(error.Invalid, "record mongo adapter store is nil")
	}

	collection := store.Collection(recordCollection)

	return &mongoRecordAdapter{records: collection}, nil
}

// GetRecords gets records by filter
func (a *mongoRecordAdapter) GetRecords(ctx context.Context, startDate, endDate string, minCount, maxCount int) ([]record.Model, *error.DomainErrorBase) {

	var recordsBson []recordBson
	var records []record.Model

	addFieldsStage := bson.D{
		{"$addFields", bson.M{
			"onlyDate": bson.M{
				"$dateToString": bson.M{
					"format": "%Y-%m-%d",
					"date":   "$createdAt",
				},
			},
			"totalCount": bson.M{
				"$sum": "$counts",
			},
		}},
	}

	matchStage := bson.D{
		{"$match", bson.M{
			"onlyDate": bson.M{
				"$lt": endDate,
				"$gt": startDate,
			},
			"totalCount": bson.M{
				"$lt": maxCount,
				"$gt": minCount,
			},
		}},
	}

	err := a.records.Aggregate(ctx, mongo.Pipeline{addFieldsStage, matchStage}, &recordsBson)
	if err != nil {
		return nil, error.NewDomainErrorBase(error.Invalid, err.Error())
	}

	for _, m := range recordsBson {
		totalCount := sum(m.Counts...)
		r, err := record.New(record.Id(m.Id.Hex()), m.Key, m.Value, m.CreatedAt, totalCount)
		if err != nil {
			continue
		}

		records = append(records, *r)
	}

	return records, nil
}

func sum(values ...int) int {
	var total int
	for _, v := range values {
		total += v
	}

	return total
}
