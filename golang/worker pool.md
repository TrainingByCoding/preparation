# Worker Pool
 
Question:
You have a list of 10 tasks that need to be processed.
Each task takes around 1 second to complete.

Write a Go program that:
- Starts 2 worker goroutines
- Each worker picks up one task at a time from a shared queue
- Prints the task ID and the worker ID while processing
Waits for all tasks to finish before exiting

```go
package main
import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Worker is just pulling items off the 'jobs' channel
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(1 * time.Second) // Pretend to do some work
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	const numWorkers = 2 // ✅ We control the concurrency limit right here!
	const numJobs = 10

	jobs := make(chan int, numJobs) // The work queue
	var wg sync.WaitGroup

	// 1. We start up our fixed number of workers.
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// 2. Enqueue jobs into the jobs channel.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Important: Tell the workers we're done sending jobs

	// 3. We wait for every single worker to be finished.
	wg.Wait()
	fmt.Println("All jobs completed")
}

output:
Worker 2 started job 1
Worker 1 started job 2
Worker 1 finished job 2
Worker 1 started job 3
Worker 2 finished job 1
Worker 2 started job 4
Worker 2 finished job 4
Worker 2 started job 5
Worker 1 finished job 3
Worker 1 started job 6
Worker 1 finished job 6
Worker 1 started job 7
Worker 2 finished job 5
Worker 2 started job 8
Worker 1 finished job 7
Worker 1 started job 9
Worker 2 finished job 8
Worker 2 started job 10
Worker 1 finished job 9
Worker 2 finished job 10
All jobs completed
 
