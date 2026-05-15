package pool

import (
	"context"
	"time"
)

func Worker (
	ctx context.Context,
	works <-chan Work,
	results chan <- Result,
) {
	// Before running the job, check if context is already cancelled.
    // Why? If shutdown happened while this job was waiting in the channel,
    // we don't want to start it.

	for work := range works {
		select {
		case ctx.Done():
			return
		default:
		}

		start := time.Now()
		
		
	}
}