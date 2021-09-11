package db_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/db"
)

type inserter struct {
	doc func(path string) *firestore.DocumentRef
}

var _ db.Inserter = (*inserter)(nil)

func (i inserter) Insert(ctx context.Context, record db.Record, options db.InsertOptions) error {
	generateID := options.IDGenerator()
	if generateID != nil {
		if err := generateID(ctx, record); err != nil {
			return err
		}
	}
	key := record.Key()
	path := PathFromKey(key)
	docRef := i.doc(path)
	data := record.Data()
	_, err := docRef.Create(ctx, data)
	return err
}
