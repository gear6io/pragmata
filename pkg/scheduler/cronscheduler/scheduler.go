// Package cronscheduler manages cron-driven execution of pipes using robfig/cron.
package cronscheduler

import (
	"context"
	"log"
	"sync"

	"github.com/robfig/cron/v3"

	"github.com/gear6io/pragmata/pkg/executor"
	"github.com/gear6io/pragmata/pkg/pipevisitor"
	"github.com/gear6io/pragmata/pkg/sqlstore"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// Scheduler registers and fires pipe cron jobs.
// Implements factory.Service (Start/Stop) and scheduler.Scheduler (Register/Unregister).
type Scheduler struct {
	cron    *cron.Cron
	exec    executor.Executor
	store   sqlstore.SQLStore
	entries map[string]cron.EntryID
	mu      sync.Mutex
}

// New creates a Scheduler backed by the given executor and store.
func New(exec executor.Executor, store sqlstore.SQLStore) *Scheduler {
	return &Scheduler{
		cron:    cron.New(),
		exec:    exec,
		store:   store,
		entries: make(map[string]cron.EntryID),
	}
}

// Start loads all pipes with a schedule from the store, registers their cron entries,
// and begins the scheduler. Implements factory.Service.
func (s *Scheduler) Start(ctx context.Context) error {
	pipes, err := s.store.ListPipes(ctx)
	if err != nil {
		return err
	}
	for _, p := range pipes {
		if p.CopySchedule == "" {
			continue
		}
		exec, err := pipevisitor.Visit(p.Content, pipevisitor.PipeVisitorOpts{})
		if err != nil {
			return err
		}
		if err := s.Register(exec); err != nil {
			return err
		}
	}
	s.cron.Start()
	return nil
}

// Stop gracefully halts the scheduler, waiting for any running jobs to finish.
// Implements factory.Service.
func (s *Scheduler) Stop(_ context.Context) error {
	s.cron.Stop()
	return nil
}

// Register adds a pipe to the cron scheduler.
// If the pipe is already registered, it is replaced.
func (s *Scheduler) Register(pipe *pipetypes.ExecutablePipe) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if entryID, ok := s.entries[pipe.Name]; ok {
		s.cron.Remove(entryID)
	}

	pipeName := pipe.Name
	exec := s.exec
	entryID, err := s.cron.AddFunc(pipe.CopySchedule, func() {
		if err := exec.RunPipe(context.Background(), pipeName); err != nil {
			log.Printf("pipe %q run failed: %v", pipeName, err)
		}
	})
	if err != nil {
		return err
	}
	s.entries[pipe.Name] = entryID
	return nil
}

// Unregister removes a pipe from the cron scheduler.
func (s *Scheduler) Unregister(pipeID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if entryID, ok := s.entries[pipeID]; ok {
		s.cron.Remove(entryID)
		delete(s.entries, pipeID)
	}
}
