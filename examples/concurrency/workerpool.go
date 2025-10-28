package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(ctx context.Context, id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "shutting down")
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("worker", id, "no more jobs")
				return
			}
			// process job
			fmt.Printf("worker %d processing job %d\n", id, job.ID)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan Job, numJobs)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, &wg)
	}

	// produce jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs) // signal no more jobs

	// wait for workers to finish processing
	wg.Wait()

	fmt.Println("all jobs processed")
}

/*
Worker pool: producer pushes jobs into a buffered channel; a fixed number of workers (goroutines) process jobs. Use WaitGroup to wait for workers to finish and context for graceful shutdown.
*/
