package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
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
        fmt.Printf("%s done\n", j.id)
        return nil
    case <-ctx.Done(): // cancelled mid-job
        return ctx.Err()
    }
}

func main () {

    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()
    
	p := pool.NewPool(10, 1000)

	p.Start(ctx)

    // lets submit 500 jobs...
    for i := 0; i < 500; i++ {
        select {
        case <- ctx.Done():
            p.Stop()
            return

        default:
            p.Submit(SleepJob{
                id: fmt.Sprintf("jobId : %d", i),
                duration: 2000 * time.Millisecond,
            })
        }
    }
    p.Stop()

    completed := 0
    failed := 0
    for r := range p.Results() {
        if r.Err != nil {
            fmt.Printf("%s ,Error: %v, TimeDuration: %d", r.JobID,r.Err,r.Duration)
        } else {
            fmt.Printf("%s, ISDoneIn :%d\n", r.JobID,r.Duration)
            completed ++
        }
        
    }
    fmt.Printf("\ncompleted :%d, failed : %d jobs\n", completed,failed)
    
}