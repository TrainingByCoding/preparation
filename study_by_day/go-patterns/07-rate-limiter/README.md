# Day 37: Rate Limiter Pattern

## 🎯 Goal
Implement a **Rate Limiter** — control how many requests/operations happen per unit of time.

## 🧠 Key Concepts
- Rate limiting = throttle throughput to protect services from overload
- Two common algorithms:
  - **Token Bucket** — tokens added at fixed rate, consumed per request (bursts allowed)
  - **Fixed Window** — max N requests per time window (simple, but can have edge-case burst)
- In Go: `time.Ticker` + channel is the idiomatic rate limiter
- Real use: API gateway limits, DB query throttling, external service calls

## 📖 Patterns

### Simple Rate Limiter with time.Ticker
```go
// Allow 1 request per 200ms
limiter := time.NewTicker(200 * time.Millisecond)

for req := range requests {
    <-limiter.C  // block until next tick
    go handleRequest(req)
}
```

### Bursty Rate Limiter (Token Bucket style)
```go
// Allow burst of 3, then 1 per 200ms
limiter := time.NewTicker(200 * time.Millisecond)
burstyLimiter := make(chan time.Time, 3) // buffer = burst size

// Pre-fill burst tokens
for i := 0; i < 3; i++ {
    burstyLimiter <- time.Now()
}

// Continuously add tokens at rate
go func() {
    for t := range limiter.C {
        burstyLimiter <- t
    }
}()

for req := range requests {
    <-burstyLimiter // consumes a token
    go handleRequest(req)
}
```

## ✅ Learning Checklist
- [ ] Implement simple rate limiter (1 req per N ms) with time.Ticker
- [ ] Implement bursty rate limiter with pre-filled channel
- [ ] Understand the difference: burst vs smooth rate limiting
- [ ] Rate limiter per-user/per-key (map of limiters)

## 🛠️ Practice Exercises
1. **Basic rate limiter** — process 5 requests, limit to 1 per 500ms, print timestamp of each
2. **Bursty limiter** — first 3 requests instant (burst), then 1 per second
3. **Per-user rate limiter** — maintain a separate limiter per user ID
4. **BONUS** — Sliding window counter: reject if more than N calls in last 60 seconds

## 🔍 Interview Questions
- "How do you implement rate limiting in Go?"
- "What's the difference between token bucket and leaky bucket?"
- "How would you rate-limit per user in a web server?"
