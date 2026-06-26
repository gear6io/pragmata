package implsources

import (
	"context"

	"github.com/gear6io/pragmata/pkg/datastore"
	"github.com/gear6io/pragmata/pkg/modules/sources"
	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
)

type module struct {
	store datastore.DataStore
}

// NewModule returns a Module backed by the given DataStore.
func NewModule(store datastore.DataStore) sources.Module {
	return &module{store: store}
}

func (m *module) CreateSource(ctx context.Context, src *sourcetypes.Source) (*sourcetypes.Source, error) {
	if err := m.store.CreateSource(ctx, src); err != nil {
		return nil, err
	}
	return src, nil
}

func (m *module) ListSources(ctx context.Context) ([]sourcetypes.Source, error) {
	return m.store.ListSources(ctx, nil)
}

func (m *module) GetSource(ctx context.Context, name string) (*sourcetypes.Source, error) {
	return m.store.GetSource(ctx, name)
}
