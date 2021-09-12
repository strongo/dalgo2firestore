package dalgo_firestore

import (
	"context"
	"github.com/strongo/dalgo"
)

type getter struct {
}

func newGetter() getter {
	return getter{}
}

func (dtb database) Get(ctx context.Context, record dalgo.Record) error {
	//client, err := firestore.NewClient(ctx, "")
	//client.D
	panic("implement me")
}

func (dtb database) GetMulti(ctx context.Context, records []dalgo.Record) error {
	panic("implement me")
}
