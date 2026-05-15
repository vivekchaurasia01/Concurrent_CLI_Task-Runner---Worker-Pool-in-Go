package pool

import (
	"context"
	"sync"
)

type Pool struct {
	workerCount int

	jobs 		chan Job
	results 	chan Result

	wg 			sync.WaitGroup
	once        sync.Once      // guarantees results channel is closed exactly once
}

func NewPool (workerCount int, buffersize int ) *Pool {
	return &Pool{
		workerCount: workerCount,
		jobs: make(chan Job, buffersize),
		results: make(chan Result, buffersize),
	}
}

func (p *Pool) Start (ctx context.Context) {
	for i := 0; i <= p.workerCount; i++ {
		p.wg.Add(1)
		go func () {
			defer p.wg.Done()
			Worker(ctx, p.jobs, p.results)
		} ()
	}
	go func () {
		p.wg.Wait()
		p.once.Do(func() {   // Close on a closed channel = panic. Once guarantees this runs exactly once, even if Stop() is called multiple times.
			close(p.results)
		})
	} ()
}

func (p Pool) Submit (job Job) {
	
}