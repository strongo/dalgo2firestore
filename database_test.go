package dalgo_firestore

import (
	"cloud.google.com/go/firestore"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	var dtb = NewDatabase(&firestore.Client{})
	if dtb == nil {
		t.Error("NewDatabase returned nil")
	}
}
