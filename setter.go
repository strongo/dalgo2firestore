package dalgo2firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo/dal"
)

type setter struct {
	doc   func(key *dal.Key) *firestore.DocumentRef
	set   func(ctx context.Context, docRef *firestore.DocumentRef, data interface{}) (_ *firestore.WriteResult, err error)
	batch func() *firestore.WriteBatch
}

func newSetter(dtb database) setter {
	return setter{
		doc: dtb.doc,
		set: set,
		batch: func() *firestore.WriteBatch {
			return dtb.client.Batch()
		},
	}
}

func (s setter) Set(ctx context.Context, record dal.Record) error {
	key := record.Key()
	docRef := s.doc(key)
	data := record.Data()
	_, err := s.set(ctx, docRef, data)
	return err
}

func (s setter) SetMulti(ctx context.Context, records []dal.Record) error {
	batch := s.batch()
	for _, record := range records {
		key := record.Key()
		docRef := s.doc(key)
		data := record.Data()
		batch.Set(docRef, data)
	}
	_, err := batch.Commit(ctx)
	return err
}
