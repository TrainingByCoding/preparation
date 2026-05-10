# Day 13: Worker Pool Pattern

## 🎯 Today's Goal
Master the worker pool pattern - essential for interviews!

## 📚 Files to Study
1. `../../workerpool.go` - Complete worker pool implementation

## 🧠 Key Concepts
- Worker pool pattern (control concurrency)
- Job queue with channels
- WaitGroup for synchronization
- Graceful shutdown (closing channels)
- Why this is better than spawning unlimited goroutines

## ✅ Learning Checklist
- [ ] Understand the 3 components: jobs channel, workers, WaitGroup
- [ ] Know why we limit workers (prevent resource exhaustion)
- [ ] Can implement from scratch
- [ ] Know when to close the jobs channel
- [ ] Understand why workers exit after channel closes

## 🔍 Pattern Breakdown

```go
// 1. Create job queue
jobs := make(chan Job, 100)

// 2. Start fixed number of workers
for w := 1; w <= numWorkers; w++ {
    wg.Add(1)
    go worker(w, jobs, &wg)
}

// 3. Send jobs to queue
for j := 1; j <= numJobs; j++ {
    jobs <- j
}
close(jobs)  // Tell workers: no more jobs coming

// 4. Wait for all workers to finish
wg.Wait()
```

## 🛠️ Practice Exercises

### Exercise 1: Basic Worker Pool
```go
// Implement worker pool with 3 workers processing 10 jobs
// Each job: sleep 500ms, print job number
```

### Exercise 2: Worker Pool with Results
```go
// Process 20 numbers
// Return their squares in a results channel
// Use 4 workers
```

### Exercise 3: URL Fetcher
```go
// Fetch 10 URLs concurrently
// Use 3 workers
// Print which worker fetched which URL
```

### Exercise 4: Dynamic Jobs
```go
// Worker pool where jobs generate more jobs
// (e.g., web crawler finding new links)
```

## 📝 Study Steps

### Step 1: Read workerpool.go (7 min)
Trace the execution:
1. Main creates channels and starts workers
2. Workers block on jobs channel
3. Main sends jobs
4. Workers process jobs
5. Main closes jobs channel
6. Workers finish remaining jobs and exit
7. WaitGroup releases main

### Step 2: Draw the Flow (3 min)
Sketch on paper:
- Jobs channel (middle)
- Workers pulling from it (left)
- Main sending to it (right)
- WaitGroup coordinating shutdown

### Step 3: Practice (10 min)
Implement exercises from scratch

## 🎓 Why Worker Pools?

**Without worker pool** (bad):
```go
for _, job := range jobs {
    go process(job)  // Could spawn 10,000 goroutines!
}
// No control, resource exhaustion
```

**With worker pool** (good):
```go
// Only 5 goroutines, process 10,000 jobs
for w := 1; w <= 5; w++ {
    go worker(jobs)
}
// Controlled, efficient
```

## 🎯 Interview Tips

Common questions:
1. **"How do you limit concurrency in Go?"**
   → Worker pool pattern with limited workers

2. **"What happens if you don't close the jobs channel?"**
   → Workers wait forever, goroutine leak

3. **"Why use WaitGroup?"**
   → Ensure all workers finish before main exits

4. **"Buffered or unbuffered jobs channel?"**
   → Buffered if you know job count, acts as queue

## ✅ Completion
- [ ] Studied workerpool.go thoroughly
- [ ] Drew the flow diagram
- [ ] Implemented worker pool from scratch
- [ ] Can explain to interviewer
- [ ] Updated PROGRESS_TRACKER.md
- [ ] Confidence: ___/10

## 🔄 Spaced Repetition
- Review Day 9 (Channels Introduction)
- Review Day 12 (Producer-Consumer)

## 🔜 Tomorrow: Day 14 - Timeouts & Context
Learn how to prevent goroutines from running forever!
