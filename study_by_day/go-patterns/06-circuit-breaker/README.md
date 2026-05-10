# Day 36: Circuit Breaker Pattern

## 🎯 Goal
Implement a **Circuit Breaker** — stops calling a failing service and lets it recover. Critical for resilient microservices.

## 🧠 Key Concepts
- Circuit Breaker has 3 states:
  - **Closed** (normal) — requests go through
  - **Open** (failing) — requests are rejected immediately (fail fast)
  - **Half-Open** (recovering) — allow one test request; if it succeeds → Closed, if fails → Open
- Opens after N consecutive failures
- Resets after a timeout (tries half-open)
- Real use: calling external APIs, DB, microservices

## 📖 State Machine

```
    [Closed] ──failure threshold──► [Open]
       ▲                               │
       │   success                     │ timeout
       └──────── [Half-Open] ◄─────────┘
```

### Implementation Skeleton
```go
type State int
const (
    Closed   State = iota
    Open
    HalfOpen
)

type CircuitBreaker struct {
    failures    int
    maxFailures int
    state       State
    openedAt    time.Time
    timeout     time.Duration
    mu          sync.Mutex
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    // check state, maybe transition to half-open
    cb.mu.Unlock()
    
    if cb.state == Open {
        return errors.New("circuit breaker open")
    }
    
    err := fn()
    
    cb.mu.Lock()
    if err != nil {
        cb.failures++
        if cb.failures >= cb.maxFailures {
            cb.state = Open
            cb.openedAt = time.Now()
        }
    } else {
        cb.failures = 0
        cb.state = Closed
    }
    cb.mu.Unlock()
    
    return err
}
```

## ✅ Learning Checklist
- [ ] Implement all 3 states: Closed, Open, HalfOpen
- [ ] Thread-safe with sync.Mutex
- [ ] Transitions: Closed→Open on N failures, Open→HalfOpen after timeout
- [ ] HalfOpen→Closed on success, HalfOpen→Open on failure
- [ ] Know when to use circuit breaker vs retry

## 🛠️ Practice Exercises
1. **Basic Circuit Breaker** — implement Open after 3 failures, recover after 2 seconds
2. **Test it** — simulate a service that fails 5 times then succeeds, watch state transitions
3. **Metrics** — add `TotalCalls`, `FailedCalls`, `RejectedCalls` counters
4. **BONUS** — wrap an HTTP client with a circuit breaker

## 🔍 Interview Questions
- "What problem does a circuit breaker solve?"
- "What are the 3 states and when do they transition?"
- "How is circuit breaker different from retry?"
