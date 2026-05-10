# Day 12: Producer-Consumer Pattern

## 🎯 Today's Goal
Master the producer-consumer pattern - a fundamental concurrency pattern!

## 📚 Files to Study
1. `../../master100/producer_consumer.go` - Classic implementation

## 🧠 Key Concepts
- Producer goroutine (generates data)
- Consumer goroutine (processes data)
- Channel as the queue between them
- Synchronization and flow control
- Graceful shutdown

## ✅ Learning Checklist
- [ ] Understand producer role
- [ ] Understand consumer role
- [ ] Know why channel is used
- [ ] Can implement multiple producers
- [ ] Can implement multiple consumers
- [ ] Know when to close the channel

## 🔍 Pattern Breakdown

```go
// Producer: Generates data, sends to channel
func producer(ch chan<- int) {
    for i := 0; i < 10; i++ {
        ch <- i  // Send to channel
    }
    close(ch)  // Signal: no more data
}

// Consumer: Receives data, processes it
func consumer(ch <-chan int) {
    for item := range ch {  // Range until closed
        process(item)
    }
}

// Main: Connect them
func main() {
    ch := make(chan int)
    go producer(ch)
    consumer(ch)  // Blocks until channel closed
}
```

## 🛠️ Practice Exercises

### Exercise 1: Single Producer, Single Consumer
```go
// Producer generates numbers 1-100
// Consumer prints squares of numbers
```

### Exercise 2: Multiple Producers, Single Consumer
```go
// 3 producers each generate 10 numbers
// 1 consumer receives and sums all
```

### Exercise 3: Single Producer, Multiple Consumers
```go
// 1 producer generates tasks
// 3 consumers compete for tasks (first available gets it)
```

### Exercise 4: Pipeline Pattern
```go
// Stage 1: Generate numbers
// Stage 2: Square them
// Stage 3: Filter evens
// Stage 4: Print
// (Connect with channels!)
```

## 📝 Study Steps

### Step 1: Theory (5 min)

**Why this pattern?**
- Decouples data generation from processing
- Producers and consumers run independently
- Channel handles synchronization
- Natural backpressure (if consumer is slow, producer blocks)

**Channel types**:
```go
chan<- int   // Send-only (producer)
<-chan int   // Receive-only (consumer)
chan int     // Bidirectional
```

### Step 2: Study File (5 min)
Read `producer_consumer.go` and identify:
- Where producer sends data
- Where consumer receives data
- How channel is closed
- How consumer knows to stop

### Step 3: Practice (10 min)
Implement all 4 exercises

## 🎓 Key Patterns

### Pattern 1: Simple Pipeline
```go
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Usage
for n := range square(gen(1, 2, 3, 4)) {
    fmt.Println(n)  // 1, 4, 9, 16
}
```

### Pattern 2: Fan-Out (Multiple Consumers)
```go
// One producer, many consumers reading same channel
// Whoever reads first gets the job
jobs := make(chan int)
for w := 1; w <= 5; w++ {
    go worker(jobs)  // 5 workers competing
}
```

### Pattern 3: Fan-In (Multiple Producers)
```go
// Many producers, one channel
// All producers send to same channel
ch := make(chan int)
for p := 1; p <= 5; p++ {
    go producer(p, ch)
}
```

## 🚨 Common Mistakes

### Mistake 1: Forgetting to Close Channel
```go
// ❌ Bad: Consumer waits forever
go func() {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    // Forgot close(ch)!
}()
for val := range ch {  // Deadlock!
    fmt.Println(val)
}

// ✅ Good
go func() {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)  // Signal completion
}()
for val := range ch {
    fmt.Println(val)
}
```

### Mistake 2: Consumer Closes Channel
```go
// ❌ Bad: Consumer should NOT close
// Producer owns the channel, producer closes it
```

### Mistake 3: Multiple Closures
```go
// ❌ Bad: Multiple producers closing same channel
// Use sync.WaitGroup and close in main instead
```

## 🎯 Interview Tips

**Question**: "How do you share data between goroutines?"  
**Answer**: "Producer-consumer pattern with channels. Producer sends data, consumer receives. Channel handles synchronization."

**Question**: "What if producer is faster than consumer?"  
**Answer**: "Unbuffered channel creates backpressure - producer blocks. Buffered channel allows some queueing."

**Question**: "How does consumer know producer is done?"  
**Answer**: "Producer closes the channel. Consumer's range loop exits when channel is closed."

## ✅ Completion
- [ ] Studied producer_consumer.go
- [ ] Implemented single producer/consumer
- [ ] Implemented multiple producers
- [ ] Implemented multiple consumers
- [ ] Built a pipeline
- [ ] Can explain to interviewer
- [ ] Updated PROGRESS_TRACKER.md
- [ ] Confidence: ___/10

## 🔄 Spaced Repetition
- Review Day 9 (Channels)
- Review Day 11 (Goroutines)

## 🔜 Tomorrow: Day 13 - Worker Pool Pattern
Learn how to control concurrency with limited workers!
