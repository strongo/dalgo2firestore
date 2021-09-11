package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/db"
)

type deleter struct {
	delete func(ctx context.Context, key db.RecordKey) (*firestore.WriteResult, error)
	doc    func(path string) *firestore.DocumentRef
}

var _ db.Deleter = (*deleter)(nil)

func newDeleter(dtb *database) deleter {
	return deleter{
		delete: func(ctx context.Context, key db.RecordKey) (*firestore.WriteResult, error) {
			return dtb.delete(ctx, key)
		},
	}
}

func (d deleter) Delete(ctx context.Context, key db.RecordKey) error {
	_, err := d.delete(ctx, key)
	return err
}

func (dtb database) delete(ctx context.Context, key db.RecordKey) (*firestore.WriteResult, error) {
	path := PathFromKey(key)
	docRef := dtb.client.Doc(path)
	return docRef.Delete(ctx)
}
