package mgo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection is mongo collection
type Collection struct {
	collection *mongo.Collection
}

func ToObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

// Collection returns mongo collection
func (m *Store) Collection(collection string) *Collection {
	checkAgainstNil(m)

	return &Collection{m.db.Collection(collection)}
}

// InsertOne inserts new document to collection
func (c *Collection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	if c == nil || c.collection == nil {
		return nil, errors.New("mgo: collection is nil")
	}

	return c.collection.InsertOne(ctx, document)
}

// Find finds one or many document
func (c *Collection) Find(ctx context.Context, filter interface{}, decodeModels interface{}) error {
	if c == nil || c.collection == nil {
		return errors.New("mgo: collection is nil")
	}

	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return err
	}

	defer cursorClose(cursor, ctx)

	err = cursor.All(ctx, decodeModels)
	if err != nil {
		return err
	}

	return nil
}

// Aggregate aggregates collection
func (c *Collection) Aggregate(ctx context.Context, pipeline interface{}, decodeModels interface{}) error {
	if c == nil || c.collection == nil {
		return errors.New("mgo: collection is nil")
	}

	cursor, err := c.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}

	defer cursorClose(cursor, ctx)

	err = cursor.All(ctx, decodeModels)
	if err != nil {
		return err
	}

	return nil
}

func cursorClose(cursor *mongo.Cursor, ctx context.Context) {
	err := cursor.Close(ctx)
	if err != nil {
		panic(err)
	}
}
