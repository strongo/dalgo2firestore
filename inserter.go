package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo"
	"log"
)

type inserter struct {
	doc    func(key *dalgo.Key) *firestore.DocumentRef
	create func(ctx context.Context, docRef *firestore.DocumentRef, data interface{}) (_ *firestore.WriteResult, err error)
}

var _ dalgo.Inserter = (*inserter)(nil)

func newInserter(dtb database) inserter {
	return inserter{
		doc:    dtb.doc,
		create: create,
	}
}

func (i inserter) Insert(ctx context.Context, record dalgo.Record, opts ...dalgo.InsertOption) error {
	options := dalgo.NewInsertOptions(opts...)
	generateID := options.IDGenerator()
	if generateID != nil {
		if err := generateID(ctx, record); err != nil {
			return err
		}
	}
	_, err := i.insert(ctx, record)
	return err
}

func (i inserter) insert(ctx context.Context, record dalgo.Record) (*firestore.WriteResult, error) {
	key := record.Key()
	docRef := i.doc(key)
	if docRef != nil {
		log.Println("inserting document:", docRef.Path)
	}
	data := record.Data()
	return i.create(ctx, docRef, data)
}
