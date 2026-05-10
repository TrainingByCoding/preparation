# Day 40: Go Patterns Review — Mini-Project

## 🎯 Goal
Build a **mini job processing system** combining all patterns from days 31–39. This is your proof that you can assemble patterns together.

## 🏗️ What You're Building

A concurrent job processor with:
- **Singleton** — one shared logger
- **Factory** — create different job types
- **Pipeline** — parse → validate → process → log
- **Worker Pool (Fan-Out)** — 3 workers process jobs concurrently
- **Fan-In** — collect all results into one stream
- **Rate Limiter** — max 5 jobs per second
- **Circuit Breaker** — if processor fails 3 times, open circuit
- **Retry** — retry failed jobs up to 2 times
- **Context** — overall timeout; cancel all work if deadline exceeded

## 📋 Architecture

```
HTTP Request
    │
    ▼
[Rate Limiter] → reject if too many
    │
    ▼
[Job Factory] → create typed Job
    │
    ▼
[Pipeline] → validate → enrich
    │
    ▼
[Fan-Out: 3 Workers]
    │
    ├─ Worker 1 → [Circuit Breaker] → [Retry] → process
    ├─ Worker 2 → [Circuit Breaker] → [Retry] → process
    └─ Worker 3 → [Circuit Breaker] → [Retry] → process
    │
    ▼
[Fan-In] → merge results
    │
    ▼
[Singleton Logger] → log results
```

## ✅ Learning Checklist
- [ ] Each pattern works in isolation (days 31-39)
- [ ] Can explain when to use each pattern
- [ ] Mini-project compiles and processes 10 jobs with output
- [ ] Context cancels all workers cleanly on timeout
- [ ] Circuit breaker opens after failures, workers detect and stop

## 🛠️ Steps
1. Copy the boilerplate from `practice.go`
2. Implement each component (each has a TODO)
3. Wire them together in `main()`
4. Run: observe rate limiting, retries, circuit breaker opening
5. Change context timeout to 500ms — watch everything cancel early

## 🔍 Interview Readiness Check
After today, you should be able to:
- [ ] Whiteboard a fan-out + fan-in pipeline
- [ ] Explain circuit breaker states
- [ ] Write a goroutine that stops on context cancellation
- [ ] Implement rate limiting with time.Ticker
- [ ] Use sync.Once for thread-safe singletons
