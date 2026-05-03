package datastore

import (
	"context"

	"github.com/gear6io/pragmata/pkg/types/streamstoretypes"
)

// DataStore is an interface that defines the methods for interacting with the data store.
type DataStore interface {
	// Store writes a batch of events into the data store.
	Store(ctx context.Context, batch []streamstoretypes.Event) error
}
