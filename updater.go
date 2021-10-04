package dalgo2firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/strongo/dalgo"
)

type updater struct {
	dtb *database
}

func newUpdater(dtb *database) updater {
	return updater{
		dtb: dtb,
	}
}

func (u updater) Update(
	ctx context.Context,
	key *dalgo.Key,
	update []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	return u.dtb.RunInTransaction(ctx, func(ctx context.Context, d dalgo.Transaction) error {
		return d.Update(ctx, key, update, preconditions...)
	})
}

func (u updater) UpdateMulti(
	ctx context.Context,
	keys []*dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	return u.dtb.RunInTransaction(ctx, func(ctx context.Context, tx dalgo.Transaction) error {
		return tx.UpdateMulti(ctx, keys, updates, preconditions...)
	})
}

func (t transaction) Update(
	_ context.Context,
	key *dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	dr := t.dtb.doc(key)
	fsUpdates := make([]firestore.Update, len(updates))
	for i, u := range updates {
		fsUpdates[i] = getFirestoreUpdate(u)
	}
	fsPreconditions := getUpdatePreconditions(preconditions)
	return t.tx.Update(dr, fsUpdates, fsPreconditions...)
}

func (t transaction) UpdateMulti(
	_ context.Context,
	keys []*dalgo.Key,
	updates []dalgo.Update,
	preconditions ...dalgo.Precondition,
) error {
	fsPreconditions := getUpdatePreconditions(preconditions)
	for _, key := range keys {
		dr := t.dtb.doc(key)
		fsUpdates := make([]firestore.Update, len(updates))
		for i, u := range updates {
			fsUpdates[i] = getFirestoreUpdate(u)
		}
		if err := t.tx.Update(dr, fsUpdates, fsPreconditions...); err != nil {
			keyPath := PathFromKey(key)
			return fmt.Errorf("failed to update record with key: %v: %w", keyPath, err)
		}
	}
	return nil
}

func getFirestoreUpdate(update dalgo.Update) firestore.Update {
	return firestore.Update{
		Path:      update.Field,
		FieldPath: (firestore.FieldPath)(update.FieldPath),
		Value:     update.Value,
	}
}

func getUpdatePreconditions(preconditions []dalgo.Precondition) (fsPreconditions []firestore.Precondition) {
	p := dalgo.GetPreconditions(preconditions...)
	if p.Exists() {
		fsPreconditions = []firestore.Precondition{firestore.Exists}
	}
	return fsPreconditions
}
