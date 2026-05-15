package pool

import (
	"context"
	"sync"
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

type Pool struct {
	workerCount int

	works 		chan Job
	results 	chan Result

	wg 			sync.WaitGroup
}

func NewPool (workerCount int, buffersize int ) *Pool {
	return &Pool{
		workerCount: workerCount,
		works: make(chan Job, buffersize),
		results: make(chan Result, buffersize),
	}
}


