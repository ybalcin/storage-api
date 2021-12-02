package application

import (
	"context"
	"github.com/ybalcin/storage-api/internal/application/commandhandler"
	"github.com/ybalcin/storage-api/internal/application/queryhandler"
	"github.com/ybalcin/storage-api/internal/infrastructure/adapter"
	"github.com/ybalcin/storage-api/pkg/inmemorystore"
	"github.com/ybalcin/storage-api/pkg/mgo"
)

// Application provides access to application
type Application struct {
	commands *commands
	queries  *queries
}

type commands struct {
	AddCacheEntryCommand *commandhandler.AddCacheEntryCommandHandler
}

type queries struct {
	GetRecordsQuery    *queryhandler.GetRecordsQueryHandler
	GetCacheEntryQuery *queryhandler.GetCacheEntryQueryHandler
}

// Queries returns queries of app
func (a *Application) Queries() *queries {
	return a.queries
}

// Commands returns commands of app
func (a *Application) Commands() *commands {
	return a.commands
}

// New initializes new application
func New() *Application {
	ctx := context.Background()
	// from os env?
	store := mgo.NewStore(ctx, "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true", "getir-case-study")
	inMemoryClient := inmemorystore.NewClient()

	mongoAdapter, err := adapter.NewRecordMongoAdapter(store)
	if err != nil {
		panic(err)
	}

	inmemoryAdapter := adapter.NewCacheInMemoryAdapter(inMemoryClient)

	queries := &queries{
		GetRecordsQuery:    queryhandler.NewGetRecordQueryHandler(mongoAdapter),
		GetCacheEntryQuery: queryhandler.NewGetCacheEntryQueryHandler(inmemoryAdapter),
	}

	commands := &commands{
		AddCacheEntryCommand: commandhandler.NewAddCacheEntryCommandHandler(inmemoryAdapter),
	}

	return &Application{
		commands: commands,
		queries:  queries,
	}
}
