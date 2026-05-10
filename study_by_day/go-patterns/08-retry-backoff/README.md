# Day 38: Retry with Backoff Pattern

## 🎯 Goal
Implement **retry with exponential backoff** — the standard way to handle transient failures in production systems.

## 🧠 Key Concepts
- **Retry** — try the operation again on failure (not all errors are worth retrying)
- **Exponential Backoff** — wait doubles each attempt: 1s → 2s → 4s → 8s
- **Jitter** — add random delay to avoid thundering herd (all retrying at same time)
- **Max retries + deadline** — must have a limit, not retry forever
- Retryable errors: network timeout, 503 Service Unavailable
- Non-retryable: 400 Bad Request, 401 Unauthorized, business logic errors

## 📖 Pattern

```go
func withRetry(maxRetries int, fn func() error) error {
    backoff := 100 * time.Millisecond

    for attempt := 0; attempt <= maxRetries; attempt++ {
        err := fn()
        if err == nil {
            return nil  // success
        }
        if attempt == maxRetries {
            return fmt.Errorf("failed after %d attempts: %w", maxRetries+1, err)
        }
        // Add jitter: backoff ± 20%
        jitter := time.Duration(rand.Int63n(int64(backoff / 5)))
        time.Sleep(backoff + jitter)
        backoff *= 2  // exponential: double each time
    }
    return nil
}
```

## ✅ Learning Checklist
- [ ] Implement simple retry (fixed delay)
- [ ] Implement exponential backoff
- [ ] Add jitter to prevent thundering herd
- [ ] Add context for deadline/cancellation
- [ ] Distinguish retryable vs non-retryable errors

## 🛠️ Practice Exercises
1. **Fixed retry** — retry up to 3 times with 500ms delay, print each attempt
2. **Exponential backoff** — double the wait each time: 100ms → 200ms → 400ms
3. **With jitter** — add ±20% random jitter to backoff
4. **Context-aware retry** — stop retrying if context is cancelled or deadline exceeded
5. **BONUS** — `RetryableError` type — only retry if error is tagged as retryable

## 🔍 Interview Questions
- "What is exponential backoff and why is it used?"
- "What is the thundering herd problem?"
- "How do you decide which errors to retry?"
