package db_firestore

import (
	"testing"
)

func TestNewDatabase(t *testing.T) {
	var dtb = NewDatabase()
	if dtb == nil {
		t.Error("NewDatabase returned nil")
	}
}
