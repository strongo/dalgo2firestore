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
	client *firestore.Client
}

var _ dalgo.Database = (*database)(nil)

func NewDatabase() dalgo.Database {
	var dtb database
	dtb.inserter = newInserter(dtb)
	dtb.deleter = newDeleter(dtb)
	dtb.getter = newGetter(dtb)
	return dtb
}

func (dtb database) doc(key dalgo.RecordKey) *firestore.DocumentRef {
	path := dalgo.GetRecordKeyPath(key)
	return dtb.client.Doc(path)
}

func (dtb database) Upsert(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}
