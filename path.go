package dalgo2firestore

import "github.com/strongo/dalgo/dal"

// PathFromKey generates full path of a key
func PathFromKey(key *dal.Key) string {
	return key.String()
}
