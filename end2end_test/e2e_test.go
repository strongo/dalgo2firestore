package end2end

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo-firestore"
	"log"
	"os"
	"sync"
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
	db := dalgo_firestore.NewDatabase(client)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		t.Run("single", func(t *testing.T) {
			testSingleOperations(ctx, t, db)
		})
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		t.Run("multi", func(t *testing.T) {
			testMultiOperations(ctx, t, db)
		})
		wg.Done()
	}()
	wg.Wait()
}
