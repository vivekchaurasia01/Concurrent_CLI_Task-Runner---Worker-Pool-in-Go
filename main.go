package main

import (
	"context"
	"sync"
	"time"
)

type Work interface {
	run (ctx context.Context) error
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

func NewPool () *Pool {
	return &Pool{

		works: make(chan Work),
		results: make(chan Result),
	}
}


