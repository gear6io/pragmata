package scheduler

import "github.com/gear6io/pragmata/pkg/types/pipetypes"

// Scheduler manages cron schedules for COPY pipes.
type Scheduler interface {
	Register(pipe *pipetypes.Pipe) error
	Unregister(pipeID string)
}
