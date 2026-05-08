package orchestratortypes

import "time"

// TimeInterval is a half-open time range used for incremental backfill runs.
type TimeInterval struct {
	Start time.Time
	End   time.Time
}
