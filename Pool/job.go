package pool

import (
	"context"
	"time"
)

type Job interface {
	Run (ctx context.Context) error
	ID () string
}

type Result struct {
	JobID string
	Err error
	Duration time.Duration
}




