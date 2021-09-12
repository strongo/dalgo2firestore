package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo"
)

type deleter struct {
	doc    func(key dalgo.RecordKey) *firestore.DocumentRef
	delete func(ctx context.Context, docRef *firestore.DocumentRef) (_ *firestore.WriteResult, err error)
}

var _ dalgo.Deleter = (*deleter)(nil)

func newDeleter(dtb database) deleter {
	return deleter{
		doc:    dtb.doc,
		delete: delete,
	}
}

func (d deleter) Delete(ctx context.Context, key dalgo.RecordKey) error {
	docRef := d.doc(key)
	_, err := d.delete(ctx, docRef)
	return err
}

func (dtb database) DeleteMulti(ctx context.Context, keys []dalgo.RecordKey) error {
	panic("implement me")
}
