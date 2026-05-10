package main

import (
	"fmt"
	"sync"
	"time"
)

// ========================================
// Exercise 1: Basic Fan-Out
// ========================================

// TODO: Implement fanOut
// - takes a jobs channel and numWorkers
// - starts numWorkers goroutines, each reads from jobs and squares the number
// - returns a results channel
// - closes results channel when all workers are done
func fanOut(jobs <-chan int, numWorkers int) <-chan int {
	results := make(chan int, numWorkers)
	var wg sync.WaitGroup

	// TODO: implement workers and WaitGroup cleanup

	_ = wg // remove when implemented
	return results
}

func makeJobs(nums ...int) <-chan int {
	ch := make(chan int, len(nums))
	for _, n := range nums {
		ch <- n
	}
	close(ch)
	return ch
}

func exercise1() {
	jobs := makeJobs(1, 2, 3, 4, 5)
	results := fanOut(jobs, 3)
	for r := range results {
		fmt.Println(r)
	}
}

// ========================================
// Exercise 2: Identified Workers
// ========================================

// TODO: Implement fanOutIdentified
// Same as fanOut but each worker prints: fmt.Printf("Worker %d processed job %d → %d\n", id, job, job*job)
func fanOutIdentified(jobs <-chan int, numWorkers int) <-chan int {
	results := make(chan int, numWorkers)
	var wg sync.WaitGroup

	// TODO: implement — pass worker ID to each goroutine

	_ = wg
	return results
}

func exercise2() {
	jobs := makeJobs(10, 20, 30, 40, 50)
	results := fanOutIdentified(jobs, 3)
	for r := range results {
		_ = r // results already printed inside workers
	}
}

// ========================================
// Exercise 3: Fan-Out + Fan-In Pipeline
// ========================================

// Reuse fanOut from above
// TODO: also implement fanIn to merge multiple result channels
// (Or reuse the fanIn pattern from day33 — same concept)

func fanIn(channels ...<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	// TODO: implement — same as day33 fanIn
	_ = wg
	return merged
}

func exercise3() {
	// Send 10 jobs
	jobs := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	// TODO: fan-out to 3 workers, collect results, print them
	results := fanOut(jobs, 3)
	_ = results
}

// ========================================
// Exercise 4: URL Fetcher Simulation
// ========================================

type FetchResult struct {
	URL    string
	Status int
}

func simulateFetch(url string) FetchResult {
	// Simulate variable network latency
	time.Sleep(time.Duration(len(url)*10) * time.Millisecond)
	return FetchResult{URL: url, Status: 200}
}

func exercise4() {
	urls := []string{
		"https://api.service1.com/data",
		"https://api.service2.com/users",
		"https://api.service3.com/orders",
		"https://api.service4.com/products",
		"https://api.service5.com/inventory",
	}

	urlCh := make(chan string, len(urls))
	for _, u := range urls {
		urlCh <- u
	}
	close(urlCh)

	results := make(chan FetchResult, len(urls))
	var wg sync.WaitGroup

	// TODO: Start 3 worker goroutines
	// Each reads URLs from urlCh, calls simulateFetch, sends result to results
	// Close results when all workers done

	_ = wg
	_ = results
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 34: Fan-Out Pattern ===\n")

	fmt.Println("Exercise 1: Basic Fan-Out (squares of 1-5)")
	exercise1()

	fmt.Println("\nExercise 2: Identified Workers")
	exercise2()

	fmt.Println("\nExercise 3: Fan-Out + Fan-In (jobs 1-10)")
	exercise3()

	fmt.Println("\nExercise 4: URL Fetcher (3 concurrent fetchers)")
	exercise4()
}

/*
SOLUTIONS:

// Exercise 1 — fanOut:
func fanOut(jobs <-chan int, numWorkers int) <-chan int {
	results := make(chan int, numWorkers)
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for job := range jobs {
				results <- job * job
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// Exercise 2:
func fanOutIdentified(jobs <-chan int, numWorkers int) <-chan int {
	results := make(chan int, numWorkers)
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for id := 1; id <= numWorkers; id++ {
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				result := job * job
				fmt.Printf("Worker %d: job %d → %d\n", id, job, result)
				results <- result
			}
		}(id)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// Exercise 4:
wg.Add(3)
for i := 0; i < 3; i++ {
	go func() {
		defer wg.Done()
		for url := range urlCh {
			results <- simulateFetch(url)
		}
	}()
}
go func() {
	wg.Wait()
	close(results)
}()
for r := range results {
	fmt.Printf("Fetched %s → %d\n", r.URL, r.Status)
}
*/
