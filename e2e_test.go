package dalgo2firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	end2end "github.com/strongo/dalgo-end2end-tests"
	"log"
	"os"
	"testing"
	"time"
)

func TestEndToEnd(t *testing.T) {

	setEnv := func(key, value string) {
		if err := os.Setenv(key, value); err != nil {
			t.Fatal(err)
		}
	}
	setEnv("GCLOUD_PROJECT", "dalgo")
	setEnv("FIREBASE_AUTH_EMULATOR_HOST", "localhost:9099")
	setEnv("FIRESTORE_EMULATOR_HOST", "localhost:8080")

	firestoreProjectID := os.Getenv("DALGO_E2E_PROJECT_ID")

	if firestoreProjectID == "" {
		firestoreProjectID = "dalgo"
		//t.Fatalf("Environment variable DALGO_E2E_PROJECT_ID is not set")
	}
	log.Println("Firestore Project ID:", firestoreProjectID)
	//log.Println("ENV: GOOGLE_APPLICATION_CREDENTIALS:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	timeout, _ := time.ParseDuration("5s")
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(timeout))

	var client *firestore.Client
	var err error
	client, err = firestore.NewClient(ctx, firestoreProjectID)
	if err != nil {
		t.Fatalf("failed to create Firestore client: %v", err)
	}
	db := NewDatabase(client)

	end2end.TestDalgoDB(t, db)
}
