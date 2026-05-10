# Day 39: Context Cancellation Pattern

## 🎯 Goal
Master `context.Context` — the Go standard for cancellation, deadlines, and passing request-scoped values. **Every production Go service uses this.**

## 🧠 Key Concepts
- `context.Context` carries: cancellation signal, deadline, timeout, key-value pairs
- Always passed as **first argument**: `func DoWork(ctx context.Context, ...)`
- When ctx is cancelled → **all goroutines using it should stop**
- Three ways to create cancellable contexts:
  - `context.WithCancel(parent)` → manual cancel
  - `context.WithTimeout(parent, duration)` → auto-cancel after duration
  - `context.WithDeadline(parent, time.Time)` → cancel at specific time

## 📖 Pattern

### WithCancel — manual cancellation
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // always defer cancel to avoid goroutine leak

go func() {
    select {
    case <-ctx.Done():
        fmt.Println("Goroutine stopped:", ctx.Err())
        return
    case result := <-doWork():
        fmt.Println("Result:", result)
    }
}()

time.Sleep(2 * time.Second)
cancel() // stops the goroutine
```

### WithTimeout — auto-cancellation
```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

err := callExternalService(ctx) // passes ctx to respect timeout
if errors.Is(err, context.DeadlineExceeded) {
    fmt.Println("Request timed out")
}
```

## ✅ Learning Checklist
- [ ] Use `ctx.Done()` channel to stop goroutines cleanly
- [ ] Implement timeout with `WithTimeout`
- [ ] Propagate context through a call chain
- [ ] Use `context.WithValue` for request-scoped data (trace ID, user ID)
- [ ] Always `defer cancel()` to prevent context leak

## 🛠️ Practice Exercises
1. **Basic cancel** — start a worker goroutine, cancel it after 2 seconds
2. **Timeout** — simulate slow DB query with 3s delay, timeout context after 1s
3. **Context chain** — 3-level call chain (handler → service → DB), cancel propagates all the way
4. **WithValue** — pass request ID through context, read it at bottom of call chain
5. **BONUS** — cancel multiple goroutines with one context

## 🔍 Interview Questions
- "How do you stop a goroutine from outside?"
- "What is context.Background() vs context.TODO()?"
- "What happens if you forget defer cancel()?"
- "How do you pass authentication data through a request in Go?"
