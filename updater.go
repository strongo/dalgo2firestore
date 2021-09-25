package dalgo2firestore

import (
	"context"
	"github.com/strongo/dalgo"
)

type updater struct {
	dtb *database
}

func (u updater) Update(
	_ context.Context,
	_ *dalgo.Key,
	_ []dalgo.Update,
	_ ...dalgo.Precondition,
) error {
	panic("not supported")
}

func (u updater) UpdateMulti(
	_ context.Context,
	_ []*dalgo.Key,
	_ []dalgo.Update,
	_ ...dalgo.Precondition,
) error {
	panic("not supported")
}
