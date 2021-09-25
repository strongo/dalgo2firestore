package dalgo2firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	end2end "github.com/strongo/dalgo-end2end-tests"
	"log"
	"os"
	"testing"
)

func TestEndToEnd(t *testing.T) {

	firestoreProjectID := os.Getenv("DALGO_E2E_PROJECT_ID")

	if firestoreProjectID == "" {
		firestoreProjectID = "DALGO_E2E"
		//t.Fatalf("Environment variable DALGO_E2E_PROJECT_ID is not set")
	}
	log.Println("Firestore Project ID:", firestoreProjectID)
	//log.Println("ENV: GOOGLE_APPLICATION_CREDENTIALS:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	ctx := context.Background()

	var client *firestore.Client
	var err error
	client, err = firestore.NewClient(ctx, firestoreProjectID)
	if err != nil {
		t.Fatalf("failed to create Firestore client: %v", err)
	}
	db := NewDatabase(client)

	end2end.TestDalgoDB(t, db)
}
