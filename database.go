package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/db"
)

type database struct {
	inserter
	deleter
	client *firestore.Client
}

var _ db.Database = (*database)(nil)

func NewDatabase() db.Database {
	var dtb database
	dtb.inserter = newInserter(&dtb)
	dtb.deleter = newDeleter(&dtb)
	return dtb
}

func (dtb database) RunInTransaction(ctx context.Context, f func(ctx context.Context, tx db.Transaction) error, options ...db.TransactionOption) error {
	panic("implement me")
}

func (dtb database) Upsert(ctx context.Context, record db.Record) error {
	panic("implement me")
}

func (dtb database) Get(ctx context.Context, record db.Record) error {
	panic("implement me")
}

func (dtb database) Set(ctx context.Context, record db.Record) error {
	panic("implement me")
}

func (dtb database) Update(ctx context.Context, record db.Record) error {
	panic("implement me")
}

func (dtb database) GetMulti(ctx context.Context, records []db.Record) error {
	panic("implement me")
}

func (dtb database) SetMulti(ctx context.Context, records []db.Record) error {
	panic("implement me")
}

func (dtb database) UpdateMulti(c context.Context, records []db.Record) error {
	panic("implement me")
}

func (dtb database) DeleteMulti(ctx context.Context, keys []db.RecordKey) error {
	panic("implement me")
}
