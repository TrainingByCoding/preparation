# Complete Guide to Channels and Goroutines in Go

## Table of Contents
1. [Goroutines Basics](#goroutines-basics)
2. [Unbuffered Channels](#unbuffered-channels)
3. [Buffered Channels](#buffered-channels)
4. [Deadlock Scenarios](#deadlock-scenarios)
5. [Channel Synchronization Patterns](#channel-synchronization-patterns)
6. [Best Practices](#best-practices)

---

## Goroutines Basics

### What is a Goroutine?
A goroutine is a lightweight thread managed by the Go runtime. Use `go` keyword to start one.

### Example 1: Basic Goroutine
```go
package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go printNumbers()  // runs in separate goroutine
    time.Sleep(1 * time.Second)  // wait for goroutine to finish
    fmt.Println("Done")
}
```

**Key Points:**
- Main goroutine doesn't wait for spawned goroutines
- Need synchronization mechanisms (channels, WaitGroups) to coordinate
- Goroutines are cheap (can create thousands)

---

## Unbuffered Channels

### Definition
**Unbuffered channel = Synchronous communication**
- Sender blocks until receiver is ready
- Receiver blocks until sender sends
- Direct handoff between goroutines

### Example 2: Unbuffered Channel - Success
```go
package main

import "fmt"

func main() {
    ch := make(chan string)  // unbuffered channel
    
    // Goroutine 1: Sender
    go func() {
        ch <- "Hello"  // blocks until someone receives
        fmt.Println("Sent message")
    }()
    
    // Main goroutine: Receiver
    msg := <-ch  // blocks until someone sends
    fmt.Println("Received:", msg)
}
```

**Flow:**
1. Main creates goroutine
2. Goroutine tries to send → **blocks**
3. Main tries to receive → **receives immediately** (goroutine unblocks)
4. Both continue

### Example 3: Unbuffered Channel - Deadlock!
```go
package main

import "fmt"

func main() {
    ch := make(chan string)  // unbuffered
    
    // WRONG: Same goroutine send & receive
    ch <- "Hello"  // ❌ Blocks forever! No one to receive!
    msg := <-ch    // Never reached
    fmt.Println(msg)
}
```

**Why Deadlock?**
- Line 7: `ch <- "Hello"` blocks, waiting for receiver
- Line 8: `<-ch` never executes because Line 7 is stuck
- **Same goroutine cannot be both sender and receiver on unbuffered channel**

---

## Buffered Channels

### Definition
**Buffered channel = Asynchronous communication**
- Has capacity to hold values
- Sender only blocks when buffer is FULL
- Receiver only blocks when buffer is EMPTY

### Example 4: Buffered Channel - Success
```go
package main

import "fmt"

func main() {
    ch := make(chan int, 2)  // buffer capacity = 2
    
    ch <- 1  // ✅ Doesn't block (buffer: [1])
    ch <- 2  // ✅ Doesn't block (buffer: [1, 2])
    // ch <- 3  // ❌ Would block (buffer full)
    
    fmt.Println(<-ch)  // Prints: 1 (buffer: [2])
    fmt.Println(<-ch)  // Prints: 2 (buffer: [])
}
```

**Key Points:**
- Send doesn't block until buffer is full
- Same goroutine can send AND receive (using buffer as temporary storage)

### Example 5: Buffer Full - Blocks
```go
package main

import "fmt"

func main() {
    ch := make(chan int, 2)  // capacity = 2
    
    ch <- 1  // OK (buffer: [1])
    ch <- 2  // OK (buffer: [1, 2])
    ch <- 3  // ❌ DEADLOCK! Buffer full, no receiver
    
    fmt.Println(<-ch)
}
```

---

## Deadlock Scenarios

### Scenario 1: No Receiver (Unbuffered)
```go
func main() {
    ch := make(chan int)
    ch <- 42  // ❌ DEADLOCK: No goroutine to receive
}
```

### Scenario 2: No Sender (Unbuffered)
```go
func main() {
    ch := make(chan int)
    val := <-ch  // ❌ DEADLOCK: No goroutine to send
}
```

### Scenario 3: Buffer Full, No Receiver
```go
func main() {
    ch := make(chan int, 1)
    ch <- 1  // OK
    ch <- 2  // ❌ DEADLOCK: Buffer full, no receiver
}
```

### Scenario 4: Circular Dependency
```go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() {
        ch1 <- (<-ch2)  // Waits for ch2
    }()
    
    go func() {
        ch2 <- (<-ch1)  // Waits for ch1
    }()
    
    time.Sleep(1 * time.Second)  // ❌ Both goroutines deadlocked
}
```

---

## Channel Synchronization Patterns

### Pattern 1: Simple Producer-Consumer
```go
package main

import "fmt"

func producer(ch chan int) {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)  // Signal no more values
}

func consumer(ch chan int) {
    for val := range ch {  // Reads until channel closed
        fmt.Println("Received:", val)
    }
}

func main() {
    ch := make(chan int)
    go producer(ch)
    consumer(ch)
}
```

### Pattern 2: Worker Pool
```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    var wg sync.WaitGroup
    
    // Start 3 workers
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Wait for workers
    wg.Wait()
    close(results)
    
    // Collect results
    for result := range results {
        fmt.Println("Result:", result)
    }
}
```

### Pattern 3: Ping-Pong (Alternating Goroutines)
```go
package main

import "fmt"

func ping(pings chan<- string, pongs <-chan string) {
    for i := 0; i < 5; i++ {
        pings <- "ping"
        <-pongs  // Wait for pong
    }
}

func pong(pings <-chan string, pongs chan<- string) {
    for i := 0; i < 5; i++ {
        <-pings  // Wait for ping
        pongs <- "pong"
    }
}

func main() {
    pings := make(chan string)
    pongs := make(chan string)
    
    go ping(pings, pongs)
    go pong(pings, pongs)
    
    time.Sleep(1 * time.Second)
}
```

### Pattern 4: Done Signal
```go
package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Println("Working...")
    time.Sleep(1 * time.Second)
    fmt.Println("Done")
    done <- true  // Signal completion
}

func main() {
    done := make(chan bool)
    go worker(done)
    <-done  // Block until worker signals done
    fmt.Println("All workers finished")
}
```

### Pattern 5: Select Statement (Multiple Channels)
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "from ch1"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "from ch2"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
}
```

---

## Best Practices

### 1. Close Channels
- **Only sender** should close channels
- Receivers can detect closure: `val, ok := <-ch`
- Range loop exits when channel closed

```go
ch := make(chan int)
go func() {
    for i := 0; i < 3; i++ {
        ch <- i
    }
    close(ch)  // Sender closes
}()

for val := range ch {  // Exits when closed
    fmt.Println(val)
}
```

### 2. Use Directional Channels
```go
func send(ch chan<- int) {  // Send-only
    ch <- 42
}

func receive(ch <-chan int) {  // Receive-only
    val := <-ch
}
```

### 3. Avoid Leaking Goroutines
```go
// ❌ BAD: Goroutine may leak if context cancelled
func leaky() {
    ch := make(chan int)
    go func() {
        val := <-ch  // Blocks forever if no sender
        fmt.Println(val)
    }()
}

// ✅ GOOD: Use context or timeout
func safe(ctx context.Context) {
    ch := make(chan int)
    go func() {
        select {
        case val := <-ch:
            fmt.Println(val)
        case <-ctx.Done():
            return  // Exit if cancelled
        }
    }()
}
```

### 4. Choose Right Channel Type

| Use Case | Channel Type |
|----------|--------------|
| Strict synchronization needed | Unbuffered |
| Decouple sender/receiver timing | Buffered (small) |
| High throughput, bursty traffic | Buffered (large) |
| Simple signaling (done/cancel) | Unbuffered or struct{} |

---

## Common Mistakes

### Mistake 1: Forgetting to Close Channel
```go
// ❌ Range loop blocks forever
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    // Forgot to close!
}()

for val := range ch {  // Never exits
    fmt.Println(val)
}
```

### Mistake 2: Closing Channel Multiple Times
```go
ch := make(chan int)
close(ch)
close(ch)  // ❌ PANIC: close of closed channel
```

### Mistake 3: Sending on Closed Channel
```go
ch := make(chan int)
close(ch)
ch <- 1  // ❌ PANIC: send on closed channel
```

### Mistake 4: Not Handling Buffer Size
```go
// ❌ Buffer too small, may deadlock
ch := make(chan int, 1)
for i := 0; i < 10; i++ {
    ch <- i  // Blocks on i=2
}
```

---

## Quick Reference Chart

| Feature | Unbuffered | Buffered |
|---------|------------|----------|
| Creation | `make(chan T)` | `make(chan T, n)` |
| Send blocks? | Always (until received) | Only when full |
| Receive blocks? | Always (until sent) | Only when empty |
| Synchronous? | Yes | No |
| Same goroutine send/receive? | ❌ Deadlock | ✅ OK (if buffer not full) |
| Use case | Synchronization | Throughput, decoupling |

---

## Practice Tips

1. **Start simple**: Single sender, single receiver
2. **Add complexity**: Multiple senders/receivers
3. **Use `go run -race`**: Detect race conditions
4. **Print states**: Add logs to understand flow
5. **Experiment with buffer sizes**: See blocking behavior
6. **Use select**: Handle multiple channels
7. **Profile**: Use `pprof` for goroutine leaks

---

## Summary

- **Goroutines**: Lightweight concurrent functions
- **Unbuffered channels**: Direct synchronization (handoff)
- **Buffered channels**: Asynchronous with capacity
- **Deadlock**: All goroutines blocked, no progress possible
- **Close channels**: Signal no more values (sender responsibility)
- **Select**: Handle multiple channel operations
- **Patterns**: Producer-consumer, worker pool, ping-pong, done signal

**Golden Rule**: Unbuffered = Synchronous, Buffered = Asynchronous (until full)
