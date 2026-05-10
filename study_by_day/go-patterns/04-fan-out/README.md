# Day 34: Fan-Out Pattern

## 🎯 Goal
Master **Fan-Out** — distributing work from one channel to multiple workers. The backbone of parallel processing.

## 🧠 Key Concepts
- Fan-Out = **one producer → many workers**
- Workers compete to pull from the same input channel (Go channels are safe for concurrent reads)
- Controls concurrency: instead of unlimited goroutines, you fix N workers
- Combined with Fan-In = complete parallel pipeline

## 📖 Pattern

```
                ┌──► Worker 1 ──┐
Jobs Channel ───┼──► Worker 2 ──┼──► Results Channel
                └──► Worker 3 ──┘
```

### Implementation
```go
func fanOut(jobs <-chan int, numWorkers int) <-chan int {
    results := make(chan int, numWorkers)
    var wg sync.WaitGroup

    worker := func() {
        defer wg.Done()
        for job := range jobs {        // all workers compete for same jobs channel
            results <- job * job       // process job
        }
    }

    wg.Add(numWorkers)
    for i := 0; i < numWorkers; i++ {
        go worker()
    }

    go func() {
        wg.Wait()
        close(results) // close results only after ALL workers finish
    }()

    return results
}
```

## ✅ Learning Checklist
- [ ] Implement fan-out with N workers reading from shared channel
- [ ] Understand workers compete (not round-robin) for jobs
- [ ] Combine fan-out + fan-in into a full pipeline
- [ ] Know when to buffer the results channel

## 🛠️ Practice Exercises
1. **Basic Fan-Out** — 5 jobs, 3 workers, each worker squares the number
2. **Identified workers** — same as above but each worker prints its ID with the result
3. **Fan-Out + Fan-In pipeline** — send 10 jobs, 3 workers process them, fan-in results, print sorted
4. **URL fetcher simulation** — 10 URLs, 3 concurrent fetchers (use `time.Sleep` to simulate), print results as they arrive

## 🔍 Interview Questions
- "How do you limit concurrency in Go?"
- "What happens if workers finish before you close the results channel?"
- "How is fan-out different from just spawning a goroutine per job?"
