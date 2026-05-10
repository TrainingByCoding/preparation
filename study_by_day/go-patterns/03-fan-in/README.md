# Day 33: Fan-In Pattern

## 🎯 Goal
Master **Fan-In** — merging multiple channels into one. Critical for aggregating concurrent results.

## 🧠 Key Concepts
- Fan-In = **many producers → one consumer channel**
- Opposite of Fan-Out (one → many)
- Use case: multiple data sources (APIs, DB queries) running concurrently, merge results
- Implementation: one goroutine per input channel, all send to single output channel
- Close output channel only after ALL input channels are drained → use `sync.WaitGroup`

## 📖 Pattern

```
Producer A ──┐
Producer B ──┼──► merged channel ──► Consumer
Producer C ──┘
```

### Implementation
```go
func fanIn(channels ...<-chan int) <-chan int {
    merged := make(chan int)
    var wg sync.WaitGroup

    // For each input channel, start a goroutine that forwards its values
    forward := func(ch <-chan int) {
        defer wg.Done()
        for val := range ch {
            merged <- val
        }
    }

    wg.Add(len(channels))
    for _, ch := range channels {
        go forward(ch)
    }

    // Close merged channel when all forwarders are done
    go func() {
        wg.Wait()
        close(merged)
    }()

    return merged
}
```

## ✅ Learning Checklist
- [ ] Implement fanIn with variadic channels
- [ ] Understand why you need WaitGroup to close the merged channel
- [ ] Fan-In with done/cancel channel (early exit)
- [ ] Know the real-world use case (parallel API calls, merge results)

## 🛠️ Practice Exercises
1. **Basic Fan-In** — merge 3 channels (each sends 3 numbers), print all 9 in merged channel
2. **Parallel search** — simulate 3 search sources (web, image, news), fan-in results, print first 3 that arrive
3. **Fan-In with timeout** — merge 3 slow channels, use `time.After` to stop after 2 seconds
4. **BONUS** — Fan-In + Fan-Out pipeline: fan-out jobs to 3 workers, fan-in their results

## 🔍 Interview Questions
- "How do you merge results from multiple goroutines into one stream?"
- "What is the select statement used for in fan-in?"
- "Why do you need sync.WaitGroup in fan-in?"
