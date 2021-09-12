package dalgo_firestore

import "github.com/strongo/dalgo"

func PathFromKey(key dalgo.RecordKey) string {
	return dalgo.GetRecordKeyPath(key)
}
