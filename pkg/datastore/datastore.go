package datastore

import (
	"context"

	"github.com/gear6io/pragmata/pkg/config"
	"github.com/gear6io/pragmata/pkg/datastore/clickhouse"
	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
)

// DataStore is the interface for interacting with the ClickHouse data store.
type DataStore interface {
	CreateSource(ctx context.Context, src *sourcetypes.Source) error
	ListSources(ctx context.Context, match []string) ([]sourcetypes.Source, error)
	GetSource(ctx context.Context, name string) (*sourcetypes.Source, error)
}

// New returns a DataStore backed by ClickHouse using the given config.
func New(cfg config.ClickHouseConfig) (DataStore, error) {
	return clickhouse.New(cfg.URL)
}
