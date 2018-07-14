// Package gcd provides helpful functions to work with Google Cloud DataStore.
package gcd

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// DropKind drops entire kind by deleting all entities from it.
func DropKind(ctx context.Context, kind string) error {
	keys, err := GetKeys(ctx, kind)
	if err != nil {
		return fmt.Errorf("failed to get keys for kind, error: %v", err)
	}

	err = datastore.DeleteMulti(ctx, keys)
	if err != nil {
		return fmt.Errorf("failed to delete all from kind, error: %v", err)
	}

	return nil
}

// GetKeys gets slice with all kind's keys.
func GetKeys(ctx context.Context, kind string) ([]*datastore.Key, error) {
	q := datastore.NewQuery(kind).KeysOnly()
	t := q.Run(ctx)

	keys := make([]*datastore.Key, 0)
	for {
		var d interface{}

		key, err := t.Next(&d)
		if err == datastore.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to fetch next key, error: %v", err)
		}

		keys = append(keys, key)
	}

	return keys, nil
}
