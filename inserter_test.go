package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo"
	"testing"
)

type inserterMock struct {
	createCalled int
	inserter     inserter
}

func newInserterMock() *inserterMock {
	im := inserterMock{}
	im.inserter = inserter{
		doc: func(key dalgo.RecordKey) *firestore.DocumentRef {
			return nil
		},
		create: func(ctx context.Context, docRef *firestore.DocumentRef, data interface{}) (_ *firestore.WriteResult, err error) {
			im.createCalled++
			return nil, nil
		},
	}
	return &im
}
func TestInserter_Insert(t *testing.T) {
	inserterMock := newInserterMock()
	ctx := context.Background()
	key := dalgo.NewRecordKey(dalgo.RecordRef{Kind: "TestKind", ID: "test-id"})
	record := dalgo.NewRecord(key, nil)
	err := inserterMock.inserter.Insert(ctx, record, dalgo.NewInsertOptions())
	if err != nil {
		t.Errorf("expected to be successful, got error: %v", err)
	}
	if inserterMock.createCalled != 1 {
		t.Errorf("expected a single call to create(), got %v", inserterMock.createCalled)
	}
}
