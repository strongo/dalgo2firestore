package dalgo2firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/strongo/dalgo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type getter struct {
	doc    func(key *dalgo.Key) *firestore.DocumentRef
	dataTo func(ds *firestore.DocumentSnapshot, p interface{}) error
	get    func(ctx context.Context, docRef *firestore.DocumentRef) (_ *firestore.DocumentSnapshot, err error)
	getAll func(ctx context.Context, docRefs []*firestore.DocumentRef) (_ []*firestore.DocumentSnapshot, err error)
}

func newGetter(dtb database) getter {
	return getter{
		doc:    dtb.doc,
		get:    get,
		getAll: dtb.client.GetAll,
		dataTo: func(ds *firestore.DocumentSnapshot, p interface{}) error {
			return ds.DataTo(p)
		},
	}
}

func (g getter) Get(ctx context.Context, record dalgo.Record) error {
	key := record.Key()
	docRef := g.doc(key)
	docSnapshot, err := g.get(ctx, docRef)
	return docSnapshotToRecord(err, docSnapshot, record, g.dataTo)
}

func docSnapshotToRecord(
	err error,
	docSnapshot *firestore.DocumentSnapshot,
	record dalgo.Record,
	dataTo func(ds *firestore.DocumentSnapshot, p interface{}) error,
) error {
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return dalgo.NewErrNotFoundByKey(record.Key(), err)
		}
		return err
	}
	recData := record.Data()
	if err = dataTo(docSnapshot, recData); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to marshal record data into a struct of type %T", recData))
	}
	return nil
}

func (g getter) GetMulti(ctx context.Context, records []dalgo.Record) error {
	docRefs := make([]*firestore.DocumentRef, len(records), len(records))
	for i, rec := range records {
		docRefs[i] = g.doc(rec.Key())
	}
	docSnapshots, err := g.getAll(ctx, docRefs)
	if err != nil {
		return err
	}
	for i, rec := range records {
		data := rec.Data()
		if err := docSnapshots[i].DataTo(data); err != nil {
			return err
		}
	}
	return nil
}
