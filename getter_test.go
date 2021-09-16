package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo"
	"testing"
)

type getterMock struct {
	getCalled int
	getter    getter
}

func newGetterMock() *getterMock {
	var gm getterMock
	gm.getter = getter{
		doc: func(key *dalgo.Key) *firestore.DocumentRef {
			return nil
		},
		get: func(ctx context.Context, docRef *firestore.DocumentRef) (_ *firestore.DocumentSnapshot, err error) {
			gm.getCalled++
			return nil, err
		},
		dataTo: func(ds *firestore.DocumentSnapshot, p interface{}) error {
			return nil
		},
	}
	return &gm
}

func TestGetter_Get(t *testing.T) {
	gm := newGetterMock()
	ctx := context.Background()
	key := dalgo.NewKeyWithStrID("TestKind", "TestID")
	record := dalgo.NewRecord(key, nil)
	err := gm.getter.Get(ctx, record)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
