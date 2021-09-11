package dalgo_firestore

import "github.com/strongo/db"

func PathFromKey(key db.RecordKey) string {
	return db.GetRecordKeyPath(key)
}
