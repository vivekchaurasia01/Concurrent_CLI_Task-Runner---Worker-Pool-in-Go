package pool_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	pool "github.com/vivekchaurasia01/Concurrent_CLI_Task-Runner---Worker-Pool-in-Go/Pool"
)

// Fake job for testing...
type TestJob struct {
	id string
}

func (j TestJob) ID () string {
	return j.id
}

func (j TestJob) Run (ctx context.Context) error {
	time.Sleep(50 * time.Millisecond)
	return nil
}

// TEst - 01 (does pool processes all jobs)  
func TestPool_AllJobsComplete(t *testing.T) {
	p := pool.NewPool(10,100)
	ctx := context.Background()

	p.Start(ctx)

	totalJobs := 100

	go func () {
		for i := 0; i < totalJobs; i++ {
			p.Submit(TestJob{id: fmt.Sprintf("job-%d",i)})
		}
		p.Stop()
	} ()

	completed := 0

	for r := range p.Results () {
		if r.Err != nil {
			t.Errorf("job %s failed : %v", r.JobID, r.Err)
		}
		completed ++
	}
	if completed != totalJobs {
		t.Errorf("expected %d results, got %d", totalJobs, completed)
	}
}

// test - 02 Go_routine race Detection...

func TestGo_RoutineLeak (t *testing.T) {
	p := pool.NewPool(10,100)
	ctx := context.Background()

	p.Start(ctx)

	totalJobs := 0

	go func () {
	for i := 0; i < totalJobs; i++ {
		p.Submit(TestJob{id: fmt.Sprintf("job-%d",i)})
	}
	p.Stop()
	} ()

	for range p.Results() {}  //doesn't care about the VALUES,it just keeps pulling items out, making space so workers never get stuck.
}

// Test - 03 GraceFulShutDown...

func TestPool_GracefulShutdown(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())

    p := pool.NewPool(10, 100)
    p.Start(ctx)

	  // lets submit 100 jobs...
    go func() {
        for i := 0; i < 100; i++ {
            select {
            case <-ctx.Done():
                p.Stop()
                return
            default:
                p.Submit(TestJob{id: fmt.Sprintf("job-%d", i)})
            }
        }
        p.Stop()
    }()

	// cancel after 50ms — simulates Ctrl+C
    time.AfterFunc(50*time.Millisecond, cancel)

    for range p.Results() {}
}