package dalgo_firestore

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) RunInTransaction(ctx context.Context, f func(ctx context.Context, tx dalgo.Transaction) error, options ...dalgo.TransactionOption) error {
	panic("implement me")
}
