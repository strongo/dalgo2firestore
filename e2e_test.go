package dalgo2firestore

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"context"
	"github.com/strongo/dalgo/end2end"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestEndToEnd(t *testing.T) {
	log.Println("TestEndToEnd() started...")
	cmd, cmdStdout, cmdStdErr := startFirebaseEmulators(t)
	defer terminateFirebaseEmulators(cmd)
	emulatorExited := false
	go handleCommandStderr(t, cmdStdErr, &emulatorExited)
	select {
	case <-handleEmulatorClosing(t, cmd):
		emulatorExited = true
		t.Log("Firebase emulator exited")
	case <-waitForEmulatorReadiness(t, cmdStdout, &emulatorExited):
		testEndToEnd(t)
	}
}

func handleCommandStderr(t *testing.T, stderr *bytes.Buffer, emulatorExited *bool) {
	for {
		if *emulatorExited {
			return
		}
		line, err := stderr.ReadString('\n')
		if err == io.EOF {
			time.Sleep(9 * time.Millisecond)
			continue
		}
		if err != nil {
			t.Errorf("Failed to read from Firebase emulator STDERR: %v", err)
			return
		}
		t.Errorf("ERROR in Firebase emulator: " + line)
	}
}

func terminateFirebaseEmulators(cmd *exec.Cmd) {
	if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
		log.Printf("Failed to terminate Firebase emulator: %v", err)
		return
	}
}

func startFirebaseEmulators(t *testing.T) (cmd *exec.Cmd, stdout, stderr *bytes.Buffer) {
	cmd = exec.Command("firebase",
		"emulators:start",
		"-c", "./firebase/firebase.json",
		"--only", "firestore",
		"--project", "dalgo",
	)

	stdout = new(bytes.Buffer)
	stderr = new(bytes.Buffer)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	t.Log("Starting Firebase emulator...")
	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start Firebase emulator: %v", err)
	}
	return
}

func waitForEmulatorReadiness(t *testing.T, cmdOutput *bytes.Buffer, emulatorExited *bool) (emulatorsReady chan bool) {
	emulatorsReady = make(chan bool)
	//time.Sleep(3 * time.Second)
	go func() {
		t.Log("Awaiting for Firebase emulator to be ready...")
		for i := 1; true; i++ {
			if *emulatorExited {
				return
			}
			line, err := cmdOutput.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					time.Sleep(10 * time.Millisecond)
					continue
				}
				t.Errorf("Failed to read: %v", err)
				return
			}
			if strings.Contains(line, "All emulators ready!") {
				//t.Log("Firebase emulators are ready.")
				emulatorsReady <- true
				close(emulatorsReady)
			}
		}
	}()
	return
}

func handleEmulatorClosing(t *testing.T, cmd *exec.Cmd) (emulatorErrors chan error) {
	emulatorErrors = make(chan error)
	go func() {
		err := cmd.Wait()
		if err != nil {
			if err.Error() == "signal: killed" {
				t.Log("Firebase emulator killed.")
			} else {
				t.Errorf("Firebase emulator failed: %v", err)
				emulatorErrors <- err
			}
		} else {
			t.Log("Firebase emulator completed.")
		}
		close(emulatorErrors)
	}()
	return
}

func testEndToEnd(t *testing.T) {
	if err := os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080"); err != nil {
		t.Fatalf("Failed to set env variable FIRESTORE_EMULATOR_HOST: %v", err)
	}
	firestoreProjectID := os.Getenv("DALGO_E2E_PROJECT_ID")

	if firestoreProjectID == "" {
		firestoreProjectID = "dalgo"
		//t.Fatalf("Environment variable DALGO_E2E_PROJECT_ID is not set")
	}
	log.Println("Firestore Project ID:", firestoreProjectID)
	//log.Println("ENV: GOOGLE_APPLICATION_CREDENTIALS:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	ctx := context.Background()

	//var client *firestore.Client
	client, err := firestore.NewClient(ctx, firestoreProjectID)
	if err != nil {
		t.Fatalf("failed to create Firestore client: %v", err)
	}
	db := NewDatabase(client)

	end2end.TestDalgoDB(t, db)
}
