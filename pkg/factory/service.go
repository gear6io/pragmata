package factory

import "context"

// Service is the common lifecycle interface for long-running components.
type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
