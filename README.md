# Concurrent_CLI_Task Runner — Worker Pool in Go
 
# Worker Pool in Go

A concurrent job processing tool. Contain fixed number of workers, 
jobs distributed via channels.

## Why not launch a goroutine per job?

1000 goroutines competing for CPU causes scheduler thrashing.
A fixed pool keeps resources bounded regardless of job count.

## Design Decisions

**Buffered channel over mutex-protected slice**
Workers block on receive — zero CPU wasted while idle.
Backpressure is free: Submit() blocks when buffer is full.

**Context over done channel**
Context forms a tree — cancel parent, all children stop.
Every stdlib call (HTTP, DB) accepts context, so cancellation
flows all the way down automatically.

**sync.Once to close results**
Closing a closed channel panics. Multiple workers finish concurrently.
Once guarantees exactly one close regardless of race.

**Directional channels in worker signature**
`<-chan` and `chan<-` make wrong usage a compile error, not a runtime panic.

**Check ctx.Done() before each job**
Buffered channel may hold queued jobs after shutdown.
Without the check, workers drain the entire buffer after Ctrl+C.

## Run

```bash
go run Pool/Example/manual/main.go
