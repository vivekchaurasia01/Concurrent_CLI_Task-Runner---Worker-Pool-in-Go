package main

import (
	"context"
	"fmt"
	"time"

	pool "github.com/vivekchaurasia01/Concurrent_CLI_Task-Runner---Worker-Pool-in-Go/Pool"
)

type SleepJob struct {
    id       string
    duration time.Duration
}

func (j SleepJob) ID() string { 
	return j.id
}

func (j SleepJob) Run(ctx context.Context) error {
    select {
    case <-time.After(j.duration): // simulate work
        fmt.Printf("job %s done\n", j.id)
        return nil
    case <-ctx.Done(): // cancelled mid-job
        return ctx.Err()
    }
}

func main () {
	p := pool.NewPool(10, 1000)
	ctx := context.Background()
	p.Start(ctx)

	// Submit 20 jobs
    for i := 0; i < 1000; i++ {
        p.Submit(SleepJob {
			id: fmt.Sprintf("job-%d", i),
			duration: 100 * time.Millisecond,
		})
    }
    p.Stop() // close jobs channel — workers will drain and exit

    // Range over results — exits automatically when results channel closes
    // No more hardcoded "20" — this is now correct for any number of jobs
    for r := range p.Results() {
        fmt.Printf("job %s | err: %v | duration: %v\n", r.JobID, r.Err, r.Duration)
    }
}