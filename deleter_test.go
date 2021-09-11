package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/db"
	"testing"
)

func TestDeleter_Delete(t *testing.T) {
	deleteCalled := 0
	v := deleter{delete: func(ctx context.Context, key db.RecordKey) (*firestore.WriteResult, error) {
		deleteCalled++
		return nil, nil
	}}
	ctx := context.Background()
	key := db.NewRecordKey(db.RecordRef{Kind: "TestKind", ID: "test-id"})
	err := v.Delete(ctx, key)
	if err != nil {
		t.Errorf("expected to be successful, got error: %v", err)
	}
}
