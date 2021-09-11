package db_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/db"
)

type inserter struct {
	client *firestore.Client
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
	docRef := i.client.Doc(PathFromKey(key))
	_, err := docRef.Create(ctx, record.Data())
	return err
}
