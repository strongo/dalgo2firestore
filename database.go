package dalgo2firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo"
)

type database struct {
	inserter
	deleter
	getter
	setter
	updater
	client *firestore.Client
}

func (dtb database) Select(ctx context.Context, query dalgo.Query) (dalgo.Reader, error) {
	panic("implement me")
}

var _ dalgo.Database = (*database)(nil)

// NewDatabase creates new instance of dalgo interface to Firestore
func NewDatabase(client *firestore.Client) dalgo.Database {
	if client == nil {
		panic("client is a required field, got nil")
	}
	dtb := database{
		client: client,
	}
	dtb.inserter = newInserter(dtb)
	dtb.deleter = newDeleter(dtb)
	dtb.getter = newGetter(dtb)
	dtb.setter = newSetter(dtb)
	dtb.updater = newUpdater(&dtb)
	return dtb
}

func (dtb database) doc(key *dalgo.Key) *firestore.DocumentRef {
	path := PathFromKey(key)
	return dtb.client.Doc(path)
}

func (dtb database) Upsert(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}
