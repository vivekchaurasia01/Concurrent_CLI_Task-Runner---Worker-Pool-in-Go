# Concurrent_CLI_Task Runner — Worker Pool in Go

<p align="center">

<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=600&size=24&duration=2500&pause=1000&color=00ADD8&center=true&vCenter=true&width=700&lines=Concurrent+Worker+Pool+in+Go;Bounded+Concurrency+%7C+Graceful+Shutdown;Channels+%7C+Context+%7C+Goroutines" />

</p>

<p align="center">

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Concurrency](https://img.shields.io/badge/Concurrency-Goroutines-success?style=for-the-badge)
![Pattern](https://img.shields.io/badge/Pattern-Worker%20Pool-orange?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=for-the-badge)

![Context](https://img.shields.io/badge/Context-Cancellation-red?style=flat-square)
![Channels](https://img.shields.io/badge/Channels-Buffered-yellow?style=flat-square)
![Shutdown](https://img.shields.io/badge/Shutdown-Graceful-critical?style=flat-square)
![Race Safe](https://img.shields.io/badge/Race%20Safe-sync.Once-blueviolet?style=flat-square)

![Last Commit](https://img.shields.io/github/last-commit/vivekchaurasia01/Concurrent_CLI_Task-Runner---Worker-Pool-in-Go?style=flat-square)
![Repo Size](https://img.shields.io/github/repo-size/vivekchaurasia01/Concurrent_CLI_Task-Runner---Worker-Pool-in-Go?style=flat-square)
![Stars](https://img.shields.io/github/stars/vivekchaurasia01/Concurrent_CLI_Task-Runner---Worker-Pool-in-Go?style=flat-square)

</p>
 
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

