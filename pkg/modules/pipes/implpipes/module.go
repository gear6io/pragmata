package implpipes

import (
	"context"
	"fmt"
	"time"

	"github.com/gear6io/pragmata/pkg/datastore"
	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/orchestration"
	"github.com/gear6io/pragmata/pkg/pipevisitor"
	"github.com/gear6io/pragmata/pkg/prqlvisitor"
	"github.com/gear6io/pragmata/pkg/sqlstore"
	"github.com/gear6io/pragmata/pkg/types"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
	"github.com/gear6io/pragmata/pkg/valuer"
)

type module struct {
	datastore datastore.DataStore
	store     sqlstore.SQLStore
	orchest   orchestration.Orchestrator
	sched     pipes.Scheduler
}

// NewModule constructs a Module with all required dependencies.
func NewModule(
	datastore datastore.DataStore,
	store sqlstore.SQLStore,
	orchest orchestration.Orchestrator,
	sched pipes.Scheduler,
) pipes.Module {
	return &module{
		datastore: datastore,
		store:     store,
		orchest:   orchest,
		sched:     sched,
	}
}

func (m *module) CreatePipe(ctx context.Context, postable *pipetypes.PostablePipe) (*pipetypes.Pipe, error) {
	pipe, err := pipevisitor.Visit("", postable.Content, pipevisitor.PipeVisitorOpts{
		FetchSources: func(srcs ...string) ([]sourcetypes.Source, error) {
			return m.datastore.ListSources(ctx, srcs)
		},
		SourceValidator: prqlvisitor.NewSourceValidator,
	})
	if err != nil {
		return nil, err
	}

	// create storable flavor
	storable := pipetypes.StorablePipe{
		Identifiable: types.Identifiable{
			ID: valuer.GenerateUUID(),
		},
		Pipe: *pipe,
		TimeAuditable: types.TimeAuditable{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := m.store.CreatePipe(ctx, &storable); err != nil {
		return nil, fmt.Errorf("store: %w", err)
	}
	switch pipe.Type {
	case pipetypes.PipeTypeMaterialized:
		_, err := m.orchest.StartMaterializedPipe(ctx, orchestration.MaterializedPipeParams{
			Pipe: pipe,
		})
		if err != nil {
			return nil, fmt.Errorf("start materialized pipe: %w", err)
		}
	case pipetypes.PipeTypeCopy:
		if m.sched != nil && pipe.CopySchedule != "" {
			if err := m.sched.Register(pipe); err != nil {
				return nil, fmt.Errorf("register schedule: %w", err)
			}
		}
	}
	return pipe, nil
}

func (m *module) GetPipe(ctx context.Context, name string) (*pipetypes.Pipe, error) {
	return m.store.GetPipe(ctx, name)
}

func (m *module) ListPipes(ctx context.Context) ([]*pipetypes.Pipe, error) {
	return m.store.ListPipes(ctx)
}

func (m *module) UpdatePipe(ctx context.Context, pipe *pipetypes.StorablePipe) (*pipetypes.Pipe, error) {
	if err := m.store.UpdatePipe(ctx, pipe); err != nil {
		return nil, fmt.Errorf("store: %w", err)
	}
	if pipe.Type == pipetypes.PipeTypeMaterialized {
		_, err := m.orchest.StartMaterializedPipe(ctx, orchestration.MaterializedPipeParams{Pipe: &pipe.Pipe})
		if err != nil {
			return nil, fmt.Errorf("re-sync materialized pipe: %w", err)
		}
	}
	if pipe.Type == pipetypes.PipeTypeCopy && m.sched != nil {
		if pipe.CopySchedule != "" {
			_ = m.sched.Register(&pipe.Pipe)
		} else {
			m.sched.Unregister(pipe.Name)
		}
	}
	return &pipe.Pipe, nil
}

func (m *module) DeletePipe(ctx context.Context, name string) error {
	if m.sched != nil {
		m.sched.Unregister(name)
	}
	return m.store.DeletePipe(ctx, name)
}

// func (m *module) ExecutePipe(ctx context.Context, name string, params map[string]string) (*pipetypes.ExecuteResult, error) {
// 	pipe, err := m.store.GetPipe(ctx, name)
// 	if err != nil {
// 		return nil, fmt.Errorf("load pipe: %w", err)
// 	}
// 	sql, err := executor.BuildCTE(pipe.Nodes, params)
// 	if err != nil {
// 		return nil, fmt.Errorf("build query: %w", err)
// 	}
// 	result, err := m.exec.Query(ctx, sql)
// 	if err != nil {
// 		return nil, fmt.Errorf("execute: %w", err)
// 	}
// 	return result, nil
// }
