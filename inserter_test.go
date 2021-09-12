package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo"
	"testing"
)

func TestInserter_Insert(t *testing.T) {
	createCalled := 0
	v := inserter{create: func(ctx context.Context, record dalgo.Record) (*firestore.WriteResult, error) {
		createCalled++
		return nil, nil
	}}
	ctx := context.Background()
	key := dalgo.NewRecordKey(dalgo.RecordRef{Kind: "TestKind", ID: "test-id"})
	record := dalgo.NewRecord(key, nil)
	err := v.Insert(ctx, record, dalgo.NewInsertOptions())
	if err != nil {
		t.Errorf("expected to be successful, got error: %v", err)
	}
}
