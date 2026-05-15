package pool

import (
	"context"
	"sync"
	"time"
)

type Work interface {
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

	works 		chan Work
	results 	chan Result

	wg 			sync.WaitGroup
}

func NewPool (workerCount int, buffersize int ) *Pool {
	return &Pool{
		workerCount: workerCount,
		works: make(chan Work, buffersize),
		results: make(chan Result, buffersize),
	}
}


