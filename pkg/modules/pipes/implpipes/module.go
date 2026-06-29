package implpipes

import (
	"context"
	"fmt"
	"time"

	"github.com/gear6io/pragmata/pkg/datastore"
	"github.com/gear6io/pragmata/pkg/executor"
	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/pipevisitor"
	"github.com/gear6io/pragmata/pkg/prqlvisitor"
	"github.com/gear6io/pragmata/pkg/querier"
	"github.com/gear6io/pragmata/pkg/scheduler"
	"github.com/gear6io/pragmata/pkg/sqlstore"
	"github.com/gear6io/pragmata/pkg/types"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
	"github.com/gear6io/pragmata/pkg/valuer"
)

type module struct {
	datastore datastore.DataStore
	store     sqlstore.SQLStore
	exec      executor.Executor
	sched     scheduler.Scheduler
	querier   *querier.Querier
}

// NewModule constructs a Module with all required dependencies.
func NewModule(
	datastore datastore.DataStore,
	store sqlstore.SQLStore,
	exec executor.Executor,
	sched scheduler.Scheduler,
	q *querier.Querier,
) pipes.Module {
	return &module{
		datastore: datastore,
		store:     store,
		exec:      exec,
		sched:     sched,
		querier:   q,
	}
}

func (m *module) CreatePipe(ctx context.Context, postable *pipetypes.PostablePipe) (*pipetypes.GettablePipe, error) {
	exec, err := pipevisitor.Visit(postable.Content, pipevisitor.PipeVisitorOpts{
		FetchSources: func(srcs ...string) ([]sourcetypes.Source, error) {
			return m.datastore.ListSources(ctx, srcs)
		},
		SourceValidator: prqlvisitor.NewSourceValidator,
	})
	if err != nil {
		return nil, err
	}

	storable := pipetypes.StorablePipe{
		Identifiable: types.Identifiable{
			ID: valuer.GenerateUUID(),
		},
		Pipe: exec.Pipe,
		TimeAuditable: types.TimeAuditable{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := m.store.CreatePipe(ctx, &storable); err != nil {
		return nil, fmt.Errorf("store: %w", err)
	}
	switch exec.Type {
	case pipetypes.PipeTypeMaterialized:
		_, err := m.exec.StartMaterializedPipe(ctx, executor.MaterializedPipeParams{
			Pipe: exec,
		})
		if err != nil {
			return nil, fmt.Errorf("start materialized pipe: %w", err)
		}
		if m.sched != nil && exec.CopySchedule != "" {
			if err := m.sched.Register(exec); err != nil {
				return nil, fmt.Errorf("register schedule: %w", err)
			}
		}
	case pipetypes.PipeTypeCopy:
		if m.sched != nil && exec.CopySchedule != "" {
			if err := m.sched.Register(exec); err != nil {
				return nil, fmt.Errorf("register schedule: %w", err)
			}
		}
	}
	return &storable, nil
}

func (m *module) GetPipe(ctx context.Context, name string) (*pipetypes.GettablePipe, error) {
	return m.store.GetPipe(ctx, name)
}

func (m *module) ListPipes(ctx context.Context) ([]*pipetypes.GettablePipe, error) {
	return m.store.ListPipes(ctx)
}

func (m *module) UpdatePipe(ctx context.Context, exec *pipetypes.ExecutablePipe) (*pipetypes.GettablePipe, error) {
	storable := &pipetypes.StorablePipe{
		Pipe: exec.Pipe,
		TimeAuditable: types.TimeAuditable{
			UpdatedAt: time.Now(),
		},
	}
	if err := m.store.UpdatePipe(ctx, storable); err != nil {
		return nil, fmt.Errorf("store: %w", err)
	}
	if exec.Type == pipetypes.PipeTypeMaterialized {
		_, err := m.exec.StartMaterializedPipe(ctx, executor.MaterializedPipeParams{Pipe: exec})
		if err != nil {
			return nil, fmt.Errorf("re-sync materialized pipe: %w", err)
		}
	}
	if m.sched != nil {
		if exec.CopySchedule != "" {
			_ = m.sched.Register(exec)
		} else {
			m.sched.Unregister(exec.Name)
		}
	}
	return storable, nil
}

func (m *module) DeletePipe(ctx context.Context, name string) error {
	if m.sched != nil {
		m.sched.Unregister(name)
	}
	return m.store.DeletePipe(ctx, name)
}

func (m *module) ExecutePipe(ctx context.Context, name string, params map[string]string) (*pipetypes.ExecuteResult, error) {
	gettable, err := m.store.GetPipe(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("get pipe: %w", err)
	}
	exec, err := pipevisitor.Visit(gettable.Content, pipevisitor.PipeVisitorOpts{})
	if err != nil {
		return nil, fmt.Errorf("parse pipe: %w", err)
	}
	return m.querier.Execute(ctx, exec, params)
}
