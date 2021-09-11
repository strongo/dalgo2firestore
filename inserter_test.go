package db_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/db"
	"testing"
)

func TestInserter_Insert(t *testing.T) {
	createCalled := 0
	v := inserter{create: func(ctx context.Context, record db.Record) (*firestore.WriteResult, error) {
		createCalled++
		return nil, nil
	}}
	ctx := context.Background()
	key := db.NewRecordKey(db.RecordRef{Kind: "TestKind", ID: "test-id"})
	record := db.NewRecord(key, nil)
	err := v.Insert(ctx, record, db.NewInsertOptions())
	if err != nil {
		t.Errorf("expected to be successful, got error: %v", err)
	}
}
