// Package mgo provides mongo client operations
package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
)

// Store is wraps *mongo.Database and *mongo.Client
type Store struct {
	db  *mongo.Database
	cli *mongo.Client
}

// NewStore connects mongo and initializes Store
func NewStore(ctx context.Context, uri, database string) *Store {
	checkAgainstEmpty("uri", uri)
	checkAgainstEmpty("database", database)

	db, cli := connect(ctx, uri, database)
	return &Store{
		db:  db,
		cli: cli,
	}
}

// IsConnected checks mongo is connected
func (m *Store) IsConnected() bool {
	checkAgainstNil(m)

	err := m.cli.Ping(context.Background(), readpref.Primary())
	return err == nil
}

func connect(ctx context.Context, uri, database string) (*mongo.Database, *mongo.Client) {
	var once sync.Once
	var cli *mongo.Client
	var db *mongo.Database

	once.Do(func() {
		db, cli = connectToMongo(ctx, uri, database)
	})

	return db, cli
}

func connectToMongo(ctx context.Context, uri, database string) (*mongo.Database, *mongo.Client) {
	cli, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fatal(err)
	}

	if err = cli.Connect(ctx); err != nil {
		fatal(err)
	}

	db := cli.Database(database)
	log.Printf("mgo: successfully connected to database %s", database)

	return db, cli
}
