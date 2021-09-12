package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo"
)

type inserter struct {
	doc    func(key dalgo.RecordKey) *firestore.DocumentRef
	create func(ctx context.Context, record dalgo.Record) (*firestore.WriteResult, error)
}

var _ dalgo.Inserter = (*inserter)(nil)

func newInserter(dtb *database) inserter {
	return inserter{
		create: func(ctx context.Context, record dalgo.Record) (*firestore.WriteResult, error) {
			return dtb.create(ctx, record)
		},
	}
}

func (i inserter) Insert(ctx context.Context, record dalgo.Record, options dalgo.InsertOptions) error {
	generateID := options.IDGenerator()
	if generateID != nil {
		if err := generateID(ctx, record); err != nil {
			return err
		}
	}
	_, err := i.create(ctx, record)
	return err
}

func (dtb database) create(ctx context.Context, record dalgo.Record) (*firestore.WriteResult, error) {
	key := record.Key()
	path := PathFromKey(key)
	docRef := dtb.client.Doc(path)
	data := record.Data()
	return docRef.Create(ctx, data)
}
