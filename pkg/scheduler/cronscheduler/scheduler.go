// Package cronscheduler manages cron-driven execution of COPY pipes using robfig/cron.
package cronscheduler

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/robfig/cron/v3"

	"github.com/gear6io/pragmata/pkg/orchestration"
	"github.com/gear6io/pragmata/pkg/sqlstore"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// Scheduler registers and fires COPY pipe cron jobs.
// Implements factory.Service (Start/Stop) and scheduler.Scheduler (Register/Unregister).
type Scheduler struct {
	cron    *cron.Cron
	orchest orchestration.Orchestrator
	store   sqlstore.SQLStore
	entries map[string]cron.EntryID
	mu      sync.Mutex
}

// New creates a Scheduler backed by the given orchestrator and store.
func New(orchest orchestration.Orchestrator, store sqlstore.SQLStore) *Scheduler {
	return &Scheduler{
		cron:    cron.New(),
		orchest: orchest,
		store:   store,
		entries: make(map[string]cron.EntryID),
	}
}

// Start loads all COPY pipes from the store, registers their cron entries, and
// begins the scheduler. Implements factory.Service.
func (s *Scheduler) Start(ctx context.Context) error {
	pipes, err := s.store.ListPipes(ctx)
	if err != nil {
		return fmt.Errorf("load pipes for scheduler: %w", err)
	}
	for _, p := range pipes {
		if p.Type == pipetypes.PipeTypeCopy && p.CopySchedule != "" {
			if err := s.Register(p); err != nil {
				return err
			}
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

// Register adds a COPY pipe to the cron scheduler.
// If the pipe is already registered, it is replaced.
func (s *Scheduler) Register(pipe *pipetypes.Pipe) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if entryID, ok := s.entries[pipe.Name]; ok {
		s.cron.Remove(entryID)
	}

	pipeName := pipe.Name
	orchest := s.orchest
	entryID, err := s.cron.AddFunc(pipe.CopySchedule, func() {
		if err := orchest.RunCopyPipe(context.Background(), pipeName); err != nil {
			log.Printf("copy pipe %q run failed: %v", pipeName, err)
		}
	})
	if err != nil {
		return err
	}
	s.entries[pipe.Name] = entryID
	return nil
}

// Unregister removes a COPY pipe from the cron scheduler.
func (s *Scheduler) Unregister(pipeID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if entryID, ok := s.entries[pipeID]; ok {
		s.cron.Remove(entryID)
		delete(s.entries, pipeID)
	}
}
