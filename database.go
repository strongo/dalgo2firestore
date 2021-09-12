package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo"
)

type database struct {
	inserter
	deleter
	client *firestore.Client
}

var _ dalgo.Database = (*database)(nil)

func NewDatabase() dalgo.Database {
	var dtb database
	dtb.inserter = newInserter(&dtb)
	dtb.deleter = newDeleter(&dtb)
	return dtb
}

func (dtb database) Upsert(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}
