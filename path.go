package dalgo2firestore

import "github.com/strongo/dalgo"

func PathFromKey(key *dalgo.Key) string {
	return dalgo.GetRecordKeyPath(key)
}
