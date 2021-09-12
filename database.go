package dalgo_firestore

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
	client *firestore.Client
}

var _ dalgo.Database = (*database)(nil)

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
	return dtb
}

func (dtb database) doc(key dalgo.RecordKey) *firestore.DocumentRef {
	path := dalgo.GetRecordKeyPath(key)
	return dtb.client.Doc(path)
}

func (dtb database) Upsert(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}
