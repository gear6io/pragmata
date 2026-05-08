package implpipes

import (
	"context"
	"fmt"

	"github.com/gear6io/pragmata/internal/executor"
	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/orchestration"
	"github.com/gear6io/pragmata/pkg/sqlstore"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

type module struct {
	store    sqlstore.SQLStore
	orchest  orchestration.Orchestrator
	exec     executor.Executor
	sched    pipes.Scheduler
}

// NewModule constructs a Module with all required dependencies.
func NewModule(
	store sqlstore.SQLStore,
	orchest orchestration.Orchestrator,
	exec executor.Executor,
	sched pipes.Scheduler,
) pipes.Module {
	return &module{
		store:   store,
		orchest: orchest,
		exec:    exec,
		sched:   sched,
	}
}

func (m *module) CreatePipe(ctx context.Context, pipe *pipetypes.Pipe) (*pipetypes.Pipe, error) {
	if err := m.store.CreatePipe(ctx, pipe); err != nil {
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

func (m *module) UpdatePipe(ctx context.Context, name string, pipe *pipetypes.Pipe) (*pipetypes.Pipe, error) {
	pipe.Name = name
	if err := m.store.UpdatePipe(ctx, pipe); err != nil {
		return nil, fmt.Errorf("store: %w", err)
	}
	// Re-sync orchestration for MATERIALIZED pipes on update.
	if pipe.Type == pipetypes.PipeTypeMaterialized {
		_, err := m.orchest.StartMaterializedPipe(ctx, orchestration.MaterializedPipeParams{Pipe: pipe})
		if err != nil {
			return nil, fmt.Errorf("re-sync materialized pipe: %w", err)
		}
	}
	if pipe.Type == pipetypes.PipeTypeCopy && m.sched != nil {
		if pipe.CopySchedule != "" {
			_ = m.sched.Register(pipe)
		} else {
			m.sched.Unregister(name)
		}
	}
	return pipe, nil
}

func (m *module) DeletePipe(ctx context.Context, name string) error {
	if m.sched != nil {
		m.sched.Unregister(name)
	}
	return m.store.DeletePipe(ctx, name)
}

func (m *module) ExecutePipe(ctx context.Context, name string, params map[string]string) (*pipetypes.ExecuteResult, error) {
	pipe, err := m.store.GetPipe(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("load pipe: %w", err)
	}
	sql, err := executor.BuildCTE(pipe.Nodes, params)
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}
	result, err := m.exec.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute: %w", err)
	}
	return result, nil
}
