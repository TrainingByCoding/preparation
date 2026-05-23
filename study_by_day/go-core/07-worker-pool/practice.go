package main

import (
	"fmt"
	"sync"
)

// ========================================
// Exercise 1: Basic Worker Pool
// ========================================
func basicWorkerPool() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// TODO: Start workers
	// Each worker should:
	// - Process jobs from channel
	// - Sleep 500ms
	// - Print "Worker X processing job Y"

	// TODO: Send jobs

	// TODO: Close channel and wait
}

// ========================================
// Exercise 2: Worker Pool with Results
// ========================================
func workerPoolWithResults() {
	const numWorkers = 4
	const numJobs = 20

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// TODO: Worker function that squares numbers
	worker := func(id int, jobs <-chan int, results chan<- int) {
		defer wg.Done()
		// Your code here
	}

	// TODO: Start workers

	// TODO: Send jobs (1 to 20)

	// TODO: Close jobs channel

	// TODO: Collect results in a separate goroutine

	// TODO: Wait and print results
}

// ========================================
// Exercise 3: URL Fetcher Worker Pool
// ========================================
func urlFetcher() {
	urls := []string{
		"http://example.com/1",
		"http://example.com/2",
		"http://example.com/3",
		"http://example.com/4",
		"http://example.com/5",
		"http://example.com/6",
		"http://example.com/7",
		"http://example.com/8",
		"http://example.com/9",
		"http://example.com/10",
	}

	const numWorkers = 3
	jobs := make(chan string, len(urls))
	var wg sync.WaitGroup

	// TODO: Create worker that "fetches" URL (simulate with sleep)

	// TODO: Start workers

	// TODO: Send URLs to jobs channel

	// TODO: Close and wait
}

// ========================================
// Exercise 4: Dynamic Jobs (Advanced)
// ========================================
type Task struct {
	ID         int
	CreateMore bool // If true, worker creates 2 more jobs
}

func dynamicWorkerPool() {
	// TODO: Implement worker pool where some jobs create more jobs
	// Hint: Use sync.WaitGroup carefully
	// Hint: Track when to close the jobs channel

	// This is tricky! Think about:
	// - When do you close the jobs channel?
	// - How do you track pending jobs?
}

// ========================================
// Main
// ========================================
func main() {
	fmt.Println("=== Exercise 1: Basic Worker Pool ===")
	basicWorkerPool()

	fmt.Println("\n=== Exercise 2: Worker Pool with Results ===")
	workerPoolWithResults()

	fmt.Println("\n=== Exercise 3: URL Fetcher ===")
	urlFetcher()

	fmt.Println("\n=== Exercise 4: Dynamic Jobs (Advanced) ===")
	fmt.Println("(Try this if you have extra time)")
	// dynamicWorkerPool()
}

/*
SOLUTIONS:

func basicWorkerPool() {
	const numWorkers = 3
	const numJobs = 10
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("Worker %d processing job %d\n", id, j)
			}
		}(w)
	}
	for j := 1; j <= numJobs; j++ { jobs <- j }
	close(jobs)
	wg.Wait()
}

func workerPoolWithResults() {
	const numWorkers = 4
	const numJobs = 20
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup
	worker := func(id int, jobs <-chan int, results chan<- int) {
		defer wg.Done()
		for j := range jobs { results <- j * j }
	}
	for w := 1; w <= numWorkers; w++ { wg.Add(1); go worker(w, jobs, results) }
	for j := 1; j <= numJobs; j++ { jobs <- j }
	close(jobs)
	go func() { wg.Wait(); close(results) }()
	for r := range results { fmt.Println(r) }
}

func urlFetcher() {
	urls := []string{
		"http://example.com/1","http://example.com/2","http://example.com/3",
		"http://example.com/4","http://example.com/5","http://example.com/6",
		"http://example.com/7","http://example.com/8","http://example.com/9",
		"http://example.com/10",
	}
	const numWorkers = 3
	jobs := make(chan string, len(urls))
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for url := range jobs {
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("Worker %d fetched: %s\n", id, url)
			}
		}(w)
	}
	for _, url := range urls { jobs <- url }
	close(jobs)
	wg.Wait()
}
*/
