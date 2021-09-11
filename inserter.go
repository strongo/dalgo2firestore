package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/db"
)

type inserter struct {
	create func(ctx context.Context, record db.Record) (*firestore.WriteResult, error)
}

var _ db.Inserter = (*inserter)(nil)

func newInserter(dtb *database) inserter {
	return inserter{
		create: func(ctx context.Context, record db.Record) (*firestore.WriteResult, error) {
			return dtb.create(ctx, record)
		},
	}
}

func (i inserter) Insert(ctx context.Context, record db.Record, options db.InsertOptions) error {
	generateID := options.IDGenerator()
	if generateID != nil {
		if err := generateID(ctx, record); err != nil {
			return err
		}
	}
	_, err := i.create(ctx, record)
	return err
}

func (dtb database) create(ctx context.Context, record db.Record) (*firestore.WriteResult, error) {
	key := record.Key()
	path := PathFromKey(key)
	docRef := dtb.client.Doc(path)
	data := record.Data()
	return docRef.Create(ctx, data)
}
