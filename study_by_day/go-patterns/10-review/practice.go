package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ========================================
// SINGLETON: Logger
// ========================================

type Logger struct{ prefix string }

var (
	loggerInstance *Logger
	loggerOnce     sync.Once
)

func GetLogger() *Logger {
	loggerOnce.Do(func() {
		loggerInstance = &Logger{prefix: "[JOB-SYSTEM]"}
	})
	return loggerInstance
}

func (l *Logger) Log(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s %s\n", l.prefix, time.Now().Format("15:04:05.000"), msg)
}

// ========================================
// FACTORY: Job Types
// ========================================

type JobType string

const (
	JobTypeEmail   JobType = "email"
	JobTypeReport  JobType = "report"
	JobTypeCleanup JobType = "cleanup"
)

type Job struct {
	ID      int
	Type    JobType
	Payload string
}

// TODO: Implement NewJob(id int, jobType string) Job
// Create a Job with the given ID and type, set Payload = fmt.Sprintf("payload-%d", id)
func NewJob(id int, jobType string) Job {
	// TODO: implement
	return Job{}
}

// ========================================
// PIPELINE: validate → enrich
// ========================================

// TODO: validateJobs — filter out jobs with empty Payload
func validateJobs(in <-chan Job) <-chan Job {
	out := make(chan Job)
	// TODO: forward valid jobs, discard invalid
	return out
}

// TODO: enrichJobs — add " [enriched]" to each job's Payload
func enrichJobs(in <-chan Job) <-chan Job {
	out := make(chan Job)
	// TODO: implement
	return out
}

// ========================================
// RATE LIMITER
// ========================================

type RateLimiter struct {
	ticker *time.Ticker
}

func NewRateLimiter(ratePerSecond int) *RateLimiter {
	interval := time.Second / time.Duration(ratePerSecond)
	return &RateLimiter{ticker: time.NewTicker(interval)}
}

func (r *RateLimiter) Wait() {
	<-r.ticker.C
}

// ========================================
// CIRCUIT BREAKER (simplified)
// ========================================

type CB struct {
	mu          sync.Mutex
	failures    int
	maxFailures int
	open        bool
	openedAt    time.Time
	timeout     time.Duration
}

func NewCB(maxFailures int) *CB {
	return &CB{maxFailures: maxFailures, timeout: 3 * time.Second}
}

func (cb *CB) Call(fn func() error) error {
	cb.mu.Lock()
	if cb.open && time.Since(cb.openedAt) > cb.timeout {
		cb.open = false
		cb.failures = 0
	}
	if cb.open {
		cb.mu.Unlock()
		return fmt.Errorf("circuit open")
	}
	cb.mu.Unlock()

	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()
	if err != nil {
		cb.failures++
		if cb.failures >= cb.maxFailures {
			cb.open = true
			cb.openedAt = time.Now()
			GetLogger().Log("⚡ CIRCUIT OPENED after %d failures", cb.failures)
		}
	} else {
		cb.failures = 0
	}
	return err
}

// ========================================
// RETRY
// ========================================

func withRetry(maxRetries int, fn func() error) error {
	delay := 100 * time.Millisecond
	for i := 0; i <= maxRetries; i++ {
		if err := fn(); err == nil {
			return nil
		} else if i < maxRetries {
			time.Sleep(delay)
			delay *= 2
		} else {
			return err
		}
	}
	return nil
}

// ========================================
// WORKER (FAN-OUT)
// ========================================

type Result struct {
	JobID  int
	Output string
	Err    error
}

// Simulated processor — fails randomly 30% of the time
func processJob(job Job) error {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	if rand.Float32() < 0.3 {
		return fmt.Errorf("processing failed for job %d", job.ID)
	}
	return nil
}

// TODO: Implement runWorker
// - reads jobs from jobs channel
// - stops when ctx is cancelled
// - for each job: use circuit breaker + retry to call processJob
// - send Result to results channel
func runWorker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, cb *CB, wg *sync.WaitGroup) {
	defer wg.Done()
	log := GetLogger()

	for {
		select {
		case <-ctx.Done():
			log.Log("Worker %d: context cancelled", id)
			return
		case job, ok := <-jobs:
			if !ok {
				log.Log("Worker %d: jobs channel closed", id)
				return
			}
			// TODO: use withRetry + cb.Call to process job
			// Send result regardless of success/failure
			_ = job
		}
	}
}

// ========================================
// FAN-IN: merge results
// ========================================

func fanInResults(channels ...<-chan Result) <-chan Result {
	merged := make(chan Result, 10)
	var wg sync.WaitGroup

	forward := func(ch <-chan Result) {
		defer wg.Done()
		for r := range ch {
			merged <- r
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go forward(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

// ========================================
// MAIN
// ========================================

func main() {
	log := GetLogger()
	log.Log("=== Job Processing System Starting ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Rate limiter: 3 jobs per second
	limiter := NewRateLimiter(3)

	// Circuit breaker: open after 3 failures
	cb := NewCB(3)

	// Job channel and results per worker
	jobs := make(chan int, 20)
	const numWorkers = 3
	workerJobs := make(chan Job, 20)
	var workerResults [numWorkers]chan Result
	for i := range workerResults {
		workerResults[i] = make(chan Result, 20)
	}

	// Generate 10 jobs
	go func() {
		jobTypes := []string{"email", "report", "cleanup"}
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Pipeline: read raw job IDs → create Jobs → validate → enrich → send to workers
	go func() {
		rawJobs := make(chan Job, 20)

		// Create jobs from IDs
		go func() {
			jobTypes := []string{"email", "report", "cleanup"}
			for id := range jobs {
				limiter.Wait()
				job := NewJob(id, jobTypes[id%3])
				log.Log("Created job %d (type: %s)", job.ID, job.Type)
				rawJobs <- job
			}
			close(rawJobs)
		}()

		validated := validateJobs(rawJobs)
		enriched := enrichJobs(validated)

		for job := range enriched {
			workerJobs <- job
		}
		close(workerJobs)
	}()

	// Fan-out: 3 workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go runWorker(ctx, i+1, workerJobs, workerResults[i], cb, &wg)
	}

	// Close result channels when all workers done
	go func() {
		wg.Wait()
		for i := range workerResults {
			close(workerResults[i])
		}
	}()

	// Fan-in: collect all results
	resultChs := make([]<-chan Result, numWorkers)
	for i := range workerResults {
		resultChs[i] = workerResults[i]
	}
	allResults := fanInResults(resultChs...)

	// Print results
	success, failed := 0, 0
	for r := range allResults {
		if r.Err != nil {
			log.Log("❌ Job %d FAILED: %v", r.JobID, r.Err)
			failed++
		} else {
			log.Log("✅ Job %d SUCCESS: %s", r.JobID, r.Output)
			success++
		}
	}

	log.Log("=== Done: %d success, %d failed ===", success, failed)
}

/*
HINTS for runWorker:

func runWorker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, cb *CB, wg *sync.WaitGroup) {
	defer wg.Done()
	log := GetLogger()
	for {
		select {
		case <-ctx.Done():
			log.Log("Worker %d: context cancelled", id)
			return
		case job, ok := <-jobs:
			if !ok { return }
			err := withRetry(2, func() error {
				return cb.Call(func() error {
					return processJob(job)
				})
			})
			output := ""
			if err == nil {
				output = fmt.Sprintf("processed %s: %s", job.Type, job.Payload)
			}
			results <- Result{JobID: job.ID, Output: output, Err: err}
		}
	}
}
*/
