package pool

import (
	"context"
	"time"
)

func Worker (
	ctx context.Context,
	works <-chan Job,
	results chan <- Result,
) {
	// Before running the job, check if context is already cancelled.
    // Why? If shutdown happened while this job was waiting in the channel,
    // we don't want to start it.

	for job := range works {
		select {
		case <- ctx.Done():
			return
		default:
		}

		start := time.Now()
		err := job.Run(ctx) // Polymorphism.. actually do the work, pass context so job can also cancel itself
        
		results <- Result{
			JobID: job.ID(),
			Err: err,
			Duration: time.Since(start),
		}
		
	}
}