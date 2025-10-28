# Go-Learning: Week 3 – Deep Dive into Concurrency

## Overview

This document is a deep dive into Go's concurrency model. It explains goroutines, channels, select, synchronization primitives, the context package, common patterns (worker pool), race detection, best practices and interview-style Q&A.

---

## 1. Concurrency vs Parallelism

- **Concurrency**: managing multiple tasks that can make progress.
- **Parallelism**: executing multiple tasks at the same time (requires multiple cores).

Go's concurrency model enables both: goroutines are scheduled by the Go runtime and can be executed on multiple OS threads (parallelism) when `GOMAXPROCS` > 1.

---

## 2. Goroutines

- Created with `go f()`.
- Cheap: small stack that grows/shrinks.
- Scheduling: the Go scheduler multiplexes goroutines onto OS threads.
- Lifespan: until function returns (or goroutine is blocked forever).
- Leaks: goroutine leaks occur when goroutines block forever (e.g., waiting on a channel or blocked network call) and never exit.

**Example pattern:**
- Fire-and-forget tasks (use with caution).
- Long-running workers — accept `context.Context` to allow graceful shutdown.

**Best practice:** Accept `context.Context` in functions that perform work and might need cancellation.

---

## 3. Data Races and the Race Detector

- **Data race**: two goroutines access same variable concurrently, with at least one write and no synchronization.
- Data races are undefined behavior: they can corrupt state and cause intermittent bugs.
- **Go race detector** (`-race`) instruments memory accesses at runtime and reports races. It is invaluable during development.

**How to use:**
```bash
go run -race ./path/to/program.go
go test -race ./...
```

**Windows note:** the race detector requires cgo (CGO_ENABLED=1) and a working C compiler (`gcc`) in PATH.

---

## 4. Channels

- Channels are typed: `chan T`.
- **Unbuffered** channels synchronize sender and receiver.
- **Buffered** channels have capacity `make(chan T, cap)`.
- Close channels to signal no more values; receivers can range over channels until they close.
- Do **not** close channels from the receiving side; the producer/owner should close.

**Operations:**
- Send: `ch <- v`
- Receive: `v := <-ch` or `v, ok := <-ch`
- Close: `close(ch)`

**Select:** choose among multiple channel ops.

---

## 5. Synchronization primitives (sync package)

- `sync.WaitGroup` — wait for set of goroutines.
- `sync.Mutex` — mutual exclusion lock.
- `sync.RWMutex` — allow concurrent readers, exclusive writer.
- `sync/atomic` — lock-free atomic operations, useful for counters (fast).

Prefer channels for communication; mutexes when protecting shared memory is simpler.

---

## 6. Context

- `context.Context` is used to carry deadlines, cancellation signals and request-scoped values.
- Functions performing blocking or remotely bound work should accept `ctx context.Context`.
- Do **not** store `Context` in struct fields — pass explicitly.
- `context.WithCancel`, `context.WithTimeout`, `context.WithDeadline` create derived contexts.

Pattern:
```go
ctx, cancel := context.WithTimeout(parentCtx, time.Second)
defer cancel()
doWork(ctx)
```

---

## 7. Worker Pool Pattern

- Bounded concurrency: create `n` workers reading jobs from a job channel.
- Advantages: predictable resource usage, backpressure.
- Use buffered job channel to queue tasks.
- Use WaitGroup for clean shutdown, and Context for cancellation.

---

## 8. Design Patterns and Anti-Patterns

**Prefer**
- small goroutines that do a single responsibility
- channels for communication instead of shared state
- contexts for cancellation and timeouts
- `select` with `time.After` or `context.Done()` for timeouts

**Avoid**
- spawning unbounded goroutines in response to external input (DDOS risk)
- using `time.Sleep` for synchronization logic
- forgetting to close channels or cancel contexts (resource leaks)
- long critical sections under a mutex

---

## 9. Debugging and Profiling

- **Race detector**: `-race`
- **pprof**: CPU and memory profiles for performance issues
- **GODEBUG**: various runtime debug options (e.g., `gctrace=1`)  
- Monitor goroutine count (runtime.NumGoroutine()) to detect leaks.

---

## 10. Performance considerations

- `sync/atomic` is faster for simple counters than Mutex (less contention).
- Mutex contention can be mitigated by sharding (partitioning state).
- Channels incur copying cost for complex types — pass pointers if appropriate.
- For high-throughput systems use worker pools and bounded queues.

---

## Summary of week3
- Goroutines and how Go schedules them
- Concurrency vs Parallelism
- WaitGroups for synchronization
- Race conditions & debugging with -race
- Mutex locks for safe sharing
- Channels (unbuffered & buffered)
- select statement for multi-channel coordination
- Worker Pool design pattern
- Context for cancellation & timeouts
- Graceful shutdown of concurrent systems

**Repo folder references:**
- `examples/concurrency/basic.go` -> Goroutines basics
- `examples/concurrency/race.go` -> Race conditions
- `examples/concurrency/waitgroup.go` -> Synchronization
- `examples/concurrency/mutex.go` -> Mutex lock
- `examples/channels/channel_basics.go` -> Channels
- `examples/channels/buffered_channels.go` -> Buffered channels
- `examples/channels/select.go` -> Select statement
- `examples/context/cancel.go` -> Context cancel
- `examples/context/timeout.go` -> Timeout context
- `examples/concurrency/workerpool/` -> Mini Project

---

**End of Week 3 theory**
