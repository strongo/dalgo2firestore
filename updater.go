package dalgo_firestore

import (
	"context"
	"github.com/strongo/dalgo"
)

func (dtb database) Update(
	ctx context.Context,
	key *dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	panic("implement me")
}

func (dtb database) UpdateMulti(c context.Context, records []dalgo.Record) error {
	panic("implement me")
}
