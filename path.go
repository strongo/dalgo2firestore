package dalgo2firestore

import "github.com/strongo/dalgo"

// PathFromKey generates full path of a key
func PathFromKey(key *dalgo.Key) string {
	return key.String()
}
