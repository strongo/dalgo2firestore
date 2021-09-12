package dalgo_firestore

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) Update(ctx context.Context, record dalgo.Record) error {
	panic("implement me")
}

func (dtb database) UpdateMulti(c context.Context, records []dalgo.Record) error {
	panic("implement me")
}
