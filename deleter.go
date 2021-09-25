package dalgo2firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo"
)

type deleter struct {
	doc    func(key *dalgo.Key) *firestore.DocumentRef
	delete func(ctx context.Context, docRef *firestore.DocumentRef) (_ *firestore.WriteResult, err error)
	batch  func() *firestore.WriteBatch
}

var _ dalgo.Deleter = (*deleter)(nil)

func newDeleter(dtb database) deleter {
	return deleter{
		doc:    dtb.doc,
		delete: delete,
		batch: func() *firestore.WriteBatch {
			return dtb.client.Batch()
		},
	}
}

func (d deleter) Delete(ctx context.Context, key *dalgo.Key) error {
	docRef := d.doc(key)
	_, err := d.delete(ctx, docRef)
	return err
}

func (d deleter) DeleteMulti(ctx context.Context, keys []*dalgo.Key) error {
	batch := d.batch()
	for _, key := range keys {
		docRef := d.doc(key)
		batch.Delete(docRef)
	}
	_, err := batch.Commit(ctx)
	return err
}
